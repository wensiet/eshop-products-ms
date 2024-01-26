package userService

import (
	"log/slog"
)

type User struct {
	log         *slog.Logger
	userStorage UserStorage
}

func New(log *slog.Logger, userStorage UserStorage) *User {
	return &User{log: log, userStorage: userStorage}
}
