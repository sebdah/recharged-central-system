package main

import (
	"fmt"
	"net/http"

	goLogging "github.com/op/go-logging"
	"github.com/sebdah/recharged-central-system/config"
	"github.com/sebdah/recharged-central-system/logging"
	"github.com/sebdah/recharged-shared/rpc"
	"github.com/sebdah/recharged-shared/websockets"
)

var (
	log      goLogging.Logger
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

	// Fire up the websockets communicator
	go websocketCommunicator()

	// Configure handlers
	http.HandleFunc("/ocpp-2.0j/ws", WsServer.Handler)

	// Start the HTTP server
	log.Info("Starting webserver on port %d", config.Config.GetInt("port"))
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.Config.GetInt("port")), nil)
	if err != nil {
		panic(err)
	}
}

// Communicator for websockets, reading and sending messages
func websocketCommunicator() {
	var message string
	log.Info("Starting the websocket communicator")

	for {
		message = <-WsServer.ReadMessage
		log.Debug("RECV: %s", message)
		messageType, err := rpc.ParseMessage(message)
		if err != nil {
			log.Notice("The incoming message does not match the RPC protocol")
			continue
		}

		switch {
		case messageType == 2:
			log.Info("RECV: %s", message)
		case messageType == 3:
			log.Info("RECV: %s", message)
		case messageType == 4:
			log.Info("RECV: %s", message)
		default:
			log.Error("RPC call not supported")
			continue
		}
	}
}
