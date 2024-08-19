package employeesmodelsv1

import (
	"time"

	generalsmodelsv1 "github.com/dechevarrieta1/hrhelpers/internal/v1/models/generals"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Employee struct {
	ID             primitive.ObjectID                `json:"id" bson:"_id"`
	EmployeeID     string                            `json:"employee_id" bson:"employee_id"`
	Name           string                            `json:"name" bson:"name"`
	DateOfBirth    time.Time                         `json:"date_of_birth" bson:"date_of_birth"`
	Address        string                            `json:"address" bson:"address"`
	Phone          string                            `json:"phone" bson:"phone"`
	Email          string                            `json:"email" bson:"email"`
	AcademicTitle  string                            `json:"academic_title" bson:"academic_title"`
	DateOfEntry    time.Time                         `json:"date_of_entry" bson:"date_of_entry"`
	Position       string                            `json:"position" bson:"position"`
	Department     string                            `json:"department" bson:"department"`
	WorkExperience []generalsmodelsv1.WorkExperience `json:"work_experience" bson:"work_experience"`
	Skills         []string                          `json:"skills" bson:"skills"`
	Languages      []string                          `json:"languages" bson:"languages"`
	EmployeeStatus string                            `json:"employee_status" bson:"employee_status"`
	Notes          string                            `json:"notes" bson:"notes"`
	LastEvaluation time.Time                         `json:"last_evaluation" bson:"last_evaluation"`
	Performance    PerformanceMetrics                `json:"performance,omitempty" bson:"performance,omitempty"`
}

type PerformanceMetrics struct {
	ProductivityScore  float64   `json:"productivity_score,omitempty" bson:"productivity_score,omitempty"`
	QualityScore       float64   `json:"quality_score,omitempty" bson:"quality_score,omitempty"`
	InitiativeScore    float64   `json:"initiative_score,omitempty" bson:"initiative_score,omitempty"`
	TeamworkScore      float64   `json:"teamwork_score,omitempty" bson:"teamwork_score,omitempty"`
	AdaptabilityScore  float64   `json:"adaptability_score,omitempty" bson:"adaptability_score,omitempty"`
	CommitmentScore    float64   `json:"commitment_score,omitempty" bson:"commitment_score,omitempty"`
	CustomerFeedback   float64   `json:"customer_feedback,omitempty" bson:"customer_feedback,omitempty"`
	GoalsAchievement   float64   `json:"goals_achievement,omitempty" bson:"goals_achievement,omitempty"`
	LastReviewDate     time.Time `json:"last_review_date,omitempty" bson:"last_review_date,omitempty"`
	OverallPerformance float64   `json:"overall_performance,omitempty" bson:"overall_performance,omitempty"`
}
