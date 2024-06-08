package exception

type BadRequestError struct {
	Message string
}

func (err *BadRequestError) Error() string {
	return err.Message
}

type ValidationError struct {
	Message string
}

func (err *ValidationError) Error() string {
	return err.Message
}

type UnauthorizedError struct {
	Message string
}

func (err *UnauthorizedError) Error() string {
	return err.Message
}

type NotFoundError struct {
	Message string
}

func (error *NotFoundError) Error() string {
	return error.Message
}
