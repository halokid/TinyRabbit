
package cmd

import (
	"github.com/spf13/cobra"
)

// listContainersCmd represents the listContainers command
var listContainersCmd = &cobra.Command{
	Use:   "list",
	Short: "List containers",
	Long:  `List all containers running in the specified endpoint(s).`,
	Run: func(cmd *cobra.Command, args []string) {

		tog, _ := cmd.Flags().GetBool("debug")
		logx.DebugFlag = tog

		portainer := NewPortainer()
		portainer.Endpoints = portainer.getEndpoints()

		for _, e := range portainer.Endpoints {
			printContainersForEndpoint(e)
		}
	},
}

func init() {
	containerCmd.AddCommand(listContainersCmd)
}
