package exception

type NotMatchError struct {
	Error	string
}

func NewNotMatchError(error string) NotMatchError {
	return NotMatchError{Error: error}
}