package errors

type ErrorType int64

const (
	ErrNotFound ErrorType = iota
	ErrInvalidInput
	ErrUnauthorized
	ErrForbidden
	ErrInternal
	ErrConflict
	ErrNotImplemented
	ErrUnknown
)

type AppError interface {
	Error() string
	Code() string
}

type Error struct {
	error
	Type ErrorType
}

func (e *Error) Error() string {
	return e.error.Error()
}

func (e *Error) Code() string {
	return Code(e.Type)
}

func Code(typ ErrorType) string {
	switch typ {
	case ErrNotFound:
		return "not_found"
	case ErrInvalidInput:
		return "invalid_input"
	case ErrUnauthorized:
		return "unauthorized"
	case ErrForbidden:
		return "forbidden"
	case ErrInternal:
		return "internal"
	case ErrConflict:
		return "conflict"
	case ErrNotImplemented:
		return "not_implemented"
	case ErrUnknown:
		return "unknown"
	default:
		return "unknown"
	}
}

func Wrap(err error, typ ErrorType) *Error {
	return &Error{
		error: err,
		Type:  typ,
	}
}
