package googleAuthIDTokenVerifier

import (
	"errors"
	"time"
)

var (
	// ClockSkew - five minutes
	ClockSkew = time.Minute * 5

	// Issuers is the allowed oauth token issuers
	Issuers = []string{
		"accounts.google.com",
		"https://accounts.google.com",
	}
)

func NewVerifier() *Verifier {
	return &Verifier{
		ClockSkew: ClockSkew,
		Issuers:   Issuers,
	}
}

type Verifier struct {
	ClockSkew time.Duration
	Issuers   []string
}

func (v *Verifier) VerifyIDToken(idToken string, audience []string) error {
	_, err := v.ClaimIDToken(idToken, audience)
	return err
}

func (v *Verifier) ValidateIDTokenEmail(idToken string, audience []string, email string) error {
	clms, err := v.ClaimIDToken(idToken, audience)
	if err == nil {
		if clms.Email == email {
			return nil
		} else {
			return errors.New("email not matching")
		}
	}
	return err
}

func (v *Verifier) ValidateIDTokenSubject(idToken string, audience []string, subject string) error {
	clms, err := v.ClaimIDToken(idToken, audience)
	if err == nil {
		if clms.Subject == subject {
			return nil
		} else {
			return errors.New("subject not matching")
		}
	}
	return err
}

func (v *Verifier) ClaimIDToken(idToken string, audience []string) (*ClaimSet, error) {
	certs, err := getFederatedSignonCerts()
	if err != nil {
		return nil, err
	}
	return VerifySignedJWTWithCerts(idToken, certs, audience, Issuers)
}
