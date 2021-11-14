package imgur

import (
	"net/http"

	"github.com/MedzikUser/go-imgur"
	"github.com/MedzikUser/go-screenshot-imgur/config"
)

var ClientImgur = new(imgur.Client)

func init() {
	ClientImgur.HTTPClient = new(http.Client)
	ClientImgur.ImgurClientID = config.Toml.Imgur.Id
}
