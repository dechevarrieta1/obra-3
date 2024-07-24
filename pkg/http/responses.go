package httputilsv1

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

type RestErrorResponse struct {
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func ResponseHandlers(ctx *fasthttp.RequestCtx, data interface{}, err interface{}, statusCode int, messages ...string) {

	ctx.SetContentType("application/json; charset=UTF-8")
	ctx.SetStatusCode(statusCode)

	message := ""
	for _, m := range messages {
		if message == "" {
			message = m
		} else {
			message = fmt.Sprintf("%s - %s", message, m)
		}
	}

	str := struct {
		Data    interface{} `json:"data"`
		Error   interface{} `json:"error"`
		Message interface{} `json:"message,omitempty"`
	}{
		Data:    data,
		Error:   err,
		Message: message,
	}

	serialized, err := json.Marshal(str)
	if err != nil {

		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBody([]byte("Serialization Error"))
		return
	}

	ctx.SetBody(serialized)
}
