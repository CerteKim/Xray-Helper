package xray

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func StartHandler(ctx *fasthttp.RequestCtx) {
	ret := Start()
	fmt.Fprintf(ctx, ret+"\n")
}

func StopHandler(ctx *fasthttp.RequestCtx) {
	ret := fmt.Sprintln(Stop())
	fmt.Fprintf(ctx, ret+"\n")
}
