package main

import (
	"os"
	"os/exec"

	"github.com/MedzikUser/go-screenshot-imgur/imgur"
)

func cmdFull(filename string, c cmdOpts) {
	if c.Full {
		_, err := exec.Command("scrot", filename).Output()
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
