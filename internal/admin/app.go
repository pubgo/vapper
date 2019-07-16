package admin

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/logger"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/pubgo/vapper/internal/config"
	"github.com/rs/zerolog"
	"net/http"
	"os"
)

func App() *gin.Engine {
	cfg := config.Default()
	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(cors.Default())

	// Custom logger
	subLog := zerolog.New(os.Stdout).With().
		Str("app", cfg.Cfg.App.Name).
		Logger()
	r.Use(logger.SetLogger(logger.Config{
		Logger: &subLog,
		UTC:    true,
		//SkipPath:       []string{"/test"},
		//SkipPathRegexp: regexp.MustCompile(`^/regexp\d*`),
	}))

	r.Use(sessions.Sessions("dr_session", cookie.NewStore([]byte("secret"))))
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	// app := r.Group("/login")// 用户登录

	app := r.Group("/db2rest") // 需要auth校验，不然进不来
	app.GET("/", func(ctx *gin.Context) {
		var rs []gin.H
		for _, rt := range r.Routes() {
			rs = append(rs, gin.H{
				"method":  rt.Method,
				"path":    rt.Path,
				"handler": rt.Handler,
			})
		}
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"routes":  rs,
			"version": version.Version,
			"buildV":  version.BuildV,
			"commitV": version.CommitV,
			"config":  cfg,
		})
	})

	app.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
		return
	})

	return r
}
