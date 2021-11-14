package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/MedzikUser/go-screenshot-imgur/config"
	"github.com/MedzikUser/go-screenshot-imgur/imgur"
	"github.com/MedzikUser/go-screenshot-imgur/webhook"
	"github.com/MedzikUser/go-utils/common"
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

	if c.Version {
		fmt.Printf("Version: %s | %s\n", config.Version, runtime.Version())

		os.Exit(0)
	}

	if c.Config == "" {
		c.Config = "config.toml"
	}

	config.TomlParse(c.Config)

	filename := config.Toml.Path + "/" + time.Now().Format("2006-01-02T15:04:05-0700") + config.FileExt

	if c.Select {
		_, err := exec.Command("scrot", "-s", filename).Output()
		if err != nil {
			common.Log.Error(err)
		}

		i, _, err := imgur.ClientImgur.UploadImageFromFile(filename, "")
		if err != nil {
			common.Log.Error(err)

			os.Exit(1)
		}

		url := "https://cdn.magicuser.cf/" + i.Data.IDExt

		webhook.Send(url)
	}

	/* out, err := exec.Command("ls").Output()
	if err != nil {
		common.Log.Error(err)
	}

	common.Log.Info(string(out[:])) */
}
