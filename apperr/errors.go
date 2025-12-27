package apperr

// MARK: - Invalid Argument

// InvalidArgumentError
/*
	This error is raised if one or multiple argument to a function has invalid or not-present mandatory parameters.
	Usually it means you have not properly called the function
	This is akin to an HTTP 400, GRPC InvalidArgument (3)
*/
type InvalidArgumentError struct{}

func (e InvalidArgumentError) Error() string {
	return string(e.Type())
}

func (e InvalidArgumentError) Type() ErrorType { return ErrorTypeInvalidArgument }

// MARK: - Unauthorized

// UnauthorizedError
/*
	This error is raised if you don't have the rights to access a ressource
	This is akin to an HTTP 403, GRPC PermissionDenied (7)
*/
type UnauthorizedError struct{}

func (e UnauthorizedError) Error() string {
	return string(e.Type())
}

func (e UnauthorizedError) Type() ErrorType { return ErrorTypeUnauthorized }

// MARK: - Resource Not Found

// ResourceNotFoundError
/*
	This error is raised if the resource you're trying to access doesn't exists (whether it is an url or a specific object)
	This is akin to an HTTP 404, GRPC NotFound (5)
*/
type ResourceNotFoundError struct{}

func (e ResourceNotFoundError) Error() string {
	return string(e.Type())
}

func (e ResourceNotFoundError) Type() ErrorType { return ErrorTypeResourceNotFound }

// MARK: - Conflict Error

// ResourceConflictError
/*
	This error is raised if the resource you're trying to access doesn't exists (whether it is an url or a specific object)
	This is akin to an HTTP 409, GRPC AlreadyExists (6)
*/
type ResourceConflictError struct{}

func (e ResourceConflictError) Error() string { return string(e.Type()) }

func (e ResourceConflictError) Type() ErrorType { return ErrorTypeConflict }

// MARK: - Internal Error

// InternalError
/*
	This error should be raised if an unrecoverable error has happened that cannot be covered by another error.
	Usually it means something unexpected broke
	This is akin to an HTTP 500, GRPC Internal (13)
*/
type InternalError struct{}

func (e InternalError) Error() string {
	return string(e.Type())
}

func (InternalError) Type() ErrorType { return ErrorTypeInternal }

// MARK: - Not Implemented

// NotImplementedError
/*
	This error is raised if the code hasn't yet been written for a specific function.
	It means the dev has to finish development
	This is akin to an HTTP 501, GRPC Internal (12)
*/
type NotImplementedError struct{}

func (e NotImplementedError) Error() string {
	return string(e.Type())
}

func (NotImplementedError) Type() ErrorType {
	return ErrorTypeNotImplemented
}

// MARK: - Unavailable Service

// UnavailableServiceError
/*
	This error should be raised if the service tried to call another service but it failed to respond
	It can be if the service wasn't reached or closed the connection unexpectedly
	This is akin to an HTTP 503, GRPC Unavailable (14)
*/
type UnavailableServiceError struct{}

func (e UnavailableServiceError) Error() string {
	return string(e.Type())
}

func (UnavailableServiceError) Type() ErrorType {
	return ErrorTypeUnavailable
}
