package seguridadservicesv1

import (
	"encoding/json"
	"log"

	httputilsv1 "github.com/dechevarrieta1/hrhelpers/pkg/http"
	"github.com/valyala/fasthttp"
)

type Alumno struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
}

func GetAlumnos(ctx *fasthttp.RequestCtx) {
	log.Println("[LOG][GetAlumnos] initializing....")

	alumnos := []string{"Agustin Romero", "Mariano Becerra", "Lucas Domenica", "David Echevarrieta"}
	dataResponse, _ := json.Marshal(alumnos)

	httputilsv1.ResponseHandlers(ctx, dataResponse, nil, fasthttp.StatusOK, "Alumnos retrieved")
}

func CreateAlumno(ctx *fasthttp.RequestCtx) {
	log.Println("[LOG][CreateAlumno] initializing....")
	reqData := ctx.Request.Body()
	var alumno Alumno
	if err := json.Unmarshal(reqData, &alumno); err != nil {
		log.Println("[ERROR][CreateAlumno] Error unmarshalling alumno data: ", err)
		httputilsv1.ResponseHandlers(ctx, nil, err, fasthttp.StatusBadRequest, "Invalid request body")
		return
	}

	httputilsv1.ResponseHandlers(ctx, reqData, nil, fasthttp.StatusOK, "Alumno created")
}
