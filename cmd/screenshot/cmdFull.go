package main

import (
	"os"
	"os/exec"

	"github.com/MedzikUser/go-screenshot-imgur/imgur"
	"github.com/MedzikUser/go-utils/common"
)

func cmdFull(filename string, c cmdOpts) {
	if c.Full {
		_, err := exec.Command("scrot", filename).Output()
		if err != nil {
			common.Log.Error(err)

			os.Exit(1)
		}

		imgur.Upload(filename)

		os.Exit(0)
	}
}
