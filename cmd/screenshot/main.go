package main

import (
	"os"
	"time"

	"github.com/MedzikUser/go-screenshot-imgur/config"
	"github.com/MedzikUser/go-screenshot-imgur/imgur"
	"github.com/jpillora/opts"
	"github.com/sirupsen/logrus"
)

type cmdOpts struct {
	Select      bool   `opts:"help=take screenshot in select mode"`
	Window      bool   `opts:"help=take screenshot in window mode"`
	Full        bool   `opts:"help=take screenshot in full mode"`
	DeleteImage string `opts:"help=delete image from imgur. Paste deletehash"`
	UploadFile  string `opts:"help=upload image from file to imgur"`
	UploadUrl   string `opts:"help=upload image from url to imgur"`
	Config      string `opts:"help=path to config file"`
	Version     bool   `opts:"help=get version"`
}

var log = &logrus.Logger{
	Out:       os.Stdout,
	Formatter: new(logrus.TextFormatter),
	Hooks:     make(logrus.LevelHooks),
	Level:     logrus.InfoLevel,
}

func main() {
	c := cmdOpts{}

	opts.Parse(&c)

	// --version, -v
	cmdVersion(c)

	if c.Config == "" {
		dirname, err := os.UserHomeDir()
		if err != nil {
			log.Error(err)

			c.Config = dirname + ".config.toml"
		}

		c.Config = dirname + "/.config/screenshots-imgur/config.toml"
	}

	err := config.TomlParse(c.Config)
	if err != nil {
		log.Fatal("parse toml:", err)
	}

	imgur.CreateClient()

	// --delete-image, -d
	cmdDeleteImg(c)

	// --upload-file path/to/file
	cmdUploadFile(c)
	// --upload-url https://url/img
	cmdUploadUrl(c)

	filename := config.Toml.Path + "/" + time.Now().Format("2006-01-02T15:04:05-0700") + config.FileExt

	// --select, -s
	cmdSelect(filename, c)
	// --window, -w
	cmdWindow(filename, c)
	// --full, -f
	cmdFull(filename, c)

	log.Info("Use with --help or -h to get more informations.")
}
