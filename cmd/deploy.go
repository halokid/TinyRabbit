
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// containerCmd represents the container command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "管理容器的部署",
	Long:  "./run deploy list 查看容器明细\n\r./run deploy log <deploy> <podId> 查看deploy某个pod的日志",
	Run: func(cmd *cobra.Command, args []string) {
		tog, _ := cmd.Flags().GetBool("debug")
		logx.DebugFlag = tog
		fmt.Println("This command by itself does nothing.")
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serviceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serviceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


