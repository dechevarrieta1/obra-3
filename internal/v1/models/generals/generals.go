package generalsmodelsv1

import "time"

type WorkExperience struct {
	Position  string    `json:"position" bson:"position"`
	Company   string    `json:"company" bson:"company"`
	StartDate time.Time `json:"start_date" bson:"start_date"`
	EndDate   time.Time `json:"end_date" bson:"end_date"`
}
