
package cmd

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"log"
	"strings"

	. "../db"
)

// listContainersCmd represents the listContainers command
var getDeployLogCmd = &cobra.Command{
	Use:   "log",
	Short: "显示deploy容器的log",
	Long:  `./run deploy log <svc名称> <svc第几个容器>`,
	Run: func(cmd *cobra.Command, args []string) {
		tog, _ := cmd.Flags().GetBool("debug")
		logx.DebugFlag = tog

		if len(args) < 2 {
			logx.DebugPrint("需要<svc> <num>")
			log.Fatal("缺少参数")
		}
		svcName := args[0]
		num := cast.ToInt(args[1])
		// 取得svc的第num 条记录
		pods := GetPodsBySvc(svcName, "gwx") // 默认取得生产的容器日志
		if strings.HasSuffix(svcName, "-dev") {
			pods = GetPodsBySvc(svcName, "gwxtest") // 默认取得生产的容器日志
		}
		if len(pods) < num || num == 0 {
			fmt.Println("svc的容器数量少于输入值<num> 或num为0")
			return
		}
		eId := pods[num-1].Eid
		cId := pods[num-1].Cid
		// 输出日志
		printLog(cast.ToString(eId), cId)
	},
}

func init() {
	deployCmd.AddCommand(getDeployLogCmd)
}



