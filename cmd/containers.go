package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/go-openapi/errors"
	. "github.com/halokid/ColorfulRabbit"
	"github.com/spf13/cast"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)


type cPort struct {
	IP 							string
	PrivatePort		 	int
	PublicPort			int
	Type						string
}

// Container is a Docker container
type Container struct {
	ID    		string
	Image 		string
	State 		string
	Ports			[]cPort
	Names 		[]string
}

// 容器的运行状态
var CtState map[string]string

func (p Portainer) getContainersForEndpoint(endpoint Endpoint) []Container {
	// 获取endpoint的容器信息
	// fixme: 这里少了一个 all=1 的参数， 所以只获取了状态为 running 的容器
	output := p.fetch("endpoints/" + strconv.Itoa(endpoint.ID) + "/docker/containers/json?all=1")
	logx.DebugPrint("getContainersForEndpoint output --------------", output)

	containers := make([]Container, 0)

	json.Unmarshal([]byte(output), &containers)
	logx.DebugPrint("getContainersForEndpoint containers --------------", containers)

	return containers
}

func (p Portainer) populateContainersForEndpoints(endpoints []Endpoint) []Endpoint {
	newEndpoints := []Endpoint{}
	var endpoint Endpoint

	for _, e := range endpoints {
		endpoint = e
		endpoint.Containers = p.getContainersForEndpoint(e)

		newEndpoints = append(newEndpoints, endpoint)
	}

	return newEndpoints
}

func GenCtState(endpoints []Endpoint) error {
  // 生成容器状态的map
  logx.DebugPrint("len Ctstate ------------", len(CtState))
  if len(CtState) > 0 {
    fmt.Println("xxxxx")
    return nil
  }
  CtState = make(map[string]string)
  for _, e := range endpoints {
    for _, c := range e.Containers {
      CtState[c.ID]  = c.State
    }
  }
  return nil
}

func printContainersForEndpoint(endpoint Endpoint) {
	fmt.Println(endpoint.Name, MakeEdName(endpoint.Name), endpoint.ID, "容器列表")
	fmt.Println("----")

	for _, c := range endpoint.Containers {
		fmt.Println("ID: " + c.ID[0:12] + ", Name:", c.Names, ", State:", c.State, ", Ip:", MakeEdName(endpoint.Name),
									"Port:", c.Ports, ", image: " + c.Image)
	}
	fmt.Println("----")
}


func (p *Portainer) PullImage(eId int, imageName string) bool {
	// pull image
	data := make(map[string]interface{})
	data["fromImage"] = imageName
	data["tag"] = "latest"
	pathImg := strings.Replace(imageName, "/", "%2F", -1)
	logx.DebugPrint("pathImg -----------", pathImg)
	path := "endpoints/" + cast.ToString(eId) + "/docker/images/create?fromImage=" + pathImg +
									"&tag=latest"
	logx.DebugPrint("PullImage path --------------", path)
	rspErr := p.Post(data, path)
	return rspErr
}

func (p *Portainer) CreateCt(eId int, imgName string, cName string, cPort string,
															hostPort string) (string, error) {
	// create container， 返回容器id， 创建了之后，还需要start才能运行
	//data := make(map[string]interface{})
	data := makeCtMap(imgName, cName, cPort, hostPort)
	logx.DebugPrint("CreateCt data ---------- ", data)
	path := p.URL + `/endpoints/` + cast.ToString(eId) + `/docker/containers/create?name=` + cName
	logx.DebugPrint("CreateCt path ----------------- ", path)
	js := []byte(data)
	req, err := http.NewRequest("POST", path, bytes.NewBuffer(js))
	bearerHeader := "Bearer " + p.token
	req.Header.Set("Authorization", bearerHeader)
	req.Header.Set("Content-Type", "application/json")
	CheckError(err, "----- CreateCt set req json err")

	client := &http.Client{}
	resp, err := client.Do(req)
	CheckError(err, "------ CreateCt request error")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	logx.DebugPrint("CreateCt resp ----------- ", string(body))

	bdJs, err := simplejson.NewJson(body)
	CheckError(err)
	return bdJs.Get("Id").MustString(), err
}


func (p *Portainer) StartCt(eId int, cId string, esName string) bool {
	// 启动容器
	data := make(map[string]interface{})
	path := "endpoints/" + cast.ToString(eId) + "/docker/containers/" + cId + "/start"
	rsp := p.Post(data, path)
	if !rsp {
		fmt.Println("[ERROR] ------------- eid:", eId,  ",eName:", MakeEdName(esName), ",cId:", cId, "启动错误-----" )
		return false
	}
	return true
}

func (p *Portainer) RmCt(eId int, cName string) {
	// 直接删除容器， 包括停止和删除
	// fixme: 这个删除动作会不会太粗暴， 影响服务的高可用？
	endpoints := p.getEndpoints()
	fmt.Println("endpoints 个数为------", len(endpoints))
	var eIdx Endpoint
	// 获取endpoint id
	for _, e := range endpoints {
		fmt.Println(e.ID)
		if e.ID == eId {
			eIdx = e
			break
		}
	}
	cId := ""
	cNameAdd := "/" + cName
	for _, c := range eIdx.Containers {
		if c.Names[0] == cNameAdd {
			cId = c.ID
			break
		}
	}

	fmt.Println("删除的容器ID为----------------", cId)

	// 删除
	path := p.URL + "/endpoints/" + cast.ToString(eId) + "/docker/containers/" + cId + "?force=true&v=0"
	fmt.Println("RmCt path --------------", path)
	req, err := http.NewRequest("DELETE", path, nil)
	bearerHeader := "Bearer " + p.token
	req.Header.Set("Authorization", bearerHeader)
	req.Header.Set("Content-Type", "application/json")
	CheckError(err, "----- CreateCt set req json err")
	client := &http.Client{}
	resp, err := client.Do(req)
	CheckError(err, "------ CreateCt request error")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	logx.DebugPrint("CreateCt resp ----------- ", string(body))

	fmt.Println("删除容器结果----------", string(body))
}

func (p *Portainer) StopCt() {

}

func (p *Portainer) DeployCt(eId int, cName string, imgName string, cPort string,
															hostPort string, esName string) string {
	// 统一封装的部署容器的流程
	p.RmCt(eId, cName)					// stop & rm
	pullRes := p.PullImage(eId, imgName)					// pull latest image
	logx.DebugPrint("pullRes ----------", pullRes)
		if !pullRes {
			CheckFatal(errors.New(500, "pull容器失败"))
		}
	time.Sleep(1 * time.Second)

	// create container
	cId, err := p.CreateCt(eId, imgName, cName, cPort, hostPort)
	logx.DebugPrint("cId -------------- ", cId)
	CheckFatal(err, " ------------ 创建容器失败")

	// start container
	startRes := p.StartCt(eId, cId, esName)
	logx.DebugPrint("startRes -------------- ", startRes)

	/**
	if !startRes {				// 如果启动失败， 直接crash
		fmt.Println("------------ 容器启动失败， 终止构建 ------------")
		os.Exit(500)
	}
	*/

	return cId
}









