package routesv1

import (
	"fmt"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func MercadoPago(router *fasthttprouter.Router) {
	router.GET("/mp", func(ctx *fasthttp.RequestCtx) {
		body := ctx.Request.Body()

		fmt.Println("DATA: ", string(body))

		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBodyString("OK")
	})
}
