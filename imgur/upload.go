package imgur

import (
	"fmt"
	"os"

	"github.com/MedzikUser/go-screenshot-imgur/webhook"
	"github.com/MedzikUser/go-utils/common"
)

func Upload(filename string) {
	i, _, err := ClientImgur.UploadImageFromFile(filename, "")
	if err != nil {
		common.Log.Error(err)

		os.Exit(1)
	}

	url := "https://cdn.magicuser.cf/" + i.Data.IDExt

	fmt.Println("Link        -", url)
	fmt.Println("ID          -", i.Data.ID)
	fmt.Println("Delete Hash -", i.Data.Deletehash)

	webhook.Send(url, i.Data.Deletehash)
}
