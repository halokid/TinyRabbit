
package cmd

import (
  . "../db"
  . "../gwx"
  pb "../gwx/proto"
  "context"
  "google.golang.org/grpc"
  "os"
  "strings"

  "fmt"
  . "github.com/halokid/ColorfulRabbit"
  "github.com/spf13/cast"
  "github.com/spf13/cobra"
  //"os"
)

const (
  parLen = 4
)

// listContainersCmd represents the listContainers command
var svcDeployCmd = &cobra.Command{
  Use:   "svc",
  Short: "部署deploy",
  Long:  `./run deploy svc <svcName> <imageName> <port:hostPort> <scale>`,
  Run: func(cmd *cobra.Command, args []string) {
    tog, _ := cmd.Flags().GetBool("debug")
    logx.DebugFlag = tog
    fmt.Println(len(args), args)
    if len(args) < parLen {
      fmt.Println("参数不足或错误")
    }
    fmt.Println("开始部署deploy svc")
    svcNameOri := args[0]
    imgName := args[1]
    ports := args[2]
    scaleNum := args[3]
    gwEnv := "gwx"
    if len(args) == 5 {
      gwEnv = args[4]
    }
    fmt.Println("gwEnv ------------- ", gwEnv)
    //os.Exit(11)
    //cPort, hostPort := GetPorts(ports)
    cPort := ports
    // 检查svc是否已存在
    svcName := svcNameOri + MakeEnvDeployName(gwEnv)    // 重写svcName，加上环境参数
    d := GetOneDeploy(svcName, gwEnv)
    hostPort := GetRandomPort()
    if d.ID != 0 {
      // 已存在svc，则用旧端口
      hostPort = GetPodPortBySvc(svcName)
    }

    cName := svcName			// 用服务名来作为容器名称, 添加gwx区别容器环境

    p := NewPortainer()
    es := p.getEndpoints()
    logx.DebugPrint("len es ------------", len(es))

    if len(es) == 1 &&  cast.ToInt(scaleNum) > 1{
      // 只有只有一个节点, scale 的数量不能大于1, 因为会端口冲突
      fmt.Println("只有一个节点,scale数量不能大于1")
      os.Exit(500)
    }

    var scale Scalex
    scale = &Scale{}

    // 假如deploy表里面没有记录，则采用新创建算法
    pods := GetPodsBySvc(svcName, gwEnv)
    var hitEsIdx []int
    var hitEsNames []string
    if len(pods) == 0 {
      tmpIdx := scale.Random(cast.ToInt(scaleNum), len(es))
      for _, h := range tmpIdx {
        hitEsIdx = append(hitEsIdx, es[h].ID)
        hitEsNames = append(hitEsNames, es[h].Name)
      }
    } else {      // 如果对应的服务已存在容器
      for _, v := range pods {
        hitEsIdx = append(hitEsIdx, v.Eid)
        hitEsNames = append(hitEsNames, v.Ename)
      }
    }
    logx.DebugPrint("hitEsIdx ------------- ", hitEsIdx)
    logx.DebugPrint("hitEsNames ------------- ", hitEsNames)
    //os.Exit(500)

    var dpEid []int
    var dpEpName []string
    var cIds []string
    for i, eIdx := range hitEsIdx {
      // 开始部署容器
      cId := p.DeployCt(eIdx, cName, imgName, cPort, hostPort, hitEsNames[i])
      //fmt.Println("------ 部署容器到 ----- eid:", es[eIdx].ID, "cName:", cName, "imgName: ",
      //							imgName, "ports:", ports)
      cIds = append(cIds, cId)
      dpEid = append(dpEid, eIdx)
      dpEpName = append(dpEpName, hitEsNames[i])
    }

    logx.DebugPrint("dpEid ---------------------", dpEid)
    logx.DebugPrint("dpEpName ---------------------", dpEpName)
    for i, ei := range dpEid {
      fmt.Println("------ 部署容器到 ----- eid:", ei, ", eName:", MakeEdName(dpEpName[i]), ", cId:", cIds[i], ", cName:", cName, ", imgName: ", imgName, ", cPort:", ports, ", hostPort:", hostPort)
    }

    // 添加deploy信息
    //d := GetOneDeploy(svcName)
    var nodes []string
    if d.ID == 0 {				// 还没存在这个 svc
      fmt.Println("----------------- 新增加svc", svcName, "-----------------")
      addD := Deploy{Name: svcName , Env:  gwEnv}
      AddDeploy(&addD)

      for i, er := range dpEid {
        addr := MakeEdName(dpEpName[i]) + ":" + hostPort
        nodes = append(nodes, addr)
        p := Pod{
          Did:    addD.ID,
          Name:   svcName,
          Eid:    er,
          Ename:  dpEpName[i],
          Addr: 	addr,
          Env:    gwEnv,
          Cid:    cIds[i],
        }
        err := AddPod(p)
        CheckFatal(err, "-------------添加pod记录错误")
      }

      // 更新gwx
      UpdateGwx(svcNameOri, nodes, gwEnv)

    } else {				// 已存在svc， 只更新pod
      fmt.Println("----------------- 已存在svc", svcName, "-----------------")
      for i, er := range dpEid {
        /**
        p := Pod{
          Did:    d.ID,
          Name:   svcName,
          Eid:    er,
          Ename:  dpEpName[i],
          Cid:    cIds[i],
        }
        */
        p := Pod{}
        Db.Where("name = ? and eid = ?", svcName, er).Take(&p)
        p.Cid = cIds[i]
        err := UpdatePod(&p)
        CheckFatal(err, "-------------更新pod记录错误")
      }
    }

  },
}

func UpdateGwx(svcName string, nodes []string, gwEnv string) error {
  // 更新gwx网关
  ngxAddr := Ngxx.RpcAddr
  if gwEnv == "gwxtest" {
    ngxAddr = Ngxx.RpcAddrTest
  }
  fmt.Println("gwx服务端地址:", ngxAddr)
  conn, err := grpc.Dial(ngxAddr, grpc.WithInsecure())
  CheckFatal(err, "建立RPC客户端连接句柄失败")
  defer conn.Close()
  c := pb.NewGwxClient(conn)
  // set server
  reqData := `{"act": "setServer", "val": "` + svcName + `"}`
  r, err := c.GwxDo(context.Background(), &pb.GwxReq{ReqData: reqData})
  CheckFatal(err, "发送RPC请求失败")
  fmt.Println(r.RspData)

  // set updtream
  nodesJoin := strings.Join(nodes, ";")
  reqData = `{"act": "setUpstream", "val": "` + svcName + `-!-` + nodesJoin + `"}`
  r, err = c.GwxDo(context.Background(), &pb.GwxReq{ReqData: reqData})

  // realod gwx
  reqData = `{"act": "reloadGwx", "val": "reload gwx"}`
  r, err = c.GwxDo(context.Background(), &pb.GwxReq{ReqData: reqData})

  // reaload ngx
  //Ngxx.ReloadNgx()

  return nil
}

func init() {
  deployCmd.AddCommand(svcDeployCmd)
}



