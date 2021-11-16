package imgur

import (
	"fmt"

	"github.com/MedzikUser/go-screenshot-imgur/webhook"
)

// Upload Image to Imgur, from file
func UploadFile(filename string) error {
	i, _, err := ClientImgur.UploadImageFromFile(filename, "")
	if err != nil {
		return err
	}

	next(i.Data.IDExt, i.Data.Deletehash)

	return nil
}

// Upload Image to Imgur, from url
func UploadURL(url string) error {
	i, _, err := ClientImgur.UploadImageFromURL(url, "")
	if err != nil {
		return err
	}

	next(i.Data.IDExt, i.Data.Deletehash)

	return nil
}

// Send webhook and print img info
func next(id string, delhash string) {
	url := "https://cdn.magicuser.cf/" + id

	fmt.Println("Link        -", url)
	fmt.Println("ID          -", id)
	fmt.Println("Delete Hash -", delhash)

	webhook.Send(url, delhash)
}
