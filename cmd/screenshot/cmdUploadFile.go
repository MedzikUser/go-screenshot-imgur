package main

import (
	"os"

	"github.com/MedzikUser/go-screenshot-imgur/imgur"
)

func cmdUploadFile(c cmdOpts) {
	if c.UploadFile != "" {
		err := imgur.UploadFile(c.UploadFile)
		if err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
	}
}
