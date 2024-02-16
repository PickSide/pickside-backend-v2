package util

import (
	"errors"
	"os"
	"pickside/service/types"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("TOKEN_SECRET"))

type TokenProps struct {
	UserID        uint64
	Username      string
	Email         string
	EmailVerified bool
}

func GenerateRefresh(userID uint64, username string, email string, emailVerified bool) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["username"] = username
	claims["email"] = email
	claims["email_verified"] = emailVerified
	claims["exp"] = time.Now().AddDate(1, 0, 0).Unix()

	tokenString, err := token.SignedString(secretKey)

	return tokenString, err
}

func GenerateAccess(userID uint64, username string, email string, emailVerified bool) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["username"] = username
	claims["email"] = email
	claims["email_verified"] = emailVerified
	claims["exp"] = time.Now().Add(time.Second * 5).Unix()

	tokenString, err := token.SignedString(secretKey)

	return tokenString, err
}

func ExtractClaims(tokenString string) (*types.JWTClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("error parsing claims")
	}

	userIDFloat64, ok := claims["user_id"].(float64)
	if !ok {
		return nil, errors.New("user_id not found or not a float64 in claims")
	}

	return &types.JWTClaims{
		ID:            uint64(userIDFloat64),
		Username:      claims["username"].(string),
		Email:         claims["email"].(string),
		EmailVerified: claims["email_verified"].(bool),
	}, nil
}

func IsTokenValid(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return false
	}

	return token.Valid
}
