package main

import (
	"janstupica/StickyNotes/config"
	srvr "janstupica/StickyNotes/internal/app/server"
)

func main() {
	status := config.Init()

	if !status {
		panic("Config failed to initialize.")
	}

	server := srvr.New()

	if err := server.Run(); err != nil {
		panic("Server failed to run. " + err.Error())
	}
}
