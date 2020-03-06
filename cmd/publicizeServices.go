
package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// publicizeServicesCmd represents the publicizeServices command
var publicizeServicesCmd = &cobra.Command{
	Use:   "publicize",
	Short: "Make services public",
	Long:  `Make all services except for Portainer ones public in the given endpoint(s).`,
	Run: func(cmd *cobra.Command, args []string) {
		portainer := NewPortainer()
		portainer.Endpoints = portainer.getEndpoints()

		sCount := 0
		sCountFailed := 0
		sCountTotal := 0

		for _, e := range portainer.Endpoints {
			for _, s := range e.Services {
				if !strings.Contains(s.Spec.Name, "portainer") {
					if portainer.makePublic("service", s.ID) {
						sCount++
					} else {
						sCountFailed++
					}
				}
				sCountTotal++
			}
		}

		fmt.Println("Made " + strconv.Itoa(sCount) + " services public out of " + strconv.Itoa(sCountTotal) + ", with " + strconv.Itoa(sCountFailed) + " failed.")
	},
}

func init() {
	serviceCmd.AddCommand(publicizeServicesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// publicizeServicesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// publicizeServicesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
