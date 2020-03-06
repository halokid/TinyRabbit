
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// endpointCmd represents the endpoint command
var endpointCmd = &cobra.Command{
	Use:   "endpoint",
	Short: "Work with Portainer endpoints",
	Long:  `Work with Portainer endpoints, which usually means Docker Swarms.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This command does nothing.")
	},
}

func init() {
	rootCmd.AddCommand(endpointCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// endpointCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// endpointCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
