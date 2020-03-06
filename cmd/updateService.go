
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// updateServiceCmd represents the updateService command
var updateServiceCmd = &cobra.Command{
	Use:   "update",
	Short: "Trigger an update of the specified service",
	Long: `Trigger an update of the specified service.
	Note: this will update ALL services that match the given name/ID across all given endpoints.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("updateService called with " + args[0])

		portainer := NewPortainer()
		portainer.Endpoints = portainer.getEndpoints()

		path := ""

		data := map[string]interface{}{
			"TaskTemplate": map[string]interface{}{
				"ForceUpdate": 1,
			},
		}

		for _, e := range portainer.Endpoints {
			for _, s := range e.Services {
				if s.ID == args[0] || s.Spec.Name == args[0] {
					fmt.Println("Updating service " + s.Spec.Name + "...")
					path = "/services/" + s.ID + "/?version=" + strconv.Itoa(s.Version.Index)
					portainer.Post(data, path)
				}
			}
		}
	},
}

func init() {
	serviceCmd.AddCommand(updateServiceCmd)
}
