package candidatesmodelsv1

import (
	"time"

	generalsmodelsv1 "github.com/dechevarrieta1/hrhelpers/internal/v1/models/generals"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Candidate struct {
	ID                primitive.ObjectID                `json:"id,omitempty" bson:"_id,omitempty"`
	CandidateID       string                            `json:"candidate_id,omitempty" bson:"candidate_id,omitempty"`
	Name              string                            `json:"name,omitempty" bson:"name,omitempty"`
	DateOfBirth       time.Time                         `json:"date_of_birth,omitempty" bson:"date_of_birth,omitempty"`
	Address           string                            `json:"address,omitempty" bson:"address,omitempty"`
	Phone             string                            `json:"phone,omitempty" bson:"phone,omitempty"`
	Email             string                            `json:"email,omitempty" bson:"email,omitempty"`
	AcademicTitle     string                            `json:"academic_title,omitempty" bson:"academic_title,omitempty"`
	WorkExperience    []generalsmodelsv1.WorkExperience `json:"work_experience,omitempty" bson:"work_experience,omitempty"`
	Skills            []string                          `json:"skills,omitempty" bson:"skills,omitempty"`
	Languages         []string                          `json:"languages,omitempty" bson:"languages,omitempty"`
	Resume            string                            `json:"resume,omitempty" bson:"resume,omitempty"`
	CoverLetter       string                            `json:"cover_letter,omitempty" bson:"cover_letter,omitempty"`
	Certifications    []string                          `json:"certifications,omitempty" bson:"certifications,omitempty"`
	ApplicationStatus string                            `json:"application_status,omitempty" bson:"application_status,omitempty"`
	Notes             string                            `json:"notes,omitempty" bson:"notes,omitempty"`
	LastUpdated       time.Time                         `json:"last_updated,omitempty" bson:"last_updated,omitempty"`
	Availability      string                            `json:"availability,omitempty" bson:"availability,omitempty"`
	SalaryExpectation float64                           `json:"salary_expectation,omitempty" bson:"salary_expectation,omitempty"`
	References        []Reference                       `json:"references,omitempty" bson:"references,omitempty"`
}

type Reference struct {
	Name    string `json:"name" bson:"name"`
	Contact string `json:"contact" bson:"contact"`
}
