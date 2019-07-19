package config

import (
	"github.com/pubgo/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func (t *config) InitLog() {

	zerolog.TimestampFieldName = "time"
	zerolog.LevelFieldName = "level"
	zerolog.MessageFieldName = "msg"

	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	if t.Debug {
		t.LogLevel = zerolog.DebugLevel.String()
	}

	if t.LogLevel != "" {
		_lv, err := zerolog.ParseLevel(t.LogLevel)
		errors.Wrap(err, "log level parse error")
		zerolog.SetGlobalLevel(_lv)
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, NoColor: false}).
		With().
		Caller().
		Str("pkg", "vapper").
		Logger()

}
