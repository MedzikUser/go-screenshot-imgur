package main

import (
	"os"
	"os/exec"

	"github.com/MedzikUser/go-screenshot-imgur/imgur"
	"github.com/MedzikUser/go-utils/common"
)

func cmdSelect(filename string, c cmdOpts) {
	if c.Select {
		_, err := exec.Command("scrot", "-s", filename).Output()
		if err != nil {
			common.Log.Error(err)

			os.Exit(1)
		}

		imgur.Upload(filename)

		os.Exit(0)
	}
}
