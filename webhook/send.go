package webhook

import (
	"github.com/MedzikUser/go-screenshot-imgur/config"
	"github.com/gtuk/discordwebhook"
)

func Send(url string) error {
	var username = "CDN"

	msg := discordwebhook.Message{
		Username:  &username,
		AvatarUrl: new(string),
		Embeds:    &[]discordwebhook.Embed{
			{
				Title: &url,
				Image: &discordwebhook.Image{
					Url: &url,
				},
			},
		},
	}

	return discordwebhook.SendMessage(config.Toml.Discord.Webhook, msg)
}
