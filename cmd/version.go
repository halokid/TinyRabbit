
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// containerCmd represents the container command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show version",
	Long:  `show version detail`,
	Run: func(cmd *cobra.Command, args []string) {
		version := "9527"
		fmt.Println("version:", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

}


