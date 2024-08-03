package accountshelpersv1

import (
	"context"
	"os"

	uuidhelpersv1 "github.com/dechevarrieta1/obra-3/internal/v1/helpers/uuid"
	middlewaresv1 "github.com/dechevarrieta1/obra-3/internal/v1/middlewares"
	accountsmodelsv1 "github.com/dechevarrieta1/obra-3/internal/v1/models/accounts"
	errorsutilsv1 "github.com/dechevarrieta1/obra-3/pkg/errors"
	mongoutilsv1 "github.com/dechevarrieta1/obra-3/pkg/mongo"
)

var (
	mongoConn = mongoutilsv1.MongoConnection{
		Url:      os.Getenv("OBRA_3_MONGO_URL"),
		Database: "obra-3",
	}
)

func GenerateAccountWithJWT(acc accountsmodelsv1.AccountUserRequest) (string, error) {
	accID := uuidhelpersv1.GenerateUUID()
	jwt, err := middlewaresv1.GenerateJWT(accID)
	if err != nil {
		return "", errorsutilsv1.HandleError("GenerateAccountWithJWT", "generating jwt", err)
	}

	return jwt, nil
}

func SaveAccountToMongo(acc accountsmodelsv1.AccountUserRequest) error {
	//todo save account to mongo
	client, err := mongoConn.MakeBasicConnection()
	if err != nil {
		return errorsutilsv1.HandleError("SaveAccountToMongo", "connecting to mongo", err)
	}
	db := client.Database(mongoConn.Database)
	_, err = db.Collection("accounts").InsertOne(context.Background(), acc)
	if err != nil {
		return errorsutilsv1.HandleError("SaveAccountToMongo", "inserting one document", err)
	}
	return nil
}
