package userService

import (
	"errors"
	appError "eshop-products-ms/internal/apperror"
	"eshop-products-ms/internal/config"
	models "eshop-products-ms/internal/models/product"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log/slog"
	"time"
)

type UserStorage interface {
	SaveUser(authID uint) (uint, error)
	User(id uint) (models.User, error)
}

func (u User) DecodeJWT(JWT string) (models.User, error) {
	const op = "userService.User.DecodeJWT"

	log := u.log.With(
		slog.String("op", op),
	)

	log.Info("decoding JWT")

	conf := config.Get()
	secret := conf.JWT.Secret

	token, err := jwt.Parse(JWT, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected method")
		}
		return secret, nil
	})
	if err != nil {
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		authID := claims["sub"].(uint)
		exp := claims["exp"].(int64)
		if exp < time.Now().Unix() {
			return models.User{}, fmt.Errorf("%s: %w", op, appError.ExpiredToken)
		}
		user, err := u.userStorage.User(authID)
		if err != nil {
			if errors.Is(err, appError.UserNotFound) {
				newUser, err := u.userStorage.SaveUser(authID)
				if err != nil {
					return models.User{}, fmt.Errorf("%s: %w", op, err)
				}
				user, err = u.userStorage.User(newUser)
				if err != nil {
					return models.User{}, fmt.Errorf("%s: %w", op, err)
				}
				return user, nil
			}
			return models.User{}, fmt.Errorf("%s: %w", op, err)
		}
		return user, nil
	}
	return models.User{}, fmt.Errorf("%s: %w", op, appError.InvalidToken)
}
