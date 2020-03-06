
package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Endpoint is a Docker Swarm endpoint
type Endpoint struct {
	ID         int
	Name       string
	Containers []Container
	Networks   []Network
	Services   []Service
	Tasks      []Task
	Nodes      []Node
}

func (p Portainer) getEndpoints() []Endpoint {
	endpoints := make([]Endpoint, 0)

	if endpointID != 0 {
		output := p.fetch("endpoints/" + strconv.Itoa(endpointID))

		endpoint := Endpoint{}

		json.Unmarshal([]byte(output), &endpoint)

		endpoints = append(endpoints, endpoint)
	} else {
		output := p.fetch("endpoints")

		json.Unmarshal([]byte(output), &endpoints)
	}

	endpoints = p.populateServicesForEndpoints(endpoints)
	endpoints = p.populateContainersForEndpoints(endpoints)
	endpoints = p.populateNetworksForEndpoints(endpoints)
	endpoints = p.populateNodesForEndpoints(endpoints)
	endpoints = p.populateTasksForEndpoints(endpoints)

	return endpoints
}

func (p Portainer) printEndpoints() {
	for _, e := range p.Endpoints {
		fmt.Println(strconv.Itoa(e.ID) + ": " + e.Name + " (" + strconv.Itoa(len(e.Services)) + " services, " + strconv.Itoa(len(e.Containers)) + " containers, " + strconv.Itoa(len(e.Networks)) + " networks)")
	}
	fmt.Println("----")
}
