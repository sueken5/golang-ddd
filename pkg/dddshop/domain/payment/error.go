package payment

type NotFoundAccountError struct {
	err error
}

func (e *NotFoundAccountError) Error() string {
	return e.err.Error()
}

type SameEmailAccountAlreadyExistError struct {
	err error
}

func (e *SameEmailAccountAlreadyExistError) Error() string {
	return e.err.Error()
}

type NotFoundAccountRepositoryError struct {
	err error
}

func (e *NotFoundAccountRepositoryError) Error() string {
	return e.err.Error()
}
