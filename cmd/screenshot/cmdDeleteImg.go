package main

import (
	"fmt"
	"os"

	"github.com/MedzikUser/go-screenshot-imgur/imgur"
	"github.com/MedzikUser/go-utils/common"
)

func cmdDeleteImg(c cmdOpts) {
	if c.DeleteImage != "" {
		i, _, err := imgur.ClientImgur.DeleteImageUnAuthed(c.DeleteImage)
		if err != nil {
			common.Log.Error(err)

			os.Exit(1)
		}

		if i.Success {
			fmt.Println("Deleted!")

			os.Exit(0)
		} else {
			fmt.Println("Failed deleting image!")

			os.Exit(1)
		}
	}
}
