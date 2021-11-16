package webhook

import (
	"fmt"

	"github.com/MedzikUser/go-screenshot-imgur/config"
	"github.com/gtuk/discordwebhook"
)

// Send Webhook to Discord Channel
func Send(url string, deletehash string) error {
	var username = "CDN"
	var description = fmt.Sprintf("Delete Hash: ||%s||", deletehash)

	msg := discordwebhook.Message{
		Username: &username,
		Embeds: &[]discordwebhook.Embed{
			{
				Title:       &url,
				Description: &description,
				Color:       &config.Discord.EmbedColor,
				Image:       &discordwebhook.Image{Url: &url},
			},
		},
	}

	return discordwebhook.SendMessage(config.Toml.Discord.Webhook, msg)
}
