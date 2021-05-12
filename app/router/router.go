package router

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
)

var Router = httprouter.New()

func Start() {

	Router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "*")
		}
		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})

	if viper.GetBool("xrayd.pprof") {
		Router.Handler(http.MethodGet, "/debug/pprof/*item", http.DefaultServeMux)
	}

	if static := viper.GetString("xrayd.static"); static != "" {
		Router.NotFound = http.FileServer(http.Dir(static))
	} else {
		Router.GET("/", DefaultHandler)
	}
}

func DefaultHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "%s\n", "xrayd version 0.0.1")
}
