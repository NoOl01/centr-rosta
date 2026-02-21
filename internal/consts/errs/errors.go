package errs

import "errors"

var (
	SessionNotFound        = errors.New("session not found")
	UnexpectedSignMethod   = errors.New("unexpected sign method")
	InvalidToken           = errors.New("invalid token")
	InvalidTokenClaimsType = errors.New("invalid token claims type")
	InvalidOrMissingClaim  = errors.New("invalid or missing claim")
	WrongPassword          = errors.New("wrong password")
	MissingQueryParameter  = errors.New("missing query parameter")
	MissingHeader          = errors.New("missing header")
)
