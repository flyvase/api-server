package repository

import "harvest/domain/entity"

type Auth interface {
	VerifyToken(string) error
	SetCustomClaim(entity.User, string) error
}
