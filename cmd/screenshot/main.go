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
	Upload      string `opts:"help=upload specified file to imgur"`
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

	cmdVersion(c)

	if c.Config == "" {
		c.Config = "config.toml"
	}

	err := config.TomlParse(c.Config)
	if err != nil {
		log.Fatal("parse toml:", err)
	}

	imgur.CreateClient()

	cmdDeleteImg(c)
	cmdUpload(c)

	filename := config.Toml.Path + "/" + time.Now().Format("2006-01-02T15:04:05-0700") + config.FileExt

	cmdSelect(filename, c)
	cmdWindow(filename, c)
	cmdFull(filename, c)

	log.Info("Use with --help or -h to get more informations.")
}
