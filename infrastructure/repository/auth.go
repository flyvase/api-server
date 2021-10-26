package repository

import (
	"context"

	"firebase.google.com/go/v4/auth"

	"harvest/core/apperror"
)

type Auth struct {
	Client *auth.Client
}

func (ar *Auth) VerifyToken(token string) error {
	_, err := ar.Client.VerifyIDToken(context.Background(), token)
	if err != nil {
		if auth.IsUnknown(err) {
			return apperror.Unknown{Message: err.Error()}
		}

		return err
	}

	return nil
}
