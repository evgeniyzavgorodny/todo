package errorHandler

type NotFoundError struct {
	message string
}

func (m *NotFoundError) Error() string {
	return m.message
}

func NewNotFoundError(err string) error {
	return &NotFoundError{
		message: err,
	}
}
