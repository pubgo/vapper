package config

type Cfg struct {
	App struct {
		WebPort  int    `toml:"web_port"`
		HTTPPort int    `toml:"http_port"`
		RespPort int    `toml:"resp_port"`
		RPCPort  int    `toml:"rpc_port"`
		Name     string `toml:"name"`
		Debug    bool   `toml:"debug"`
		Env      string `toml:"env"`
	} `toml:"app"`
	Db struct {
		Prefix       string `toml:"prefix"`
		Driver       string `toml:"driver"`
		DbName       string `toml:"db_name"`
		URL          string `toml:"url"`
		MaxOpenConns int    `toml:"max_open_conns"`
		MaxIdleConns int    `toml:"max_idle_conns"`
		MaxLefttime  int    `toml:"max_lefttime"`
	} `toml:"db"`
	Log struct {
		Level               string `toml:"level"`
		TimestampFieldName  string `toml:"timestamp_field_name"`
		LevelFieldName      string `toml:"level_field_name"`
		MessageFieldName    string `toml:"message_field_name"`
		ErrorFieldName      string `toml:"error_field_name"`
		CallerFieldName     string `toml:"caller_field_name"`
		ErrorStackFieldName string `toml:"error_stack_field_name"`
		TimeFieldFormat     string `toml:"time_field_format"`
		Writers             []struct {
			Type string `toml:"type"`
			URL  string `toml:"url"`
		} `toml:"writers"`
	} `toml:"log"`
	Templates struct {
		Slack   string `toml:"slack"`
		Email   string `toml:"email"`
		Webhook string `toml:"webhook"`
	} `toml:"templates"`
}
