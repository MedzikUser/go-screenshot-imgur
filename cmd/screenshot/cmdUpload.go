package main

import (
	"os"

	"github.com/MedzikUser/go-screenshot-imgur/imgur"
)

func cmdUpload(c cmdOpts) {
	if c.Upload != "" {
		err := imgur.Upload(c.Upload)
		if err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
	}
}
