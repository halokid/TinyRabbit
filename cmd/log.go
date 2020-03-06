
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// containerCmd represents the container command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "日志管理相关操作",
	Long:  `日志相关的一切功能入口参数`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This command by itself does nothing.")
	},
}

func init() {
	rootCmd.AddCommand(logCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serviceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serviceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
