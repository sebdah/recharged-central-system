package main

import (
	"github.com/op/go-logging"
	"github.com/sebdah/recharged-central-system/config"
)

var log = logging.MustGetLogger("main")

func main() {
	// Configure logging
	setupLogging()

	// Welcome message
	log.Info("Starting re:charged central system")
	log.Info("Environment: %s", config.Env)
}
