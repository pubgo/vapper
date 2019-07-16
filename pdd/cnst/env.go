package cnst

import "github.com/rs/zerolog"

const (
	DateFormat = "2006-01-02"
	TimeFormat = "2006-01-02 15:04:05"
)

var Env = struct {
	Dev  string
	Stag string
	Prod string
}{
	Dev:  "dev",
	Stag: "stag",
	Prod: "prod",
}

func MatchEnv(e string) bool {
	return Env.Dev == e || Env.Stag == e || Env.Prod == e
}

func MatchLogLevel(e string) bool {
	return zerolog.DebugLevel.String() == e ||
		zerolog.ErrorLevel.String() == e ||
		zerolog.WarnLevel.String() == e ||
		zerolog.FatalLevel.String() == e ||
		zerolog.InfoLevel.String() == e ||
		zerolog.PanicLevel.String() == e
}

const (
	SlackCode        = "slack"
	SlackUrlKey      = "url"
	SlackTemplateKey = "template"
	SlackChannelKey  = "channel"
)

const (
	MailCode        = "mail"
	MailTemplateKey = "template"
	MailServerKey   = "server"
	MailUserKey     = "user"
)

const (
	WebhookCode        = "webhook"
	WebhookTemplateKey = "template"
	WebhookUrlKey      = "url"
)
