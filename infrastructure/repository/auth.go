package repository

import (
	"context"
	"harvest/core/exception"
	"harvest/domain/entity"

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

func (a *AuthImpl) SetCustomClaim(u entity.User, id string) error {
	claims := map[string]interface{}{"id": id}
	err := a.Client.SetCustomUserClaims(context.Background(), u.Uid, claims)
	if err != nil {
		return err
	}

	return nil
}
