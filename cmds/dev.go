package cmds

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pubgo/errors"
	"github.com/pubgo/vapper/internal/jsbuild"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"mime"
	"net/http"
	"os"
	"path"
	"strings"
	"text/template"
)

var _index = template.Must(template.New("index").Parse(
	`<html>
	<head>
		<meta charset="utf-8">
		<link href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm" crossorigin="anonymous">
		<script src="https://code.jquery.com/jquery-3.2.1.slim.min.js" integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN" crossorigin="anonymous"></script>
		<script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/js/bootstrap.min.js" integrity="sha384-JZR6Spejh4U02d8jOt6vLEHfe/JQGiRRSQQxSfFWpi1MquVdAyjUar5+76PVCmYl" crossorigin="anonymous"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/ace/1.3.3/ace.js"></script>
	</head>
	<body id="wrapper" style="margin: 0;">
		<div id="progress-holder" style="width: 100%; padding: 25%;">
			<div class="progress">
				<div id="progress-bar" class="progress-bar" role="progressbar" style="width: 0%" aria-valuenow="0" aria-valuemin="0" aria-valuemax="100"></div>
			</div>
		</div>
		<script>
			window.jsgoProgress = function(count, total) {
				var value = (count * 100.0) / (total * 1.0);
				var bar = document.getElementById("progress-bar");
				bar.style.width = value+"%";
				bar.setAttribute('aria-valuenow', value);
				if (count === total) {
					document.getElementById("progress-holder").style.display = "none";
				}
			}
		</script>
		<script src="{{ .Script }}"></script>
	</body>
</html>`,
))

var (
	_build = jsbuild.Default()
)

func initDevCmd(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().BoolVarP(&_build.Options.Verbose, "verbose", "v", false, "print the names of packages as they are compiled")
	cmd.Flags().BoolVarP(&_build.Options.Quiet, "quiet", "q", false, "suppress non-fatal warnings")
	cmd.Flags().BoolVarP(&_build.Options.Minify, "minify", "m", true, "minify generated code")
	cmd.Flags().BoolVar(&_build.Options.Color, "color", terminal.IsTerminal(int(os.Stderr.Fd())) && os.Getenv("TERM") != "dumb", "colored output")
	cmd.Flags().StringVar(&_build.Tags, "tags", "", "a list of build tags to consider satisfied during the build")
	cmd.Flags().BoolVar(&_build.Options.MapToLocalDisk, "localmap", false, "use local paths for sourcemap")
	cmd.Flags().StringVarP(&_build.Addr, "http", "", ":8080", "HTTP bind address to serve")
	cmd.Flags().BoolVarP(&_build.OnlyHash, "hash", "", true, "only hash path")
	cmd.Flags().BoolVarP(&options.Watch, "watch", "w", true, "watch for changes to the source files")
	return cmd
}

func jsHandler(w http.ResponseWriter, r *http.Request) {
	_m := _build.MainUrl()
	w.Write(_m.Content)
}

func pkgHandler(w http.ResponseWriter, r *http.Request) {
	us := strings.Split(strings.ReplaceAll(r.URL.Path, "/pkg/", ""), ".")
	_pkg := _build.GetByHash(strings.Join(us[:len(us)-2],"."))
	if _pkg == nil {
		fmt.Println(us)
		fmt.Fprintf(w, "not found %s", r.URL.Path)
		return
	}

	w.Header().Set("Cache-Control", "public,max-age=31536000,immutable")
	//fmt.Println(mime.TypeByExtension(path.Ext(r.URL.Path)))

	w.Header().Set("Content-Type", mime.TypeByExtension(path.Ext(r.URL.Path)))
	w.Write(_pkg.Content)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	_m := _build.MainUrl()
	buf := &bytes.Buffer{}
	errors.Wrap(_index.Execute(buf, struct {
		Script string
	}{Script: fmt.Sprintf("/js/%s.%s.js", _m.Path, _m.Hash)}), "mainTemplateMinified error")
	w.Write(buf.Bytes())
}

