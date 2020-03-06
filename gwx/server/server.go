package main

import (
  . ".."
  pb "../proto"
  "context"
  "github.com/bitly/go-simplejson"
  . "github.com/halokid/ColorfulRabbit"
  "google.golang.org/grpc"
  "google.golang.org/grpc/reflection"
  "log"
  "net"
  "strings"
  "fmt"
)

const (
  port = ":9528"
  maxQueueSize = 10000      // 服务端最长的处理队列长度，为避免服务端压力过大
  maxWorkds = 50            // 并发处理的worker数量
)

type server struct {}

func (s *server) GwxDo(ctx context.Context, req *pb.GwxReq) (*pb.GwxRsp, error) {
  log.Printf("收到请求: %s", req.ReqData)
  //return &pb.GwxRsp{RspData: "doing: " + req.ReqData}, nil
  reqData := req.ReqData
  reqJs, err := simplejson.NewJson([]byte(reqData))
  CheckError(err, "读取客户端数据格式失败", reqJs)
  if err != nil {
    return &pb.GwxRsp{}, nil
  }
  act := reqJs.Get("act").MustString()

  if act == "setServer" {
    fmt.Println("setServer doing...")
    servName := reqJs.Get("val").MustString()
    Ngxx.SetServer(servName)
  } else if act == "setUpstream" {
    fmt.Println("setUpstream doing...")
    val := reqJs.Get("val").MustString()
    var nodes []string
    upstream := strings.Split(val, "-!-")[0]
    nodesOri := strings.Split(val, "-!-")[1]
    for _, sx := range strings.Split(nodesOri, ";") {
      nodes = append(nodes, sx)
    }
    Ngxx.SetUpstream(upstream, nodes)
  } else if act == "reloadGwx" {
    fmt.Println("reloadGwx doing...")
    Ngxx.ReloadNgx()
  }

  return &pb.GwxRsp{RspData: "reqData: " + req.ReqData}, nil
}

func main() {
  lis, err := net.Listen("tcp", port)
  CheckFatal(err, "建立RPC服务端监听端口失败")
  log.Println("gwx rpc服务端运行:", port)
  s := grpc.NewServer()
  pb.RegisterGwxServer(s, &server{})
  reflection.Register(s)
   if err := s.Serve(lis); err != nil {
     log.Fatalf("建立RPC服务端失败: %v", err)
   }
}





