package helper

import (
	"go-todo-list/entity"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user entity.User) string {
	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(1 * time.Minute).Unix(),
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := t.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		panic(err)
	}

	return token
}
