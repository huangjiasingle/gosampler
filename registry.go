package main

import (
	"fmt"

	// "github.com/docker/distribution/digest"
	"github.com/docker/distribution/manifest"
	// "github.com/docker/libtrust"
	"github.com/heroku/docker-registry-client/registry"
)

func main() {
	url := "http://172.17.11.2:5000"
	username := "" // anonymous
	password := "" // anonymous

	hub, err := registry.New(url, username, password)
	manifest, err := hub.Manifest("172.17.11.2:5000/test-app", "14")
	fmt.Println(manifest, err)
}
