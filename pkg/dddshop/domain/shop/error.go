package shop

type ItemIsAlreadySoldError struct {
	err error
}

func (e *ItemIsAlreadySoldError) Error() string {
	return e.err.Error()
}
