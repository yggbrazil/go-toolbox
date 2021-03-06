package jwt

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/yggbrazil/go-toolbox/log"
)

var claims map[string]interface{}

// Middleware for Echo framework
func MiddlewareForHeader(secret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {

			tokenRequest := c.Request().Header.Get(HEADER)

			if secret == "" {
				return &echo.HTTPError{Code: 400, Message: "Chave não informada"}
			}

			token := strings.Replace(tokenRequest, "Bearer ", "", -1)

			if token == "" {
				return &echo.HTTPError{Code: 401, Message: "Header Authorization não informado corretamente"}
			}

			claims, err = VerifyTokenAndGetClaims(token, secret)

			if err != nil {
				l := log.New()
				l.Error(err)
				return &echo.HTTPError{Code: 401, Message: "Token inválido"}
			}

			c.Set(CLAIMS, claims)

			return next(c)
		}
	}
}
