package repository

type Auth interface {
	VerifyToken(string) error
}
