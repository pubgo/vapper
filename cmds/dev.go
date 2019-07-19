package cmds

import (
	"github.com/pubgo/errors"
	"github.com/pubgo/vapper/internal/jsbuild"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

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

func init() {
	rootCmd.AddCommand(initDevCmd(&cobra.Command{
		Use:   "dev",
		Short: "dev",
		Run: func(cmd *cobra.Command, args []string) {
			defer errors.Assert()

			_curDir, err := os.Getwd()
			errors.Panic(err)

			_build.RootPath = _curDir
			_build.Build()
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
