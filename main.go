package main

import (
	"fmt"
	"net/http"

	goLogging "github.com/op/go-logging"
	"github.com/sebdah/recharged-central-system/config"
	"github.com/sebdah/recharged-central-system/logging"
	"github.com/sebdah/recharged-shared/websockets"
)

var (
	log      = goLogging.MustGetLogger("main")
	WsServer *websockets.Server
)

func main() {
	// Configure logging
	logging.Setup()

	// Welcome message
	log.Info("Starting re:charged central system")
	log.Info("Environment: %s", config.Env)

	// Setup Websockets endpoint
	WsServer = websockets.NewServer()
	http.HandleFunc("/ocpp-2.0j/ws", WsServer.Handler)

	// Start the HTTP server
	log.Info("Starting webserver on port %d", config.Config.GetInt("port"))
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.Config.GetInt("port")), nil)
	if err != nil {
		panic(err)
	}
}
