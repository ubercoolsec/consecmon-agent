package lib

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	log "github.com/sirupsen/logrus"
)

type ContainerEngineOpts struct {
	ScanAllImages bool
}

func EnumRunningContainers(opts *ContainerEngineOpts) {
	log.Info("Enumerating running Docker containers")

	cli, err := client.NewClientWithOpts(client.WithVersion("1.40"))
	if err != nil {
		log.Fatal("Failed to created Docker client: ", err)
	}

	if opts.ScanAllImages {
		images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
		if err != nil {
			log.Fatal("Failed to enumerate local Docker images")
		}

		for _, image := range images {
			fmt.Printf("%s\n", image.ID)
		}
	} else {
		containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
		if err != nil {
			log.Fatal("Failed to list running containers: ", err)
		}

		for _, container := range containers {
			fmt.Printf("%s %s [%s]\n", container.ID, container.Image, container.ImageID)
		}
	}
}
