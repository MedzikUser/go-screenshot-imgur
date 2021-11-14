package config

import (
	"github.com/BurntSushi/toml"
	"github.com/MedzikUser/go-utils/common"
)

type Config struct {
	Path    string
	Imgur   ConfigImgur
	Discord ConfigDiscord
}

type ConfigImgur struct {
	Id     string
	Secret string
	Album  string
}

type ConfigDiscord struct {
	Webhook string
}

var Toml Config

func TomlParse(configPath string) {
	_, err := toml.DecodeFile(configPath, &Toml)

	if err != nil {
		common.Log.Fatal("decode toml file: ", err)
	}
}
