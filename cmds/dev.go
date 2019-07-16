package cmds

import (
	"github.com/pubgo/errors"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ReverseProxy(target string) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		_url, err := url.Parse(target)
		errors.Panic(err)
		httputil.NewSingleHostReverseProxy(_url).ServeHTTP(writer, request)
	}
}

func initWebCmd(cmd *cobra.Command) *cobra.Command {
	return cmd
}

// WebCmd ...
func init() {

	rootCmd.AddCommand(initWebCmd(&cobra.Command{
		Use:   "web",
		Short: "web",
		Run: func(cmd *cobra.Command, args []string) {
			defer errors.Assert()

			dt, err := ioutil.ReadFile("static/index.html")
			errors.Panic(err)
			
			http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				_, err := w.Write(dt)
				errors.Panic(err)
			})

			http.HandleFunc("/frontend.js", ReverseProxy("http://localhost:3000"))
			errors.Panic(http.ListenAndServe(":8080", nil))
		},
	}))
}
