
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// containerCmd represents the container command
var containerCmd = &cobra.Command{
	Use:   "container",
	Short: "Work with Docker Swarm containers",
	Long:  `Subcommands of container work with Docker Swarm containers.`,
	Run: func(cmd *cobra.Command, args []string) {
		tog, _ := cmd.Flags().GetBool("debug")
		logx.DebugFlag = tog
		fmt.Println("This command by itself does nothing.")
	},
}

func init() {
	rootCmd.AddCommand(containerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serviceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serviceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


