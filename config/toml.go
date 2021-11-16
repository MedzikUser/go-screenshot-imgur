package config

import (
	"strings"

	"github.com/BurntSushi/toml"
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

func TomlParse(configPath string) error {
	_, err := toml.DecodeFile(configPath, &Toml)

	if err != nil && !strings.Contains(err.Error(), "no such file or directory") {
		return err
	}

	return nil
}
