package repository

import (
	"harvest/src/application/gateway"
)

type AuthImpl struct {
	Client gateway.Auth
}

func (a *AuthImpl) VerifyToken(token string) error {
	return a.Client.VerifyIDToken(token)
}
