package config

func AllowedOrigin() []string {
	if Environment == "prod" {
		return []string{"https://retail.flyvase.net"}
	} else {
		return []string{"http://localhost:3000"}
	}
}
