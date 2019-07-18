package config

import (
	"github.com/pubgo/errors"
	"github.com/pubgo/vapper/pdd/cnst"
	config2 "github.com/pubgo/vapper/pdd/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
	"sync"
)

type config struct {
	Cfg      *config2.Cfg
	Debug    bool
	Env      string
	LogLevel string
}

func (t *config) InitCfg() {
	log.Debug().Msg("init config")

	t.Debug = t.Cfg.App.Debug
	if _d, ok := os.LookupEnv("debug"); ok {
		t.Debug = _d == "true" || _d == "1" || _d == "ok"
	}
	if _d := viper.Get("debug"); _d != nil {
		t.Debug = viper.GetBool("debug")
	}

	t.Env = t.Cfg.App.Env
	if _e, ok := os.LookupEnv("env"); ok {
		t.Env = _e
		errors.T(!cnst.MatchEnv(t.Env), "the env value is not match,value(%s)", _e)
	}
	if _d := viper.Get("env"); _d != nil {
		t.Env = viper.GetString("env")
		errors.T(!cnst.MatchEnv(t.Env), "the env value is not match,value(%s)", t.Env)
	}

	t.LogLevel = zerolog.DebugLevel.String()
	if _l, ok := os.LookupEnv("log_level"); ok {
		t.LogLevel = _l
		errors.T(!cnst.MatchLogLevel(t.LogLevel), "the env value is not match,value(%s)", _l)
	}
	if _d := viper.Get("ll"); _d != nil {
		t.LogLevel = viper.GetString("ll")
		errors.T(!cnst.MatchLogLevel(t.LogLevel), "the env value is not match,value(%s)", t.LogLevel)
	}
	log.Debug().Msg("init config ok")
}

func (t *config) Init() {

	// cfg 初始化
	t.InitCfg()

	// log初始化
	t.InitLog()

	// 数据库初始化

}

func (t *config) Parse() {
	log.Debug().Msg("parse config")
	errors.Wrap(viper.Unmarshal(t.Cfg), "config parse error")
	log.Debug().Msg("parse config ok")
}

var cfg *config
var once sync.Once

func Default() *config {
	once.Do(func() {
		cfg = &config{Cfg: new(config2.Cfg)}
		cfg.Env = cnst.Env.Dev
		cfg.LogLevel = zerolog.DebugLevel.String()
		cfg.Debug = true
	})
	return cfg
}
