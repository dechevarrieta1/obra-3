package middlewaresv1

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/valyala/fasthttp"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

func AuthMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		tokenString := ctx.Request.Header.Peek("Authorization")
		if len(tokenString) == 0 {
			ctx.SetStatusCode(fasthttp.StatusUnauthorized)
			ctx.SetBodyString("Missing token")
			return
		}

		token, err := jwt.Parse(string(tokenString), func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected token method: %v", token.Header["alg"])
			}
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			ctx.SetStatusCode(fasthttp.StatusUnauthorized)
			ctx.SetBodyString("Invalid token")
			return
		}

		next(ctx)
	}
}

func GenerateJWT(accountID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"account_id": accountID,
		"exp":        time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
