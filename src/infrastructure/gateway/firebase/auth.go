package firebase

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

func InitializeAuth(app *firebase.App) *auth.Client {
	auth, err := app.Auth(context.Background())
	if err != nil {
		panic(err)
	}

	return auth
}

type AuthImpl struct {
	Client *auth.Client
}

func (a *AuthImpl) VerifyIDToken(token string) error {
	_, err := a.Client.VerifyIDToken(context.Background(), token)

	return err
}
