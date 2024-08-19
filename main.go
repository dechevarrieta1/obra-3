package main

import (
	"fmt"
	"log"
	"strings"

	middlewaresv1 "github.com/dechevarrieta1/hrhelpers/internal/v1/middlewares"
	routesv1 "github.com/dechevarrieta1/hrhelpers/internal/v1/routes"

	"github.com/valyala/fasthttp"
)

func main() {
	initFastHttp()
}

const defaultPort = "8080"

func initFastHttp() fasthttp.RequestHandler {
	var fasthttpHandler fasthttp.RequestHandler
	//? what port to use
	port := defaultPort
	if strings.TrimSpace(port) == "" {
		port = defaultPort
	}

	router := routesv1.InitRoutes()

	handlerwithLogging := middlewaresv1.LoggerMiddleware(router.Handler)

	log.Println("Starting http server on port: ", port)
	server := fasthttp.Server{
		Name:               "hrhelpers-controller",
		ReadBufferSize:     4096 * 3,
		Handler:            handlerwithLogging,
		MaxRequestBodySize: 100560416,
	}
	log.Println("Starting http server on port: ", port)
	if err := server.ListenAndServe(":" + defaultPort); err != nil {
		//Throw
		fmt.Errorf("Error starting http server: %v", err)
	}
	return fasthttpHandler
}
