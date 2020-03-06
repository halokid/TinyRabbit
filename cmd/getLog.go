
package cmd

import (
	"fmt"
	"github.com/prometheus/common/log"
	"github.com/spf13/cobra"
)

// listContainersCmd represents the listContainers command
var getLogCmd = &cobra.Command{
	Use:   "get",
	Short: "获取容器日志",
	Long:  `需要endpoint id 和 容器id, ./ocx log get <eid> <cid>`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 || len(args) < 2 {
			logx.DebugPrint("需要endpoint id 和 容器id, ./ocx log get <eid> <cid>")
			log.Fatal("缺少参数")
		}
		fmt.Println(args)
		eId, cId  := args[0], args[1]
		printLog(eId, cId)
	},
}

func init() {
	logCmd.AddCommand(getLogCmd)
}



