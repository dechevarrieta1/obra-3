package candidateshelpersv1

import (
	"context"
	"encoding/json"
	"log"
	"os"

	candidatesmodelsv1 "github.com/dechevarrieta1/hrhelpers/internal/v1/models/candidates"
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

func GetAllCandidatesByQuery() ([]candidatesmodelsv1.Candidate, error) {
	//TODO get all the candidates from mongo
	candidates := []candidatesmodelsv1.Candidate{}
	// make the conn to mongo
	client, err := mongoConn.MakeBasicConnection()
	if err != nil {
		return candidates, errorsutilsv1.HandleError("[GetAllCandidatesByQuery]", "connecting to mongo", err)
	}
	db := client.Database(mongoConn.Database).Collection("candidates")

	cursor, err := db.Find(context.Background(), bson.M{})
	if err != nil {
		return candidates, errorsutilsv1.HandleError("[GetAllCandidatesByQuery]", "getting candidates", err)
	}

	if err = cursor.All(context.Background(), &candidates); err != nil {
		return candidates, errorsutilsv1.HandleError("[GetAllCandidatesByQuery]", "getting candidates", err)
	}

	return candidates, nil
}

func GetCandidatesByFilter(query []byte) ([]candidatesmodelsv1.Candidate, error) {
	candidates := []candidatesmodelsv1.Candidate{}
	// make the conn to mongo
	client, err := mongoConn.MakeBasicConnection()
	if err != nil {
		return candidates, errorsutilsv1.HandleError("[GetCandidatesByFilter]", "connecting to mongo", err)
	}
	db := client.Database(mongoConn.Database).Collection("candidates")
	filter := bson.M{}
	if err := json.Unmarshal(query, &filter); err != nil {
		return candidates, errorsutilsv1.HandleError("[GetCandidatesByFilter]", "unmarshalling query", err)
	}

	var bsonFilter bson.D
	for key, value := range filter {
		bsonFilter = append(bsonFilter, bson.E{Key: key, Value: value})
	}
	cursor, err := db.Find(context.Background(), bsonFilter)
	if err != nil {
		return candidates, errorsutilsv1.HandleError("[GetCandidatesByFilter]", "getting candidates", err)
	}

	if err = cursor.All(context.Background(), &candidates); err != nil {
		return candidates, errorsutilsv1.HandleError("[GetCandidatesByFilter]", "getting candidates", err)
	}

	return candidates, nil
}

func CreateCanidateByQuery(candidate candidatesmodelsv1.Candidate) error {
	log.Println("[LOG][SaveCandidateToMongo] initializing....")
	client, err := mongoConn.MakeBasicConnection()
	if err != nil {
		return errorsutilsv1.HandleError("[SaveCandidateToMongo]", "connecting to mongo", err)
	}
	db := client.Database(mongoConn.Database)
	candidateID, err := db.Collection("candidates").InsertOne(context.Background(), candidate)
	if err != nil {
		return errorsutilsv1.HandleError("[SaveCandidateToMongo]", "inserting one document", err)
	}
	log.Println("[LOG][SaveCandidateToMongo] Candidate saved ID: ", candidateID.InsertedID)
	return nil
}

func UpdateCandidateByQuery(candidateID string, candidate candidatesmodelsv1.Candidate) (candidatesmodelsv1.Candidate, error) {
	log.Println("[LOG][UpdateCandidateByQuery] initializing....")

	candidateRes := candidatesmodelsv1.Candidate{}
	client, err := mongoConn.MakeBasicConnection()
	if err != nil {
		return candidateRes, errorsutilsv1.HandleError("[UpdateCandidateByQuery]", "connecting to mongo", err)
	}
	db := client.Database(mongoConn.Database).Collection("candidates")

	filter := bson.M{"candidate_id": candidateID}
	update := bson.M{"$set": candidate}
	err = db.FindOneAndUpdate(context.Background(), filter, update).Decode(&candidateRes)
	if err != nil {
		return candidateRes, errorsutilsv1.HandleError("[UpdateCandidateByQuery]", "updating candidate", err)
	}

	return candidateRes, nil
}

func DeleteCandidateByQuery(id string) (candidatesmodelsv1.Candidate, error) {
	log.Println("[LOG][DeleteCandidateByQuery] initializing....")

	candidateRes := candidatesmodelsv1.Candidate{}
	client, err := mongoConn.MakeBasicConnection()
	if err != nil {
		return candidateRes, errorsutilsv1.HandleError("[DeleteCandidateByQuery]", "connecting to mongo", err)
	}
	db := client.Database(mongoConn.Database).Collection("candidates")

	filter := bson.M{"candidate_id": id}
	err = db.FindOneAndDelete(context.Background(), filter).Decode(&candidateRes)
	if err != nil {
		return candidateRes, errorsutilsv1.HandleError("[DeleteCandidateByQuery]", "deleting candidate", err)
	}

	return candidateRes, nil
}
