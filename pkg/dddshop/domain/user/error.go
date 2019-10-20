package user

type NotFoundUserError struct {
	err error
}

func (e *NotFoundUserError) Error() string {
	return e.err.Error()
}

type SameEmailUserAlreadyExistError struct {
	err error
}

func (e *SameEmailUserAlreadyExistError) Error() string {
	return e.err.Error()
}

type NotFoundUserRepositoryError struct {
	err error
}

func (e *NotFoundUserRepositoryError) Error() string {
	return e.err.Error()
}
