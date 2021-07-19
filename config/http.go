package config

func AllowedOrigin() string {
	if Environment == "dev" {
		return "http://localhost:3000"
	} else if Environment == "stg" {
		return "https://retail-dot-flyvase-stg.an.r.appspot.com"
	} else {
		return "https://flyvase.net"
	}
}
