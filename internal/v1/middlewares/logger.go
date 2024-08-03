package middlewaresv1

import (
	"log"
	"time"

	"github.com/valyala/fasthttp"
)

func LoggerMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		start := time.Now()

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
