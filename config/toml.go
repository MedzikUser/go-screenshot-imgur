package config

import (
	"strings"

	"github.com/BurntSushi/toml"
)

var Toml Config

// Parse Toml config
func TomlParse(configPath string) error {
	_, err := toml.DecodeFile(configPath, &Toml)

	if err != nil && !strings.Contains(err.Error(), "no such file or directory") {
		return err
	}

	return nil
}
