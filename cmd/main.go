package main

import (
	"nats-streaming-webclient/internal/handlers"
	"syscall"
)
const DefaultPort = "8224"


func main()  {
	port := DefaultPort
	if value, ok := syscall.Getenv("NATS_WEBCLIENT_PORT"); ok {
		port = value
	}

	handlers.Start(port)
}