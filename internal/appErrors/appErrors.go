package appErrors

const (
	BadArg     = "0001"
	InternalDB = "0002"
	Security   = "0003"
)

type AppError struct {
	Type    string
	Message string
	Err     error
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}
