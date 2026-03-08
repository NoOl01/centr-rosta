package errs

import "errors"

// 400
var (
	MissingBody   = New(BadRequest, errors.New("missing body"))
	MissingHeader = New(BadRequest, errors.New("missing header"))
	MissingQuery  = New(BadRequest, errors.New("missing query"))

	InvalidBody   = New(BadRequest, errors.New("invalid body"))
	InvalidHeader = New(BadRequest, errors.New("invalid header"))
	InvalidQuery  = New(BadRequest, errors.New("invalid query"))

	WrongTimeFormat = New(BadRequest, errors.New("wrong time format"))

	AlreadyExists = New(BadRequest, errors.New("record already exists"))
)

// 401
var (
	MissingAuthorizationToken = New(Unauthorized, errors.New("missing authorization token"))
	MissingSessionID          = New(Unauthorized, errors.New("missing redis id"))

	InvalidToken          = New(Unauthorized, errors.New("invalid token"))
	InvalidTokenClaims    = New(Unauthorized, errors.New("invalid token claims"))
	InvalidOrMissingClaim = New(Unauthorized, errors.New("invalid or missing claim"))
	InvalidSessionID      = New(Unauthorized, errors.New("invalid redis id"))
	SessionNotFound       = New(Unauthorized, errors.New("session not found"))

	UnexpectedSignMethod = New(Unauthorized, errors.New("unexpected sign method"))

	WrongPassword = New(Unauthorized, errors.New("wrong password"))
)

// 403
var (
	AccessDenied = New(Forbidden, errors.New("access denied"))
)

// 404
var (
	RecordNotFound = New(NotFound, errors.New("record not found"))
)

// 500
var (
	InternalError   = New(InternalServerError, errors.New("internal error"))
	DbInternalError = New(InternalServerError, errors.New("database internal error"))
)
