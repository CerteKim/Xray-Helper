package router

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func versionHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	fmt.Fprintf(ctx, "xrayd version 0.0.1")
}

func Router(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/version":
		versionHandler(ctx)
	default:
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
	}
}
