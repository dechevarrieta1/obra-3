package accountsmodelsv1

type AccountUserRequest struct {
	Name         string `json:"name" bson:"name"`
	LastName     string `json:"last_name" bson:"last_name"`
	AccountID    string `json:"account_id" bson:"account_id"`
	Company      string `json:"company" bson:"company"`
	Email        string `json:"email" bson:"email"`
	Password     string `json:"password" bson:"password"`
	Country      string `json:"country" bson:"country"`
	TimeCreation string `json:"time_creation" bson:"time_creation"`
}

type AccountUserLogin struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}
