package repository

import (
	"context"

	"firebase.google.com/go/v4/auth"
)

type AuthImpl struct {
	Client *auth.Client
}

func (a *AuthImpl) VerifyToken(token string) error {
	_, err := a.Client.VerifyIDToken(context.Background(), token)
	return err
}
