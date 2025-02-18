package routesv1

import (
	"github.com/buaazp/fasthttprouter"
	healthservicesv1 "github.com/dechevarrieta1/hrhelpers/internal/v1/services/v1/health"
)

func HealthCheck(router *fasthttprouter.Router) {
	router.GET("/health", healthservicesv1.HealthCheckHandler)
}
