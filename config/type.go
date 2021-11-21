package config

// Toml config struct
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

type discord struct {
	EmbedColor string
}
