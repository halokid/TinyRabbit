
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// networkCmd represents the network command
var networkCmd = &cobra.Command{
	Use:   "network",
	Short: "Work with Docker networks",
	Long:  `Subcommands of the network command work with Docker networks.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This command by itself does nothing.")
	},
}

func init() {
	rootCmd.AddCommand(networkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// networkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// networkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
