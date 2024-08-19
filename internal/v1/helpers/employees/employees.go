package employeeshelpersv1

import (
	"context"
	"encoding/json"
	"log"
	"os"

	employeesmodelsv1 "github.com/dechevarrieta1/hrhelpers/internal/v1/models/employees"
	errorsutilsv1 "github.com/dechevarrieta1/hrhelpers/pkg/errors"
	mongoutilsv1 "github.com/dechevarrieta1/hrhelpers/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	mongoConn = mongoutilsv1.MongoConnection{
		Url:      os.Getenv("HRHELPERS_MONGO_URL"),
		Database: "hrhelpers",
	}
)

func GetAllEmployeesByQuery() ([]employeesmodelsv1.Employee, error) {
	employees := []employeesmodelsv1.Employee{}
	// make the conn to mongo
	client, err := mongoConn.MakeBasicConnection()
	if err != nil {
		return employees, errorsutilsv1.HandleError("[GetAllEmployeesByQuery]", "connecting to mongo", err)
	}
	db := client.Database(mongoConn.Database).Collection("employees")
	cursor, err := db.Find(context.Background(), bson.M{})
	if err != nil {
		return employees, errorsutilsv1.HandleError("[GetAllEmployeesByQuery]", "getting employees", err)
	}
	if err = cursor.All(context.Background(), &employees); err != nil {
		return employees, errorsutilsv1.HandleError("[GetAllEmployeesByQuery]", "getting employees", err)
	}
	log.Println("[LOG][GetAllEmployeesByQuery] Employees retrieved: ", employees)
	return employees, nil
}

func GetEmployeesByFilter(query []byte) ([]employeesmodelsv1.Employee, error) {
	employees := []employeesmodelsv1.Employee{}
	client, err := mongoConn.MakeBasicConnection()
	if err != nil {
		return employees, errorsutilsv1.HandleError("[GetEmployeesByFilter]", "connecting to mongo", err)
	}
	db := client.Database(mongoConn.Database).Collection("employees")
	filter := bson.M{}
	if err := json.Unmarshal(query, &filter); err != nil {
		return employees, errorsutilsv1.HandleError("[GetCandidatesByFilter]", "unmarshalling query", err)
	}
	var bsonFilter bson.D
	for key, value := range filter {
		bsonFilter = append(bsonFilter, bson.E{Key: key, Value: value})
	}
	cursor, err := db.Find(context.Background(), bsonFilter)
	if err != nil {
		return employees, errorsutilsv1.HandleError("[GetEmployeesByFilter]", "getting employees", err)
	}
	if err = cursor.All(context.Background(), &employees); err != nil {
		return employees, errorsutilsv1.HandleError("[GetEmployeesByFilter]", "getting employees", err)
	}
	log.Println("[LOG][GetEmployeesByFilter] Employees retrieved: ", employees)
	return employees, nil
}

func CreateEmployeeByQuery(employee employeesmodelsv1.Employee) error {
	client, err := mongoConn.MakeBasicConnection()
	if err != nil {
		return errorsutilsv1.HandleError("[CreateEmployeeByQuery]", "connecting to mongo", err)
	}
	db := client.Database(mongoConn.Database).Collection("employees")
	_, err = db.InsertOne(context.Background(), employee)
	if err != nil {
		return errorsutilsv1.HandleError("[CreateEmployeeByQuery]", "creating employee", err)
	}
	log.Println("[LOG][CreateEmployeeByQuery] Employee created: ", employee)
	return nil
}

func UpdateEmployeeByQuery(employeeID string, employee employeesmodelsv1.Employee) error {
	client, err := mongoConn.MakeBasicConnection()
	if err != nil {
		return errorsutilsv1.HandleError("[UpdateEmployeeByQuery]", "connecting to mongo", err)
	}
	db := client.Database(mongoConn.Database).Collection("employees")
	_, err = db.UpdateOne(context.Background(), bson.M{"employee_id": employeeID}, bson.M{"$set": employee})
	if err != nil {
		return errorsutilsv1.HandleError("[UpdateEmployeeByQuery]", "updating employee", err)
	}
	log.Println("[LOG][UpdateEmployeeByQuery] Employee updated: ", employee)
	return nil
}

func DeleteEmployeeByQuery(employeeID string) error {
	client, err := mongoConn.MakeBasicConnection()
	if err != nil {
		return errorsutilsv1.HandleError("[DeleteEmployeeByQuery]", "connecting to mongo", err)
	}
	db := client.Database(mongoConn.Database).Collection("employees")
	_, err = db.DeleteOne(context.Background(), bson.M{"employee_id": employeeID})
	if err != nil {
		return errorsutilsv1.HandleError("[DeleteEmployeeByQuery]", "deleting employee", err)
	}
	log.Println("[LOG][DeleteEmployeeByQuery] Employee deleted: ", employeeID)
	return nil
}
