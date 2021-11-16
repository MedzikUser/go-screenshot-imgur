package main

import (
	"os"

	"github.com/MedzikUser/go-screenshot-imgur/imgur"
)

// --upload-url https://url/img
func cmdUploadUrl(c cmdOpts) {
	if c.UploadUrl != "" {
		err := imgur.UploadURL(c.UploadUrl)
		if err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
	}
}
