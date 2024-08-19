package routesv1

import (
	"github.com/buaazp/fasthttprouter"
	middlewaresv1 "github.com/dechevarrieta1/hrhelpers/internal/v1/middlewares"
	candidateservicesv1 "github.com/dechevarrieta1/hrhelpers/internal/v1/services/v1/candidates"
)

func CandidatesRoutes(router *fasthttprouter.Router) {
	router.GET("/candidates", middlewaresv1.AuthMiddleware(candidateservicesv1.GetCandidates))
	router.GET("/candidates/query", middlewaresv1.AuthMiddleware(candidateservicesv1.GetCandidatesFiltered))
	router.POST("/candidates", middlewaresv1.AuthMiddleware(candidateservicesv1.CreateCandidate))
	router.PUT("/candidates/:id", middlewaresv1.AuthMiddleware(candidateservicesv1.UpdateCandidate))
	router.DELETE("/candidates/:id", middlewaresv1.AuthMiddleware(candidateservicesv1.DeleteCandidate))
}
