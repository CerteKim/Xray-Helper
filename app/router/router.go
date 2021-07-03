package router

import (
	"fmt"
	"xrayd/app/service/xray"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func Router(ctx *fasthttp.RequestCtx) *router.Router {
	r := router.New()
	r.GET("/version", versionHandler)
	r.GET("/api/v1/start", xray.StartHandler)
	r.GET("/api/v1/stop", xray.StopHandler)
	return r
}

func versionHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "xrayd version 0.0.1")
}
