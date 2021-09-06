package repository

type Auth interface {
	VerifyToken(string) error
	SetCustomClaim(string) error
}
