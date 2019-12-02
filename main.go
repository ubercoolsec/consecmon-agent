package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	consecmon "github.com/ubercoolsec/consecmon-agent/lib"
)

var ProgramName string = "consecmon-agent"
var ProgramVersion string = "0.0.1"

func loggerInit() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.JSONFormatter{})
}

func configInit() {
	consecmon.InitConfig()
}

func runModeStandalone() {
	consecmon.EnumRunningContainers(&consecmon.ContainerEngineOpts{
		ScanAllImages: true})
	consecmon.ScanImage()

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
