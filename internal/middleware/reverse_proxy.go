package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
)

func ReverseProxy(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		(&httputil.ReverseProxy{
			Director: func(req *http.Request) {
				r := c.Request

				req = r
				req.URL.Scheme = "http"
				req.URL.Host = target
				req.Header["my-header"] = []string{r.Header.Get("my-header")}
				delete(req.Header, "My-Header")
			},
		}).ServeHTTP(c.Writer, c.Request)
	}
}
