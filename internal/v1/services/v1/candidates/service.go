package candidateservicesv1

import (
	"encoding/json"
	"log"

	candidateshelpersv1 "github.com/dechevarrieta1/hrhelpers/internal/v1/helpers/candidates"
	candidatesmodelsv1 "github.com/dechevarrieta1/hrhelpers/internal/v1/models/candidates"
	httputilsv1 "github.com/dechevarrieta1/hrhelpers/pkg/http"
	"github.com/valyala/fasthttp"
)

func GetCandidates(ctx *fasthttp.RequestCtx) {
	log.Println("[LOG][GetCandidates] initializing....")
	candidates, err := candidateshelpersv1.GetAllCandidatesByQuery(ctx.UserValue("company").(string))
	if err != nil {
		log.Println("[ERROR][GetCandidates] Error getting candidates: ", err)
		httputilsv1.ResponseHandlers(ctx, nil, err, fasthttp.StatusInternalServerError, "Error getting candidates")
	}
	log.Println("[LOG][GetCandidates] Candidates retrieved")
	httputilsv1.ResponseHandlers(ctx, candidates, nil, fasthttp.StatusOK, "Candidates retrieved")
}

func GetCandidatesFiltered(ctx *fasthttp.RequestCtx) {
	log.Println("[LOG][GetCandidatesFiltered] initializing....")
	candidates, err := candidateshelpersv1.GetCandidatesByFilter(ctx.Request.Body(), ctx.UserValue("company").(string))
	if err != nil {
		log.Println("[ERROR][GetCandidatesFiltered] Error getting candidates: ", err)
		httputilsv1.ResponseHandlers(ctx, nil, err, fasthttp.StatusInternalServerError, "Error getting candidates")
	}
	log.Println("[LOG][GetCandidatesFiltered] Candidates retrieved")
	httputilsv1.ResponseHandlers(ctx, candidates, nil, fasthttp.StatusOK, "Candidates retrieved")
}

func CreateCandidate(ctx *fasthttp.RequestCtx) {
	log.Println("[LOG][CreateCandidate] initializing....")

	candidate := candidatesmodelsv1.Candidate{}
	if err := json.Unmarshal(ctx.Request.Body(), &candidate); err != nil {
		log.Println("[ERROR][CreateCandidate] Error unmarshalling candidate data: ", err)
		httputilsv1.ResponseHandlers(ctx, nil, err, fasthttp.StatusBadRequest, "Invalid request body")
		return
	}

	if err := candidateshelpersv1.CreateCanidateByQuery(candidate, ctx.UserValue("company").(string)); err != nil {
		log.Println("[ERROR][CreateCandidate] Error saving candidate to mongo: ", err)
		httputilsv1.ResponseHandlers(ctx, nil, err, fasthttp.StatusInternalServerError, "Error saving candidate")
		return
	}

	httputilsv1.ResponseHandlers(ctx, nil, nil, fasthttp.StatusOK, "Candidate created")
}

func UpdateCandidate(ctx *fasthttp.RequestCtx) {
	log.Println("[LOG][UpdateCandidate] initializing....")
	candidateID := ctx.UserValue("id").(string)
	candidateData := candidatesmodelsv1.Candidate{}
	if err := json.Unmarshal(ctx.Request.Body(), &candidateData); err != nil {
		log.Println("[ERROR][UpdateCandidate] Error unmarshalling candidate data: ", err)
		httputilsv1.ResponseHandlers(ctx, nil, err, fasthttp.StatusBadRequest, "Invalid request body")
		return
	}

	candidateUpd, err := candidateshelpersv1.UpdateCandidateByQuery(candidateID, ctx.UserValue("company").(string), candidateData)
	if err != nil {
		log.Println("[ERROR][UpdateCandidate] Error updating candidate: ", err)
		httputilsv1.ResponseHandlers(ctx, nil, err, fasthttp.StatusInternalServerError, "Error updating candidate")
		return
	}

	log.Println("[LOG][UpdateCandidate] Candidate updated")
	httputilsv1.ResponseHandlers(ctx, candidateUpd, nil, fasthttp.StatusOK, "Candidate retrieved")
}

func DeleteCandidate(ctx *fasthttp.RequestCtx) {
	log.Println("[LOG][DeleteCandidate] initializing....")
	candidateID := ctx.UserValue("id").(string)
	candidateDel, err := candidateshelpersv1.DeleteCandidateByQuery(candidateID, ctx.UserValue("company").(string))
	if err != nil {
		log.Println("[ERROR][DeleteCandidate] Error deleting candidate: ", err)
		httputilsv1.ResponseHandlers(ctx, nil, err, fasthttp.StatusInternalServerError, "Error deleting candidate")
		return
	}
	log.Println("[LOG][DeleteCandidate] Candidate deleted")
	httputilsv1.ResponseHandlers(ctx, candidateDel, nil, fasthttp.StatusOK, "Candidate deleted")
}
