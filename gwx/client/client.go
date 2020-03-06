package main

import (
  "context"
  "fmt"
  . "github.com/halokid/ColorfulRabbit"
  "google.golang.org/grpc"
  pb "../proto"
)

const (
  addr = "localhost:9528"
)

func main() {
  conn, err := grpc.Dial(addr, grpc.WithInsecure())
  CheckFatal(err, "建立RPC客户端连接句柄失败")
  defer conn.Close()
  c := pb.NewGwxClient(conn)

  // 发送动作给服务端
  // server
  reqData := `{"act": "setServer", "val": "xxx"}`
  //reqData := `hello`
  r, err := c.GwxDo(context.Background(), &pb.GwxReq{ReqData: reqData})

  // upstream
  //reqData = `{"act": "setUpstream", "val": "xxx-!-8.8.8.8:2222"}`
  reqData = `{"act": "setUpstream", "val": "xxx-!-8.8.8.8:2222;9.9.9.9:3333"}`
  r, err = c.GwxDo(context.Background(), &pb.GwxReq{ReqData: reqData})

  CheckFatal(err, "发送RPC请求失败")
  fmt.Println(r.RspData)
}