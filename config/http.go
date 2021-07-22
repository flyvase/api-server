package config

func AllowedOrigin() string {
	if Environment == "stg" {
		return "https://retail-dot-flyvase-stg.an.r.appspot.com"
	} else if Environment == "prod" {
		return "https://flyvase.net"
	} else {
		return "http://localhost:3000"
	}
}
