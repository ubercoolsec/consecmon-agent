package lib

import (
	log "github.com/sirupsen/logrus"
	scanners "github.com/ubercoolsec/consecmon-agent/modules/scanners"
)

type ScannerOpt struct {
	ScanWithTrivy  bool
	ScanWithDockle bool
}

func ScanImage(opts *ScannerOpt, imageChannel chan string, resultChannel chan string) {
	for imageID := range imageChannel {
		log.Info("Scanning image: ", imageID)

		resultChannel <- scanners.RunTrivyScanner(imageID)
		resultChannel <- scanners.RunDockleScanner(imageID)
	}

	close(resultChannel)
	log.Info("Scanner finished")
}
