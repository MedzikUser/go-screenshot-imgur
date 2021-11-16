package main

import (
	"os"
	"os/exec"

	"github.com/MedzikUser/go-screenshot-imgur/imgur"
)

// --window, -w
func cmdWindow(filename string, c cmdOpts) {
	if c.Window {
		_, err := exec.Command("scrot", "-u", filename).Output()
		if err != nil {
			log.Fatal(err)
		}

		err = imgur.UploadFile(filename)
		if err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
	}
}
