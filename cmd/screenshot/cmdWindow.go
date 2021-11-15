package main

import (
	"os"
	"os/exec"

	"github.com/MedzikUser/go-screenshot-imgur/imgur"
)

func cmdWindow(filename string, c cmdOpts) {
	if c.Window {
		_, err := exec.Command("scrot", "-s", "-b", filename).Output()
		if err != nil {
			log.Fatal(err)
		}

		err = imgur.Upload(filename)
		if err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
	}
}