func init() {
	rootCmd.AddCommand(initDevCmd(&cobra.Command{
		Use:   "dev",
		Short: "dev",
		Run: func(cmd *cobra.Command, args []string) {
			defer errors.Assert()

			_curDir, err := os.Getwd()
			errors.Panic(err)
			_build.RootPath = _curDir

			r := mux.NewRouter()
			r.HandleFunc("/", indexHandler)
			r.PathPrefix("/pkg").HandlerFunc(pkgHandler)
			r.PathPrefix("/js").HandlerFunc(jsHandler)
			r.PathPrefix("/services").HandlerFunc(pkgHandler)
			http.Handle("/", r)

			go _build.Build()
			fmt.Println("ok")
			errors.Panic(http.ListenAndServe(":8080", nil))
		},
	}))
}

/*
func ServeStatic(name string, w http.ResponseWriter, req *http.Request, mimeType string) error {
	var file billy.File
	var err error
	file, err = assets.Assets.Open(name)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, req)
			return nil
		}
		http.Error(w, fmt.Sprintf("error opening %s", name), 500)
		return nil
	}
	defer file.Close()

	w.Header().Set("Cache-Control", "public,max-age=31536000,immutable")
	if mimeType == "" {
		w.Header().Set("Content-Type", mime.TypeByExtension(path.Ext(req.URL.Path)))
	} else {
		w.Header().Set("Content-Type", mimeType)
	}

	_, noCompress := file.(httpgzip.NotWorthGzipCompressing)
	gzb, isGzb := file.(httpgzip.GzipByter)

	if isGzb && !noCompress && strings.Contains(req.Header.Get("Accept-Encoding"), "gzip") {
		w.Header().Set("Content-Encoding", "gzip")
		if err := WriteWithTimeout(w, gzb.GzipBytes()); err != nil {
			http.Error(w, fmt.Sprintf("error streaming gzipped %s", name), 500)
			return err
		}
	} else {
		if err := StreamWithTimeout(w, file); err != nil {
			http.Error(w, fmt.Sprintf("error streaming %s", name), 500)
			return err
		}
	}
	return nil

}

func StreamWithTimeout(w io.Writer, r io.Reader) error {
	c := make(chan error, 1)
	go func() {
		_, err := io.Copy(w, r)
		c <- err
	}()
	select {
	case err := <-c:
		if err != nil {
			return err
		}
		return nil
	case <-time.After(config.WriteTimeout):
		return errors.New("timeout")
	}
}

func WriteWithTimeout(w io.Writer, b []byte) error {
	return StreamWithTimeout(w, bytes.NewBuffer(b))
}

type Pather interface {
	Path() string
}

type Handler struct {
	Cache      *cache.Cache
	Fileserver services.Fileserver
	Database   services.Database
	Waitgroup  *sync.WaitGroup
	Queue      *queue.Queue
	mux        *http.ServeMux
	shutdown   chan struct{}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func (h *Handler) storeError(ctx context.Context, err error, req *http.Request) {

	if err == queue.TooManyItemsQueued {
		// If the server is getting flooded by a DOS, this will prevent database flooding
		return
	}

	// ignore errors when logging an error
	store.StoreError(ctx, h.Database, store.Error{
		Time:  time.Now(),
		Error: err.Error(),
		Ip:    req.Header.Get("X-Forwarded-For"),
	})

}

func (h *Handler) IconHandler(w http.ResponseWriter, req *http.Request) {
	if err := ServeStatic(req.URL.Path, w, req, "image/x-icon"); err != nil {
		http.Error(w, "error serving static file", 500)
	}
}

func (h *Handler) CssHandler(w http.ResponseWriter, req *http.Request) {
	if err := ServeStatic(req.URL.Path, w, req, "text/css"); err != nil {
		http.Error(w, "error serving static file", 500)
	}
}

func (h *Handler) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}
*/
