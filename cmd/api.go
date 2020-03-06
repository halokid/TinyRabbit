
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	. "github.com/halokid/ColorfulRabbit"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cast"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	tokenFile = "/ocx.token"
	tokenTimeOut = 3600 * 5
)

func (p Portainer) fetch(item string) string {
	bearerHeader := "Bearer " + p.token
	requestURL := p.URL + "/" + item
	req, err := http.NewRequest("GET", requestURL, nil)

	req.Header.Set("Authorization", bearerHeader)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}

func (p Portainer) login() string {
	// 读取token缓存
	tokenCache := ReadToken()
	if tokenCache != "" {
		return tokenCache
	}

	var data map[string]interface{}

	token := ""
	url := p.URL + "/auth"
	authString := `{"Username": "` + p.username + `", "Password": "` + p.password + `"}`

	//logx.DebugPrint(authString)

	jsonBlock := []byte(authString)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBlock))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		CheckFatal(err, "-------[ERROR] auth")
		//panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	_ = json.Unmarshal(body, &data)

	token = data["jwt"].(string)
	logx.DebugPrint("http code ---------", resp.StatusCode)
	logx.DebugPrint("token -----------------", token)

	// 写入token缓存
	timeNow := time.Now().Unix()
	tokenSet := cast.ToString(timeNow) + ":" + token
	WriteToken(tokenSet)

	return token
}

func ReadToken() string {
  // 读取auth的token文件
  home, err := homedir.Dir()
  c, err := ioutil.ReadFile(home + tokenFile)
  CheckError(err, "读取token文件失败")
  if err != nil {
  	return ""
	}
	fmt.Println("读取认证缓存:", string(c))
  cSpl := strings.Split(string(c), ":")
 	if len(cSpl) < 2 {
		return string(c)
	}

  timeToken := cSpl[0]
	logx.DebugPrint("timeToken ------------", timeToken)
  timeNow := time.Now().Unix()
	logx.DebugPrint("timeNow ------------", timeNow)
  timeUtil := cast.ToInt(timeNow) - cast.ToInt(timeToken)
	logx.DebugPrint("timeUtil ------------", timeUtil)
	logx.DebugPrint("tokenTimeOut ------------", tokenTimeOut)
	if timeUtil > tokenTimeOut {
		return ""
	} else {
		return cSpl[1]
	}
}

func WriteToken(token string) error {
 // 写入auth的 token文件
 home, err := homedir.Dir()
 tb :=[]byte(token)
 err = ioutil.WriteFile(home + tokenFile , tb, 0777)
 CheckError(err, "写入token文件失败")
 return err
}

func (p Portainer) makePublic(resourceType string, id string) bool {
	data := map[string]interface{}{
		"Type":       resourceType,
		"Public":     true,
		"ResourceID": id,
	}
	return p.Post(data, "resource_controls")
}

func (p Portainer) Post(data map[string]interface{}, path string) bool {
	bearerHeader := "Bearer " + p.token
	requestURL := p.URL + "/" + path

	logx.DebugPrint("requestURl --------", requestURL)

	bytesData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(bytesData))

	req.Header.Set("Authorization", bearerHeader)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	if p.verbose {
		fmt.Println("Sent request with data: " + string(bytesData))
		fmt.Println("Status " + resp.Status + " received from API, response was: " + string(body))
	}

	if resp.StatusCode == 200 {
		return true
	}

	logx.DebugPrint("Post rsp -----------", string(body), resp.StatusCode)
	return false
}
