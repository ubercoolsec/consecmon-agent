package main

import (
	"flag"
	"os"

	log "github.com/sirupsen/logrus"
	consecmon "github.com/ubercoolsec/consecmon-agent/lib"
)

var scanAllImage = flag.Bool("scan-all-image", false, "Scan all images available locally")
var apiURL = flag.String("api-url", "", "HTTP connector URL")
var apiSecret = flag.String("api-secret", "", "API secret to include in Authorization header")

var ProgramName string = "consecmon-agent"
var ProgramVersion string = "0.0.1"

func loggerInit() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.JSONFormatter{})
}

func configInit() {
	flag.Parse()
}

func runModeStandalone() {
	scannerChannel := make(chan string)

	go consecmon.EnumRunningContainers(&consecmon.ContainerEngineOpts{
		ScanAllImages: *scanAllImage}, scannerChannel)
	consecmon.ScanImage(scannerChannel)

	// Start delivery module and setup queue
	// Enumerate docker containers
	// Run modules on docker image for each container
	// Queue JSON output from module for delivery to server
}

func runOnce() {
	runModeStandalone()
}

func main() {
	loggerInit()
	configInit()

	log.Info("Running ", ProgramName, " Version: ", ProgramVersion)
	runOnce()
}
