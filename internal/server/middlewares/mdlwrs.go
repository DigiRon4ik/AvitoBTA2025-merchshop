package middlewares

import "github.com/golang-jwt/jwt/v5"

const authHeader = "Authorization"

type tokenManager interface {
	ParseClaims(string) (*jwt.MapClaims, error)
}

type Middlewares struct {
	tknMng tokenManager
}

func NewMiddlewares(tokenManager tokenManager) *Middlewares {
	return &Middlewares{tknMng: tokenManager}
}
