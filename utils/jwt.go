package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func GenerateToken(userId uint, email string) (string, error) {
	if userId == 0 {
		return "", fmt.Errorf("jwt error: token id is empty")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(privateKey))
}

func VerifyToken(token string) (uint, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(privateKey), nil
	})

	if err != nil {
		log.Printf("token err: %v", err)
		return 0, errors.New("could not parse token")
	}

	if !parsedToken.Valid {
		log.Printf("token err: %v", err)
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	userIdFloat, ok := claims["userId"].(float64)
	if !ok {
		return 0, errors.New("userid claim is not a valid float64")
	}

	userId := uint(userIdFloat)

	return userId, nil
}
