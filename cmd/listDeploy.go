package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	. "../db"
)

// listContainersCmd represents the listContainers command
var listDeployCmd = &cobra.Command{
	Use:   "list",
	Short: "显示deploy详细信息",
	Long:  `./run deploy list`,
	Run: func(cmd *cobra.Command, args []string) {
		tog, _ := cmd.Flags().GetBool("debug")
		logx.DebugFlag = tog
		fmt.Println("list deploy")
		p := NewPortainer()
		endpoints := p.getEndpoints()
		GenCtState(endpoints)
		fmt.Println("终端数量为:", len(endpoints))
		deploys := GetDeploys()
		for _, d := range deploys {
			fmt.Println("==========================================")
			fmt.Println("id: ", d.ID, ", name:", d.Name, ", 环境:", MakeEnvName(d.Env))
			fmt.Println("------ pods ------")
			pods := GetPods(d.ID)
			for _, p := range pods {
				//fmt.Println("endpointID:", p.Eid, ", eAddr:", MakeEdName(p.Ename), ", cid:", p.Cid)
				fmt.Println("endpointID:", p.Eid, ", Addr:", p.Addr, ", State:", CtState[p.Cid], ", cid:", p.Cid)
			}
			fmt.Println()
		}

	},
}

func init() {
	deployCmd.AddCommand(listDeployCmd)
}



