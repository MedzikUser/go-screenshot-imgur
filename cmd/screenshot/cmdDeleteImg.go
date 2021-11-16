package main

import (
	"os"

	"github.com/MedzikUser/go-screenshot-imgur/imgur"
)

// --delete-image, -d
func cmdDeleteImg(c cmdOpts) {
	if c.DeleteImage != "" {
		i, _, err := imgur.ClientImgur.DeleteImageUnAuthed(c.DeleteImage)
		if err != nil {
			log.Fatal(err)
		}

		if i.Success {
			log.Info("Deleted!")

			os.Exit(0)
		} else {
			log.Fatal("Failed deleting image!")
		}
	}
}
