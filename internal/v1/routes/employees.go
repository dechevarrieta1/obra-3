package routesv1

import (
	"github.com/buaazp/fasthttprouter"
	employeesservicesv1 "github.com/dechevarrieta1/hrhelpers/internal/v1/services/v1/employees"
)

func EmployeesRoutes(router *fasthttprouter.Router) {
	//* employee encryption collection
	//todo check if this implementation is valid
	// router.POST("/employees/encrypt", (employeesservicesv1.EncryptEmployee))

	//* employee crud
	router.GET("/employees", (employeesservicesv1.GetEmployees))
	router.GET("/employees/query", (employeesservicesv1.GetEmployeesFiltered))
	router.POST("/employees", (employeesservicesv1.CreateEmployee))
	router.PUT("/employees/:id", (employeesservicesv1.UpdateEmployee))
	router.DELETE("/employees/:id", (employeesservicesv1.DeleteEmployee))
	//* employee metrics
	router.PUT("/employees/:id/metrics", (employeesservicesv1.UpdateEmployeeMetrics))
	// router.GET("/employees/:id/metrics", (employeesservicesv1.GetEmployeeMetrics))
}
