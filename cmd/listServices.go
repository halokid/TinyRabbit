
package cmd

import (
	"github.com/spf13/cobra"
)

var filterServicesBroken bool

// listServicesCmd represents the listServices command
var listServicesCmd = &cobra.Command{
	Use:   "list",
	Short: "List services",
	Long:  `List all services running in the specified endpoint(s).`,
	Run: func(cmd *cobra.Command, args []string) {
		portainer := NewPortainer()
		portainer.Endpoints = portainer.getEndpoints()

		for _, e := range portainer.Endpoints {
			if filterServicesBroken {
				printBrokenServicesForEndpoint(e)
			} else {
				printServicesForEndpoint(e)
			}
		}
	},
}

func init() {
	serviceCmd.AddCommand(listServicesCmd)
	listServicesCmd.Flags().BoolVarP(&filterServicesBroken, "broken", "b", false, "Display only broken services")
}
