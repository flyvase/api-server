package config

func AllowedOrigin() []string {
	if Environment == "stg" {
		return []string{"https://retail-dot-flyvase-stg.an.r.appspot.com"}
	} else if Environment == "prod" {
		return []string{"https://retail.flyvase.net"}
	} else {
		return []string{"http://localhost:3000"}
	}
}
