package main

import (
	"os"

	"github.com/MedzikUser/go-screenshot-imgur/imgur"
)

// --upload-file path/to/file
func cmdUploadFile(c cmdOpts) {
	if c.UploadFile != "" {
		err := imgur.UploadFile(c.UploadFile)
		if err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
	}
}
