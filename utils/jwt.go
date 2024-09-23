package utils

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func GenerateUserToken(userId uint, email string) (string, error) {
	if userId == 0 {
		return "", fmt.Errorf("jwt error: user id is empty")
	} else if email == "" {
		return "", fmt.Errorf("jwt error: email is empty")
	} else {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email":  email,
			"userId": userId,
			"exp":    time.Now().Add(time.Hour * 2).Unix(),
		})
		return token.SignedString([]byte(privateKey))
	}
}

func ValidateUserToken(c *gin.Context) error {
	token, err := getToken(c)
	if err != nil {
		return err
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}
	return errors.New("invalid token provided")
}

func getToken(c *gin.Context) (*jwt.Token, error) {
	tokenString := getTokenFromRequest(c)
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return privateKey, nil
	})
	return token, err
}

func getTokenFromRequest(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, "")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}
