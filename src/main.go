package main

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"strings"
)

type Config struct {
	ServiceName string
	StaticFiles string
}

func main() {
	tServiceName := "TouchingWASM"
	tPrefix := "/" + strings.Trim(tServiceName, "/")
	tConfig := &Config{
		ServiceName: tPrefix,
		StaticFiles: "src/webapp/public",
	}

	tRouter := chi.NewRouter()

	tRouter.Mount(tPrefix, http.StripPrefix(tPrefix, http.FileServer(http.Dir(tConfig.StaticFiles))))
	if tPrefix != "/" {
		tRouter.Get(tPrefix, func(aWriter http.ResponseWriter, aRequest *http.Request) {
			http.Redirect(aWriter, aRequest, tPrefix+"/", http.StatusMovedPermanently)
		})
	}

	tRouter.Mount("/api/v1", tConfig.RestAPI())

	tErr := http.ListenAndServe(":8888", tRouter)
	if tErr != nil {
		log.Fatal("[ERROR] failed to listen and serve")
	}
}

func (aConfig *Config) RestAPI() http.Handler {
	tRouter := chi.NewRouter()
	return tRouter
}
