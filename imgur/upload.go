package imgur

import (
	"fmt"

	"github.com/MedzikUser/go-screenshot-imgur/webhook"
)

func Upload(filename string) error {
	i, _, err := ClientImgur.UploadImageFromFile(filename, "")
	if err != nil {
		return err
	}

	url := "https://cdn.magicuser.cf/" + i.Data.IDExt

	fmt.Println("Link        -", url)
	fmt.Println("ID          -", i.Data.ID)
	fmt.Println("Delete Hash -", i.Data.Deletehash)

	webhook.Send(url, i.Data.Deletehash)

	return nil
}
