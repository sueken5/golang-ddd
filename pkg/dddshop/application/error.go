package application

type ItemIsAlreadySoldError struct {
	err error
}

func (e *ItemIsAlreadySoldError) Error() string {
	return e.err.Error()
}

type UnkwonUserError struct {
	err error
}

func (e *UnkwonUserError) Error() string {
	return e.err.Error()
}

type NotPaymentAccountRegisteredError struct {
	err error
}

func (e *NotPaymentAccountRegisteredError) Error() string {
	return e.err.Error()
}

type UserAlreadyExistError struct {
	err error
}

func (e *UserAlreadyExistError) Error() string {
	return e.err.Error()
}

type PaymentAlreadySetupError struct {
	err error
}

func (e *PaymentAlreadySetupError) Error() string {
	return e.err.Error()
}

type UnexpectedError struct {
	err error
}

func (e *UnexpectedError) Error() string {
	return e.err.Error()
}
