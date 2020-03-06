
package cmd

import (
	"github.com/spf13/cobra"
)

// listLabelsCmd represents the listLabels command
var listLabelsCmd = &cobra.Command{
	Use:   "labels",
	Short: "List all labels",
	Long:  `List all labels associated with the given endpoint(s).`,
	Run: func(cmd *cobra.Command, args []string) {
		portainer := NewPortainer()
		portainer.Endpoints = portainer.getEndpoints()

		for _, e := range portainer.Endpoints {
			printServiceLabelsForEndpoint(e)
		}
	},
}

func init() {
	serviceCmd.AddCommand(listLabelsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listLabelsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listLabelsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
