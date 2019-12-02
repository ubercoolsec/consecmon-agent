package lib

import (
	log "github.com/sirupsen/logrus"
)

type ScannerOpt struct {
	ScanWithTrivy  bool
	ScanWithDockle bool
}

func ScanImage(imageChannel chan string, resultChannel chan string) {
	for imageID := range imageChannel {
		log.Info("Scanning image: ", imageID)
		// Exec scan with Trivy and Dockle here
		resultChannel <- "Dummy Result"
	}

	close(resultChannel)
	log.Info("Scanner finished")
}
