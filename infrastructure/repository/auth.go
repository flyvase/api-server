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
