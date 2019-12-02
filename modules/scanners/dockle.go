package scanners

import (
	log "github.com/sirupsen/logrus"
)

func RunDockleScanner(imageID string) string {
	log.Info("Running Dockle scanner on: ", imageID)

	return "Sample Result from Dockle"
}
