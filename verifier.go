package googleAuthIDTokenVerifier

import "time"

var (
	// ClockSkew - five minutes
	ClockSkew = time.Minute * 5

	// Issuers is the allowed oauth token issuers
	Issuers = []string{
		"accounts.google.com",
		"https://accounts.google.com",
	}
)

type Verifier struct{}

func (v *Verifier) VerifyIDToken(idToken string, audience []string) error {
	_, err := v.ClaimIDToken(idToken, audience)
	return err
}

func (v *Verifier) ClaimIDToken(idToken string, audience []string) (*ClaimSet, error) {
	certs, err := getFederatedSignonCerts()
	if err != nil {
		return nil, err
	}
	return VerifySignedJWTWithCerts(idToken, certs, audience, Issuers)
}
