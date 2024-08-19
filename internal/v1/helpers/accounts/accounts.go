package accountshelpersv1

import (
	"context"
	"log"
	"os"

	uuidhelpersv1 "github.com/dechevarrieta1/hrhelpers/internal/v1/helpers/uuid"
	middlewaresv1 "github.com/dechevarrieta1/hrhelpers/internal/v1/middlewares"
	accountsmodelsv1 "github.com/dechevarrieta1/hrhelpers/internal/v1/models/accounts"
	errorsutilsv1 "github.com/dechevarrieta1/hrhelpers/pkg/errors"
	mongoutilsv1 "github.com/dechevarrieta1/hrhelpers/pkg/mongo"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

var (
	mongoConn = mongoutilsv1.MongoConnection{
		Url:      os.Getenv("HRHELPERS_MONGO_URL"),
		Database: "hrhelpers",
	}
)

func GenerateJWT(acc accountsmodelsv1.AccountUserRequest) (string, error) {
	accID := uuidhelpersv1.GenerateUUID()
	jwt, err := middlewaresv1.GenerateJWT(accID)
	if err != nil {
		return "", errorsutilsv1.HandleError("[Accountshelpersv1][GenerateAccountWithJWT]", "generating jwt", err)
	}

	return jwt, nil
}

func SaveAccountToMongo(acc accountsmodelsv1.AccountUserRequest) error {
	client, err := mongoConn.MakeBasicConnection()
	if err != nil {
		return errorsutilsv1.HandleError("[Accountshelpersv1][SaveAccountToMongo]", "connecting to mongo", err)
	}
	db := client.Database(mongoConn.Database)
	_, err = db.Collection("accounts").InsertOne(context.Background(), acc)
	if err != nil {
		return errorsutilsv1.HandleError("[Accountshelpersv1][SaveAccountToMongo]", "inserting one document", err)
	}
	return nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errorsutilsv1.HandleError("[Accountshelpersv1][HashPassword]", "hashing password", err)
	}

	return string(hashedPassword), nil
}

func LoginAccount(account accountsmodelsv1.AccountUserLogin) (string, error) {
	accountLoggin := accountsmodelsv1.AccountUserRequest{}
	log.Println("[LOG][LoginAccount] initializing....")
	client, err := mongoConn.MakeBasicConnection()
	if err != nil {
		return "", errorsutilsv1.HandleError("[Accountshelpersv1][LoginAccount]", "connecting to mongo", err)
	}
	db := client.Database(mongoConn.Database)
	err = db.Collection("accounts").FindOne(context.Background(), bson.M{"email": account.Email}).Decode(&accountLoggin)
	if err != nil {
		return "", errorsutilsv1.HandleError("[Accountshelpersv1][LoginAccount]", "No account recovered", err)
	}
	log.Println("[LOG][LoginAccount] Account recovered: ", accountLoggin)
	if !CheckPasswordHash(account.Password, accountLoggin.Password) {
		return "", errorsutilsv1.HandleError("[Accountshelpersv1][LoginAccount]", "Password does not match", err)
	}
	//generate JWT
	jwt, err := GenerateJWT(accountLoggin)
	if err != nil {
		return "", errorsutilsv1.HandleError("[Accountshelpersv1][LoginAccount]", "generating jwt", err)
	}
	return jwt, nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Println("[ERROR][CheckPasswordHash] Error checking password hash: ", err)
		return false
	} else {
		log.Println("[LOG][CheckPasswordHash] Password hash checked")
		return true
	}
}
