package appError

import (
	"errors"
	"log/slog"
)

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

var errorsMap map[error]bool

func init() {
	errorsMap = map[error]bool{
		InvalidMetadata: true,
		Unauthorized:    true,
		UserNotFound:    true,
		ExpiredToken:    true,
		InvalidToken:    true,
	}
}

func LogIfNotApp(err error, logger *slog.Logger) {
	if !errorsMap[err] {
		logger.Error(err.Error())
	}
}
