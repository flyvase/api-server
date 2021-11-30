package gateway

type Auth interface {
	VerifyIDToken(string) error
}
