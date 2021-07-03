package xray

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func StartHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, Start()+"\n")
}

func StopHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, fmt.Sprintln(Stop())+"\n")
}
