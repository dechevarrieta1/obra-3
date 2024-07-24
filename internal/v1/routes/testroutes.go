package routesv1

import (
	"github.com/buaazp/fasthttprouter"
	testservicesv1 "github.com/dechevarrieta1/obra-3/internal/v1/services/v1/test"
)

func TestRoutes(router *fasthttprouter.Router) {
	//TODO IMPLEMENT MIDDLEWARES
	router.GET("/test", testservicesv1.TestService)
}
