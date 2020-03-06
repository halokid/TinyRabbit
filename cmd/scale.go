package cmd

import (
  . "github.com/halokid/ColorfulRabbit"
  "github.com/spf13/cast"
  "math/rand"
  "time"

  . "../db"
)

/**
弹性计算策略
*/

const (
  minPort = 30000
  maxPort = 32767
)

type Scalex interface {
  Random(scaleNum int, esLen int) []int
}

type Scale struct {
  // 策略种类，  随机， 均衡等
  Strategy      string
}

func (s *Scale) Random(scaleNum int, esLen int) []int {
  // 随机策略，返回endpointId slice
  //esSl := MakeRange(0, esLen - 1)
  //esSl := []int{0, 1, 2, 3, 4}
  //hitEsIdx := RandomSl(scaleNum, esSl)

  esSl := make([]int, 0)
  for i := 0; i < esLen; i++ {
    esSl = append(esSl, i)
  }
  hitEsIdx := RandomSl(scaleNum, esSl)

  return hitEsIdx
}

func GetRandomPort() string {
  // 获取随机端口， 排除现有已经存在svc的端口
  MakePort:
  rand.Seed(time.Now().Unix())
  hostPort := cast.ToString(RandInt(minPort, maxPort))
  pods := GetPodsByPort(hostPort)
  if len(pods) > 0 {
    goto MakePort
  }
  return hostPort
}






