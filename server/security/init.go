package security

import "os"

var secret = os.Getenv("JWT_SECRET")

func init() {
	if secret == "" {
		secret = "jwt-secret"
	}
}
