
package cmd

import (
	"github.com/spf13/cobra"
)

// listEndpointsCmd represents the listEndpoints command
var listEndpointsCmd = &cobra.Command{
	Use:   "list",
	Short: "List endpoints",
	Long: `List the endpoints available to your user on the
	currently configured Portainer instance.`,
	Run: func(cmd *cobra.Command, args []string) {
		portainer := NewPortainer()
		portainer.Endpoints = portainer.getEndpoints()

		portainer.printEndpoints()
	},
}

func init() {
	endpointCmd.AddCommand(listEndpointsCmd)
}
