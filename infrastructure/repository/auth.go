package repository

import (
	"context"
	"harvest/core/exception"

	"firebase.google.com/go/v4/auth"
)

type AuthImpl struct {
	Client *auth.Client
}

func (a *AuthImpl) VerifyToken(token string) error {
	_, err := a.Client.VerifyIDToken(context.Background(), token)
	if err != nil {
		if auth.IsUnknown(err) {
			return exception.UnknownError{Message: err.Error()}
		}
		return err
	}

	return nil
}

func (a *AuthImpl) SetCustomClaim(id string) error {
	claims := map[string]interface{}{"admin": true}
	err := a.Client.SetCustomUserClaims(context.Background(), id, claims)
	if err != nil {
		return err
	}

	return nil
}
