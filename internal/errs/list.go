package errs

import "errors"

var (
	ErrNotfound                    = errors.New("not found")
	ErrUserNotFound                = errors.New("user not found")
	ErrProductNotfound             = errors.New("product not found")
	ErrInvalidProductID            = errors.New("invalid product id")
	ErrInvalidRequestBody          = errors.New("invalid request body")
	ErrInvalidFieldValue           = errors.New("invalid field value")
	ErrInvalidProductName          = errors.New("invalid product name, min 4 symbols")
	ErrUsernameAlreadyExists       = errors.New("username already exists")
	ErrIncorrectUsernameOrPassword = errors.New("incorrect username or password")
	ErrInvalidToken                = errors.New("invalid token")
	ErrSomethingWentWrong          = errors.New("something went wrong")
	ErrItemNotAvailable   = errors.New("item is not available for selected dates")
	ErrInvalidStatus      = errors.New("invalid booking status")
	ErrInvalidDateRange   = errors.New("end date must be after start date")
	ErrPastDate          = errors.New("cannot book dates in the past")
)
