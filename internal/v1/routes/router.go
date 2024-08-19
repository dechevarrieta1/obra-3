package routesv1

import (
	"github.com/buaazp/fasthttprouter"
)

func InitRoutes() *fasthttprouter.Router {
	router := fasthttprouter.New()

	//* OBRA-3 ROUTES //

	HealthCheck(router)
	AccountRoutes(router)
	CandidatesRoutes(router)
	EmployeesRoutes(router)
	return router
}
