package scanners

import (
	log "github.com/sirupsen/logrus"
)

// Execute Trivy Scan using imageID and return result JSON
// We can start with handling JSON as a string but its not memory optimized
// We can probably use temporary files to output the result
func RunTrivyScanner(imageID string) string {
	log.Info("Running Trivy scanner on: ", imageID)

	return "Sample Result JSON from Trivy"
}
