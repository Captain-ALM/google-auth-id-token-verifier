package googleAuthIDTokenVerifier

import "github.com/golang-jwt/jwt/v5"

type ClaimSet struct {
	jwt.RegisteredClaims
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Locale        string `json:"locale"`
	HostedDomain  string `json:"hd"`
}
