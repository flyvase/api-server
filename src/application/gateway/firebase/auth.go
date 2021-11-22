package firebase

type Auth interface {
	VerifyIDToken(string) error
}
