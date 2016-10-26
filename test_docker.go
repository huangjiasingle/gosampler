package main

import (
	"fmt"
	"github.com/fsouza/go-dockerclient"
)

func main() {
	endpoint := "unix:///var/run/docker.sock"
	client, _ := docker.NewClient(endpoint)
	cts, _ := client.ListContainers(docker.ListContainersOptions{All: false})

	fmt.Printf("containers is %#v", cts)
}
