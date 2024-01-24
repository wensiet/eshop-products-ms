package appError

import "errors"

var (
	InvalidTitle    = errors.New("empty file")
	InvalidPrice    = errors.New("invalid price")
	InvalidQuantity = errors.New("invalid quantity")
	ProductNotFound = errors.New("not found")
)

var (
	InvalidToken = errors.New("invalid token")
	UserNotFound = errors.New("user not found")
	ExpiredToken = errors.New("token expired")
)

var (
	InvalidMetadata = errors.New("invalid metadata")
)

var (
	Unauthorized = errors.New("unauthorized")
)
