
package cmd

import (
	"github.com/spf13/cobra"
)

// listNetworksCmd represents the listNetworks command
var listNetworksCmd = &cobra.Command{
	Use:   "list",
	Short: "List networks",
	Long:  `List all networks associated with the given endpoint(s).`,
	Run: func(cmd *cobra.Command, args []string) {
		portainer := NewPortainer()
		portainer.Endpoints = portainer.getEndpoints()

		for _, e := range portainer.Endpoints {
			printNetworksForEndpoint(e)
		}
	},
}

func init() {
	networkCmd.AddCommand(listNetworksCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listNetworksCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listNetworksCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
