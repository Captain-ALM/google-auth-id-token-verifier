package googleAuthIDTokenVerifier

import "errors"

var (
	ErrInvalidToken = errors.New("invalid token")

	ErrPublicKeyNotFound = errors.New("no public key found for given kid")

	ErrWrongSignature = errors.New("wrong token signature")

	ErrNoIssueTimeInToken = errors.New("no issue time in token")

	ErrNoExpirationTimeInToken = errors.New("no expiration time in token")

	ErrExpirationTimeTooFarInFuture = errors.New("expiration time too far in future")

	ErrTokenUsedTooEarly = errors.New("token used too early")

	ErrTokenUsedTooLate = errors.New("token used too late")
)
