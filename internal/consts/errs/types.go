package errs

type Type int

const (
	BadRequest Type = iota
	Unauthorized
	Forbidden
	NotFound
	RequestTimeout
	InternalServerError
)
