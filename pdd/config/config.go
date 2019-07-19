package config

type Cfg struct {
	App struct {
		Port    int    `toml:"port"`
		Name    string `toml:"name"`
		Debug   bool   `toml:"debug"`
		Env     string `toml:"env"`
		SiteURL string `toml:"site_url"`
	} `toml:"app"`
	Services []struct {
		Name      string `toml:"name"`
		ServerURL string `toml:"server_url"`
		ProxyURI  string `toml:"proxy_uri"`
	} `toml:"services"`
	Pkg struct {
		URL   string `toml:"url"`
		Token string `toml:"token"`
	} `toml:"pkg"`
}
