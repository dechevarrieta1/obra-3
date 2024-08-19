package middlewaresv1

import (
	"log"
	"time"

	"github.com/valyala/fasthttp"
)

var (
	corsAllowHeaders     = "authorization,Content-Type,x-api-key,x-account-id, x-agent-migration, x-scope, x-section"
	corsAllowMethods     = "HEAD,GET,POST,PUT,DELETE,OPTIONS"
	corsAllowOrigin      = "*"
	corsAllowCredentials = "true"
)

func LoggerMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		start := time.Now()
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", corsAllowCredentials)
		ctx.Response.Header.Set("Access-Control-Allow-Headers", corsAllowHeaders)
		ctx.Response.Header.Set("Access-Control-Allow-Methods", corsAllowMethods)
		ctx.Response.Header.Set("Access-Control-Allow-Origin", corsAllowOrigin)
		next(ctx)

		duration := time.Since(start)
		logRequest(ctx, duration)
	}
}

func logRequest(ctx *fasthttp.RequestCtx, duration time.Duration) {
	method := string(ctx.Method())
	path := string(ctx.Path())
	statusCode := ctx.Response.StatusCode()
	clientIP := ctx.RemoteIP()
	userAgent := string(ctx.UserAgent())
	contentLength := len(ctx.Request.Body())

	log.Printf("[%s] | %s |  %s | %d | %s | %d | %s | %dms",
		clientIP,
		method,
		path,
		statusCode,
		userAgent,
		contentLength,
		time.Now().Format(time.RFC3339),
		duration.Milliseconds(),
	)
}
