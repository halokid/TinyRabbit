
package cmd

import (
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"os"
)

// listContainersCmd represents the listContainers command
var deployContainersCmd = &cobra.Command{
	Use:   "create",
	Short: "部署单个容器",
	Long:  `部署容器的相关操作.`,
	Run: func(cmd *cobra.Command, args []string) {
		tog, _ := cmd.Flags().GetBool("debug")
		logx.DebugFlag = tog
		if len(args) == 0 {
			logx.DebugPrint("./run <endpointId> <imageName> <containerName>" +
				" <potr:hostPort> ")
			os.Exit(404)
		}
		eId := cast.ToInt(args[0])
		imgName := args[1]
		cName := args[2]
		ports := args[3]
		cPort, hostPort := GetPorts(ports)

		p := NewPortainer()
		/*
		// 停止删除容器
		p.RmCt(eId, cName)

		// pull image
		pullRes := p.PullImage(eId, imgName)
		logx.DebugPrint("pullRes ----------", pullRes)
		if !pullRes {
			fmt.Println("pull容器失败")
			os.Exit(500)
		}

		time.Sleep(1 * time.Second)
		// create container
		cId, err := p.CreateCt(eId, imgName, cName, cPort, hostPort)
		logx.DebugPrint("cId -------------- ", cId)
		CheckFatal(err, " ------------ 创建容器失败")

		// start container
		startRes := p.StartCt(eId, cId)
		logx.DebugPrint("startRes -------------- ", startRes)
		*/

		p.DeployCt(eId, cName, imgName, cPort, hostPort, "eId为"+cast.ToString(eId))
	},
}

func init() {
	containerCmd.AddCommand(deployContainersCmd)
}



