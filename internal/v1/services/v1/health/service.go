package healthservicesv1

import "github.com/valyala/fasthttp"

func HealthCheckHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBodyString("OK")
}
