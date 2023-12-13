package googleAuthIDTokenVerifier

import "errors"

var (
	ErrInvalidToken = errors.New("invalid token")

	ErrPublicKeyNotFound = errors.New("no public key found for given kid")

	ErrInvalidIssuer = errors.New("invalid issuer")

	ErrInvalidAudience = errors.New("invalid audience")
)
