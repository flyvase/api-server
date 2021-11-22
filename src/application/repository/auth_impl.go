package repository

import (
	"harvest/src/application/gateway/firebase"
)

type AuthImpl struct {
	Client firebase.Auth
}

func (a *AuthImpl) VerifyToken(token string) error {
	return a.Client.VerifyIDToken(token)
}
