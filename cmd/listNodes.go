
package cmd

import (
	"github.com/spf13/cobra"
)

// listNodesCmd represents the listNodes command
var listNodesCmd = &cobra.Command{
	Use:   "list",
	Short: "List nodes",
	Long:  `List all nodes in the specified endpoint(s).`,
	Run: func(cmd *cobra.Command, args []string) {
		portainer := NewPortainer()
		portainer.Endpoints = portainer.getEndpoints()

		for _, e := range portainer.Endpoints {
			printNodesForEndpoint(e)
		}
	},
}

func init() {
	nodeCmd.AddCommand(listNodesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listNodesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listNodesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
