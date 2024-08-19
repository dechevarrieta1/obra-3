package employeeshelpersv1

import (
	"context"
	"log"

	employeesmodelsv1 "github.com/dechevarrieta1/hrhelpers/internal/v1/models/employees"
	errorsutilsv1 "github.com/dechevarrieta1/hrhelpers/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

// todo metrics functions based in hours of works
func UpdateEmployeeMetrics(employeeID string, query employeesmodelsv1.PerformanceMetrics) (employeesmodelsv1.PerformanceMetrics, error) {
	log.Println("[LOG][UpdateEmployeeMetrics] initializing....")
	employeeMetrics := employeesmodelsv1.PerformanceMetrics{}
	client, err := mongoConn.MakeBasicConnection()
	if err != nil {
		return employeeMetrics, errorsutilsv1.HandleError("[UpdateEmployeeMetrics]", "connecting to mongo", err)
	}
	filter := bson.M{"_id": employeeID}
	update := bson.M{"$set": bson.M{"metrics": query}}

	db := client.Database(mongoConn.Database).Collection("employees")

	err = db.FindOneAndUpdate(context.Background(), filter, update).Decode(&employeeMetrics)
	if err != nil {
		return employeeMetrics, errorsutilsv1.HandleError("[UpdateEmployeeMetrics]", "updating employee metrics", err)
	}
	log.Println("[LOG][UpdateEmployeeMetrics] Employee metrics updated: ", employeeMetrics)

	return employeeMetrics, nil
}
