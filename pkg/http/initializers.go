package httputilsv1

import (
	"fmt"
	"log"
	"os"
	"strings"

	routesv1 "github.com/dechevarrieta1/obra-3/internal/v1/routes"
	"github.com/valyala/fasthttp"
)

const defaultPort = "8080"

var (
	corsAllowHeaders     = "authorization,Content-Type,x-api-key,x-account-id, x-agent-migration, x-scope, x-section, agent-id, channel-id"
	corsAllowMethods     = "HEAD,GET,POST,PUT,DELETE,OPTIONS"
	corsAllowOrigin      = "*"
	corsAllowCredentials = "true"
)

func InitFastHttp() fasthttp.RequestHandler {
	var fasthttpHandler fasthttp.RequestHandler
	log.Println("Starting up HTTP server on port:", defaultPort)

	port := os.Getenv("OBRA_3_PORT")
	if strings.TrimSpace(port) == "" {
		port = defaultPort
	}
	router := routesv1.InitRoutes() //TODO REFACTOR THIS
	fasthttpHandler = CORS(router.Handler)

	server := fasthttp.Server{
		Name:               "obra-3-controller",
		ReadBufferSize:     4096 * 3,
		Handler:            fasthttpHandler,
		MaxRequestBodySize: 100560416,
	}
	if err := server.ListenAndServe(":" + defaultPort); err != nil {
		//Throw
		fmt.Errorf("Error starting http server: %v", err)
	}
	return fasthttpHandler
}
