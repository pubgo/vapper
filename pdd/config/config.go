package config

type Cfg struct {
	App struct {
		Port      int    `toml:"port"`
		Name      string `toml:"name"`
		Debug     bool   `toml:"debug"`
		Env       string `toml:"env"`
		ServerURL string `toml:"server_url"`
		SiteURL   string `toml:"site_url"`
		URLPrefix string `toml:"url_prefix"`
	} `toml:"app"`
}
