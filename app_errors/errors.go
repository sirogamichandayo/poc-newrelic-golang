package apperrors

type AppError struct {
	Err       error
	ErrStatus int
}

func (e *AppError) Error() string {
	return e.Err.Error()
}

func (e *AppError) Status() int {
	return e.ErrStatus
}
