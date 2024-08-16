package main

import (
	"os"
	"strconv"
)

func debugModeActive() bool {

	debugModeActive := false

	if value := os.Getenv("DEBUG_MODE"); value != "" {
		if parsedBool, err := strconv.ParseBool(value); err == nil {
			debugModeActive = parsedBool
		} else {
			panic(err)
		}
	}

	return debugModeActive
}

func modoDesktopActivo() bool {
	return os.Getenv("DESKTOP") == "ENABLED"
}
