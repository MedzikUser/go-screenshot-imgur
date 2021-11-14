package main

import (
	"time"

	"github.com/MedzikUser/go-screenshot-imgur/config"
	"github.com/MedzikUser/go-screenshot-imgur/imgur"
	"github.com/jpillora/opts"
)

type cmdOpts struct {
	Select  bool   `opts:"help=take screenshot in select mode"`
	Window  bool   `opts:"help=take screenshot in window mode"`
	Full    bool   `opts:"help=take screenshot in full mode"`
	Config  string `opts:"help=path to config file"`
	Version bool   `opts:"help=get version"`
}

func main() {
	c := cmdOpts{}

	opts.Parse(&c)

	cmdVersion(c)

	if c.Config == "" {
		c.Config = "config.toml"
	}

	config.TomlParse(c.Config)

	imgur.CreateClient()

	filename := config.Toml.Path + "/" + time.Now().Format("2006-01-02T15:04:05-0700") + config.FileExt

	cmdSelect(filename, c)
}
