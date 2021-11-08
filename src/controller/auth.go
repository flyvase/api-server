package controller

import "harvest/src/domain/repository"

func VerifyAuthToken(token string, authR repository.Auth) error {
	return authR.VerifyToken(token)
}
