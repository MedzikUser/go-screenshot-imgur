package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/MedzikUser/go-screenshot-imgur/imgur"
	"github.com/MedzikUser/go-screenshot-imgur/webhook"
	"github.com/MedzikUser/go-utils/common"
)

func cmdSelect(filename string, c cmdOpts) {
	if c.Select {
		_, err := exec.Command("scrot", "-s", filename).Output()
		if err != nil {
			common.Log.Error(err)
		}

		i, _, err := imgur.ClientImgur.UploadImageFromFile(filename, "")
		if err != nil {
			common.Log.Error(err)

			os.Exit(1)
		}

		url := "https://cdn.magicuser.cf/" + i.Data.IDExt

		fmt.Println("Link        -", url)
		fmt.Println("ID          -", i.Data.ID)
		fmt.Println("Delete hash -", i.Data.Deletehash)

		webhook.Send(url)
	}
}
