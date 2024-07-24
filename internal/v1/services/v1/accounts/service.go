package accountservicesv1

import (
	"log"

	"encoding/json"

	accountshelpersv1 "github.com/dechevarrieta1/obra-3/internal/v1/helpers/accounts"
	accountsmodelsv1 "github.com/dechevarrieta1/obra-3/internal/v1/models/accounts"
	httputilsv1 "github.com/dechevarrieta1/obra-3/pkg/http"
	"github.com/valyala/fasthttp"
)

func CreateAccount(ctx *fasthttp.RequestCtx) {
	log.Println("[LOG][CreateAccount] initializing....")

	account := accountsmodelsv1.AccountUserRequest{}
	dataToUnmarshal := ctx.Request.Body()
	if err := json.Unmarshal(dataToUnmarshal, &account); err != nil {
		httputilsv1.ResponseHandlers(ctx, nil, err, fasthttp.StatusBadRequest, "Invalid request body")
		return
	}

	accountCreated, err := accountshelpersv1.GenerateAccountWithJWT(account)
	if err != nil {
		httputilsv1.ResponseHandlers(ctx, nil, err, fasthttp.StatusInternalServerError, "Error generating JWT")
		return
	}

	httputilsv1.ResponseHandlers(ctx, accountCreated, nil, fasthttp.StatusOK, "Account created")
}

func LoginAccount(ctx *fasthttp.RequestCtx) {
	log.Println("[LOG][TestService] initializing....")
	httputilsv1.ResponseHandlers(ctx, nil, nil, fasthttp.StatusOK, "Login success")
}
