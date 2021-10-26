package controller

import "harvest/domain/repository"

func VerifyAuthToken(token string, authR repository.Auth) error {
	return authR.VerifyToken(token)
}
