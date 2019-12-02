package lib

import (
	log "github.com/sirupsen/logrus"
)

func ConnectScanResult(resultChan chan string) {
	for result := range resultChan {
		log.Info("Result received: ", result)
		// Use configured connector to deliver result
	}
}
