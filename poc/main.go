package main

import (
	"log"
	"os/exec"
	"strings"
	"sync"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func main() {
	// cli, err := client.NewEnvClient()
	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"))
	if err != nil {
		panic(err)
	}

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(len(images))

	for _, image := range images {
		// fmt.Println("Image Name: ", image.RepoTags)
		// fmt.Println("Image SHA256 Sum: ", image.ID)

		go func(imageSHAID string) {
			imageID := strings.Split(imageSHAID, ":")
			cmd := exec.Command("trivy", "-f", "json", "-o", imageID[1]+"-trivy.json", imageSHAID)
			err := cmd.Run()
			if err != nil {
				log.Printf("trivy scan failed at %s and SHA sum %s with: %s\n", image.RepoTags, imageSHAID, err)
			}

		}(image.ID)

		go func(imageSHAID string) {
			imageID := strings.Split(imageSHAID, ":")
			cmd := exec.Command("dockle", "-f", "json", "-o", imageID[1]+"-dockle.json", imageSHAID)
			err := cmd.Run()
			if err != nil {
				log.Printf("dockle scan failed at %s and SHA sum %s with: %s\n", image.RepoTags, imageSHAID, err)
			}

		}(image.ID)
	}

	wg.Wait()
}

// .[].Vulnerabilities[].VulnerabilityID
// .[].Vulnerabilities[] | {VulnerabilityID,PkgName}
