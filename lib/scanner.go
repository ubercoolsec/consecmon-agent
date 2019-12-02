package lib

import (
	log "github.com/sirupsen/logrus"
)

func ScanImage(imageChannel chan string) {
	for imageId := range imageChannel {
		log.Info("Scanning image: ", imageId)
	}

	log.Info("Scanner finished")
}
