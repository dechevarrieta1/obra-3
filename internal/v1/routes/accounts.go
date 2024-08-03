package routesv1

import (
	"github.com/buaazp/fasthttprouter"
	middlewaresv1 "github.com/dechevarrieta1/obra-3/internal/v1/middlewares"
	accountservicesv1 "github.com/dechevarrieta1/obra-3/internal/v1/services/v1/accounts"
)

func AccountRoutes(router *fasthttprouter.Router) {
	router.POST("/accounts/create", accountservicesv1.CreateAccount)
	router.POST("/accounts/login", middlewaresv1.AuthMiddleware(accountservicesv1.LoginAccount))
}
