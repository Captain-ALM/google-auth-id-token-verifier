package googleAuthIDTokenVerifier

import (
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/exp/slices"
	"time"
)

func VerifySignedJWTWithCerts(token string, certs *Certs, allowedAuds []string, issuers []string) (*ClaimSet, error) {
	claimSet, err := jwt.ParseWithClaims(token, &ClaimSet{}, func(token *jwt.Token) (interface{}, error) {
		kid, vl := token.Header["kid"].(string)
		if !vl {
			return nil, ErrInvalidToken
		}
		key, found := certs.Keys[kid]
		if found {
			return key, nil
		}
		return nil, ErrPublicKeyNotFound
	}, jwt.WithLeeway(time.Duration(ClockSkew.Seconds())), jwt.WithIssuedAt(), jwt.WithExpirationRequired())
	if err != nil {
		return nil, err
	}

	Iss, err := claimSet.Claims.GetIssuer()
	if err != nil {
		return nil, err
	}
	Aud, err := claimSet.Claims.GetAudience()
	if err != nil {
		return nil, err
	}

	found := false
	for _, issuer := range issuers {
		if issuer == Iss {
			found = true
			break
		}
	}
	if !found {
		return nil, ErrInvalidIssuer
	}

	audFound := false
	for _, aud := range allowedAuds {
		if slices.Contains(Aud, aud) {
			audFound = true
			break
		}
	}
	if !audFound {
		return nil, ErrInvalidAudience
	}

	claims, vl := claimSet.Claims.(*ClaimSet)
	if !vl {
		return nil, ErrInvalidToken
	}

	return claims, nil
}
