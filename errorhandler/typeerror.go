package errorhandler

type NotFoundError struct {
	Message string
}

type ConflictError struct {
	Message string
}

type BadRequestError struct {
	Message string
}

type UnauthorizedError struct {
	Message string
}

type ForBiddenError struct {
	Message string
}

type InternalServerError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

func (e *BadRequestError) Error() string {
	return e.Message
}

func (e *ConflictError) Error() string {
	return e.Message
}

func (e *UnauthorizedError) Error() string {
	return e.Message
}

func (e *ForBiddenError) Error() string {
	return e.Message
}

func (e *InternalServerError) Error() string {
	return e.Message
}
