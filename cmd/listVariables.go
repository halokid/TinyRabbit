
package cmd

import (
	"github.com/spf13/cobra"
)

// listVariablesCmd represents the listVariables command
var listVariablesCmd = &cobra.Command{
	Use:   "variables",
	Short: "List environment variables",
	Long:  `List all environment variables for a given endpoint(s).`,
	Run: func(cmd *cobra.Command, args []string) {
		portainer := NewPortainer()
		portainer.Endpoints = portainer.getEndpoints()

		for _, e := range portainer.Endpoints {
			printServiceVariablesForEndpoint(e)
		}
	},
}

func init() {
	serviceCmd.AddCommand(listVariablesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listVariablesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listVariablesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
