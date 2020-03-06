package cmd

import (
	. "github.com/halokid/ColorfulRabbit"
	"math/rand"
	"strings"
	"time"
)

/**
通用代码封装
 */

var epNamesMap map[string]string

func init() {
	epNamesMap = make(map[string]string)
	epNamesMap["local"] = "8.8.8.8"
}

func GetPorts(ports string) (string, string) {
  // 拆分端口
  cPort := strings.Split(ports, ":")[0]
	hostPort := strings.Split(ports, ":")[1]
	return cPort, hostPort
}

func MakeRange(min, max int) []int {
	// 获取range slice
	a := make([]int, max - min+ 1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func RandomSl(n int, sl []int) []int {
	// 随机抽取数组的n个不重复元素
	//rn := make([]int, n)
	var rn []int
	i := 0
	rand.Seed(time.Now().Unix())
	for i < n {
		Loop:
		//if k := rand.Intn(len(sl) - 1); IndexOfI(k, rn) == -1 {
		if k := rand.Intn(len(sl)); IndexOfI(k, rn) == -1 {
			logx.DebugPrint("k -----------", k)
			rn = append(rn, k)
		} else {
			goto Loop
		}
		i++
		logx.DebugPrint("i -----------", i)
		logx.DebugPrint("rn -------------", rn)
	}
	return rn
}

func MakeEdName(edName string) string {
	// 封装输出 endpoint名称
	if edIp, ok := epNamesMap[edName]; ok {
		return edIp
	}
	return edName
}

func MakeEnvName(env string) string {
	// 封装输出 运行环境名称
	envMap:= make(map[string]string)
	envMap["gwx"] = "生产"
	envMap["gwxtest"] = "开发"
	return envMap[env]
}

func MakeEnvDeployName(env string) string {
	if env == "gwxtest" {
		return "-dev"
	}
	return ""
}

