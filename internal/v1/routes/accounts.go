package routesv1

import (
	"github.com/buaazp/fasthttprouter"
	accountservicesv1 "github.com/dechevarrieta1/hrhelpers/internal/v1/services/v1/accounts"
)

func AccountRoutes(router *fasthttprouter.Router) {
	router.POST("/accounts/create", accountservicesv1.CreateAccount)
	router.POST("/accounts/login", accountservicesv1.LoginAccount)
}
