package imgur

import (
	"net/http"

	"github.com/MedzikUser/go-imgur"
	"github.com/MedzikUser/go-screenshot-imgur/config"
)

var ClientImgur = imgur.Client{}

func CreateClient() {
	ClientImgur.HTTPClient = new(http.Client)
	ClientImgur.Imgur.ClientID = config.Toml.Imgur.Id
	ClientImgur.Imgur.AccessToken = config.Toml.Imgur.Secret
}
