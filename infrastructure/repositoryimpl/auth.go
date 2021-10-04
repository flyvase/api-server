package repositoryimpl

import (
	"context"
	"harvest/core/apperror"

	"firebase.google.com/go/v4/auth"
)

type Auth struct {
	Client *auth.Client
}

func (a *Auth) VerifyToken(token string) error {
	_, err := a.Client.VerifyIDToken(context.Background(), token)
	if err != nil {
		if auth.IsUnknown(err) {
			return apperror.Unknown{Message: err.Error()}
		}
		return err
	}

	return nil
}
