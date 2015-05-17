package main

import (
	"os"

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

// Configure logging
func setupLogging() {
	// Create a logging backend
	backend := logging.NewLogBackend(os.Stderr, "", 0)

	// Set formatting
	format := logging.MustStringFormatter("%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}")
	backendFormatter := logging.NewBackendFormatter(backend, format)

	// Use the backends
	logging.SetBackend(backendFormatter)
}
