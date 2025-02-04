package routesv1

import (
	"github.com/buaazp/fasthttprouter"
)

func InitRoutes() *fasthttprouter.Router {
	router := fasthttprouter.New()

	//* HR-HELPERS ROUTES //

	HealthCheck(router)
	Seguridad(router)
	AccountRoutes(router)
	CandidatesRoutes(router)
	EmployeesRoutes(router)
	MercadoPago(router)
	return router
}
