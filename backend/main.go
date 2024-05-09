package main

import (
	"janstupica/StickyNotes/config"
)

func main() {
	status := config.Init()

	if !status {
		panic("Config failed to initialize.")
	}
}
