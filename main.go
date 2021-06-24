package main

import (
	"flag"
	"fmt"
	d "github/axieCD/package/docker"
	"log"
	"os"

	"github.com/docker/docker/client"
)

type deployment struct {
	container_name string
	container_port string
	image_registry string
	envfile        string
	mount          string
}

// Read the config file from the current directory and marshal
// into the conf config struct.

func main() {

	deploy := new(deployment)

	flag.StringVar(&deploy.container_name, "c", "", "Container name")
	flag.StringVar(&deploy.image_registry, "i", "", "Container image")
	flag.StringVar(&deploy.container_port, "p", "", "Port open")
	flag.StringVar(&deploy.envfile, "e", "", "Env file")
	flag.StringVar(&deploy.mount, "m", "", "Mount volume ")
	flag.Parse()

	host, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Println("----- Deploy new container \n----- Name: " + deploy.container_name + "\n----- Image: " + deploy.image_registry + "\n----- Port: " + deploy.container_port + "\n----- Env file: " + deploy.envfile + "\n----- Volume mount: " + deploy.mount + "\n----> Host: " + host)

	client, err := client.NewEnvClient()
	if err != nil {
		fmt.Printf("Unable to create docker client: %s", err)
	}

	// Stops and removes a container
	d.StopAndRemoveContainer(client, deploy.container_name)

	// Load env from file
	env := d.LoadEnv(deploy.envfile)

	// Run new container
	err = d.RunContainer(client, deploy.image_registry, deploy.container_name, deploy.container_port, env)
	if err != nil {
		log.Println(err)
	}
}
