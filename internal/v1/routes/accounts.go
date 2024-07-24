package routesv1

import (
	"github.com/buaazp/fasthttprouter"
	accountservicesv1 "github.com/dechevarrieta1/obra-3/internal/v1/services/v1/accounts"
)

func TestRoutes(router *fasthttprouter.Router) {
	router.GET("/accounts/create", accountservicesv1.CreateAccount)
	router.POST("/accounts/login", accountservicesv1.LoginAccount)
}
