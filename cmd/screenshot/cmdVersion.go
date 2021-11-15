package main

import (
	"os"
	"runtime"

	"github.com/MedzikUser/go-screenshot-imgur/config"
)

func cmdVersion(c cmdOpts) {
	if c.Version {
		log.Infof("Version: %s | %s\n", config.Version, runtime.Version())

		os.Exit(0)
	}
}
