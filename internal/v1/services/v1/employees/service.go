package employeesservicesv1

import (
	"encoding/json"
	"log"

	employeeshelpersv1 "github.com/dechevarrieta1/hrhelpers/internal/v1/helpers/employees"
	employeesmodelsv1 "github.com/dechevarrieta1/hrhelpers/internal/v1/models/employees"
	httputilsv1 "github.com/dechevarrieta1/hrhelpers/pkg/http"
	"github.com/valyala/fasthttp"
)

func GetEmployees(ctx *fasthttp.RequestCtx) {
	log.Println("[LOG][GetEmployees] initializing....")
	employees, err := employeeshelpersv1.GetAllEmployeesByQuery()
	if err != nil {
		log.Println("[ERROR][GetEmployees] Error getting employees: ", err)
		httputilsv1.ResponseHandlers(ctx, nil, err, fasthttp.StatusInternalServerError, "Error getting employees")
	}
	log.Println("[LOG][GetEmployees] Employees retrieved")
	httputilsv1.ResponseHandlers(ctx, employees, nil, fasthttp.StatusOK, "Employees retrieved")
}

func GetEmployeesFiltered(ctx *fasthttp.RequestCtx) {
	log.Println("[LOG][GetEmployeesFiltered] initializing....")
	employees, err := employeeshelpersv1.GetEmployeesByFilter(ctx.Request.Body())
	if err != nil {
		log.Println("[ERROR][GetEmployeesFiltered] Error getting employees: ", err)
		httputilsv1.ResponseHandlers(ctx, nil, err, fasthttp.StatusInternalServerError, "Error getting employees")
	}
	log.Println("[LOG][GetEmployeesFiltered] Employees retrieved")
	httputilsv1.ResponseHandlers(ctx, employees, nil, fasthttp.StatusOK, "Employees retrieved")
}

func CreateEmployee(ctx *fasthttp.RequestCtx) {
	log.Println("[LOG][CreateEmployee] initializing....")

	employee := employeesmodelsv1.Employee{}
	if err := json.Unmarshal(ctx.Request.Body(), &employee); err != nil {
		log.Println("[ERROR][CreateEmployee] Error unmarshalling employee data: ", err)
		httputilsv1.ResponseHandlers(ctx, nil, err, fasthttp.StatusBadRequest, "Invalid request body")
		return
	}

	if err := employeeshelpersv1.CreateEmployeeByQuery(employee); err != nil {
		log.Println("[ERROR][CreateEmployee] Error saving employee to mongo: ", err)
		httputilsv1.ResponseHandlers(ctx, nil, err, fasthttp.StatusInternalServerError, "Error saving employee")
		return
	}

	httputilsv1.ResponseHandlers(ctx, nil, nil, fasthttp.StatusOK, "Employee created")
}

func UpdateEmployee(ctx *fasthttp.RequestCtx) {
	log.Println("[LOG][UpdateEmployee] initializing....")
	employeeID := ctx.UserValue("id").(string)
	employeeData := employeesmodelsv1.Employee{}
	if err := json.Unmarshal(ctx.Request.Body(), &employeeData); err != nil {
		log.Println("[ERROR][UpdateEmployee] Error unmarshalling employee data: ", err)
		httputilsv1.ResponseHandlers(ctx, nil, err, fasthttp.StatusBadRequest, "Invalid request body")
		return
	}

	if err := employeeshelpersv1.UpdateEmployeeByQuery(employeeID, employeeData); err != nil {
		log.Println("[ERROR][UpdateEmployee] Error updating employee: ", err)
		httputilsv1.ResponseHandlers(ctx, nil, err, fasthttp.StatusInternalServerError, "Error updating employee")
		return
	}

	httputilsv1.ResponseHandlers(ctx, nil, nil, fasthttp.StatusOK, "Employee updated")
}

func DeleteEmployee(ctx *fasthttp.RequestCtx) {
	log.Println("[LOG][DeleteEmployee] initializing....")
	employeeID := ctx.UserValue("id").(string)
	if err := employeeshelpersv1.DeleteEmployeeByQuery(employeeID); err != nil {
		log.Println("[ERROR][DeleteEmployee] Error deleting employee: ", err)
		httputilsv1.ResponseHandlers(ctx, nil, err, fasthttp.StatusInternalServerError, "Error deleting employee")
		return
	}

	httputilsv1.ResponseHandlers(ctx, nil, nil, fasthttp.StatusOK, "Employee deleted")
}

func UpdateEmployeeMetrics(ctx *fasthttp.RequestCtx) {
	log.Println("[LOG][UpdateEmployeeMetrics] initializing....")

	employeeID := ctx.UserValue("id").(string)
	employeeMetrics := employeesmodelsv1.PerformanceMetrics{}
	if err := json.Unmarshal(ctx.Request.Body(), &employeeMetrics); err != nil {
		log.Println("[ERROR][UpdateEmployeeMetrics] Error unmarshalling employee metrics data: ", err)
		httputilsv1.ResponseHandlers(ctx, nil, err, fasthttp.StatusBadRequest, "Invalid request body")
		return
	}

	employeeRes, err := employeeshelpersv1.UpdateEmployeeMetrics(employeeID, employeeMetrics)
	if err != nil {
		log.Println("[ERROR][UpdateEmployeeMetrics] Error updating employee metrics: ", err)
		httputilsv1.ResponseHandlers(ctx, nil, err, fasthttp.StatusInternalServerError, "Error updating employee metrics")
		return
	}

	httputilsv1.ResponseHandlers(ctx, employeeRes, nil, fasthttp.StatusOK, "Employee metrics updated")
}

// func GetEmployeeMetrics(ctx *fasthttp.RequestCtx) {
// 	log.Println("[LOG][GetEmployeeMetrics] initializing....")

// 	employeeID := ctx.UserValue("id").(string)
// 	employeeMetrics, err := employeeshelpersv1.GetEmployeeMetricsByID(employeeID)
// 	if err != nil {
// 		log.Println("[ERROR][GetEmployeeMetrics] Error getting employee metrics: ", err)
// 		httputilsv1.ResponseHandlers(ctx, nil, err, fasthttp.StatusInternalServerError, "Error getting employee metrics")
// 		return
// 	}

// 	httputilsv1.ResponseHandlers(ctx, employeeMetrics, nil, fasthttp.StatusOK, "Employee metrics retrieved")
// }
