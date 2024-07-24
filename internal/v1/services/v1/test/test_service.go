package testservicesv1

import (
	"math/rand"

	httputilsv1 "github.com/dechevarrieta1/obra-3/pkg/http"
	"github.com/valyala/fasthttp"
)

func TestService(ctx *fasthttp.RequestCtx) {

	number := rand.Float64()
	httputilsv1.ResponseHandlers(ctx, number, nil, fasthttp.StatusOK, "Random number generated succesfully")
}
