package lib

import (
	"bytes"

	log "github.com/sirupsen/logrus"
)

type ConnectorParam struct {
	ScannerName    string
	ScannerVersion string
	ScanResult     bytes.Buffer
}

func ConnectScanResult(resultChan chan string) {
	for result := range resultChan {
		log.Info("Result received: ", result)
		// Use configured connector to deliver result
	}
}
