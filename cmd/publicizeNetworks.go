
package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// publicizeNetworksCmd represents the publicizeNetworks command
var publicizeNetworksCmd = &cobra.Command{
	Use:   "publicize",
	Short: "Make networks public",
	Long:  `Make all networks public (except for Portainer ones) in the given endpoint(s).`,
	Run: func(cmd *cobra.Command, args []string) {
		portainer := NewPortainer()
		portainer.Endpoints = portainer.getEndpoints()

		nCount := 0
		nCountFailed := 0
		nCountTotal := 0

		for _, e := range portainer.Endpoints {
			for _, n := range e.Networks {
				if !strings.Contains(n.Name, "portainer") {
					if portainer.makePublic("service", n.ID) {
						nCount++
					} else {
						nCountFailed++
					}
				}
				nCountTotal++
			}
		}

		fmt.Println("Made " + strconv.Itoa(nCount) + " networks public out of " + strconv.Itoa(nCountTotal) + ", with " + strconv.Itoa(nCountFailed) + " failed.")
	},
}

func init() {
	networkCmd.AddCommand(publicizeNetworksCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// publicizeNetworksCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// publicizeNetworksCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
