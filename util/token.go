package util

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"me/pickside/types"
	"os"
	"time"
)

var secretKey = []byte(os.Getenv("TOKEN_SECRET"))

func GenerateToken(userId float64, username string, email string, emailVerified bool, tokenType string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	expTime := make(map[string]interface{})
	expTime["refreshToken"] = time.Now().AddDate(1, 0, 0).Unix()    //1 year
	expTime["accessToken"] = time.Now().Add(time.Minute * 5).Unix() //5m minutes

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userId
	claims["username"] = username
	claims["email"] = email
	claims["email_verified"] = emailVerified
	claims["exp"] = expTime[tokenType]

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func GetTokenClaims(tokenString string) (*types.JWTClaims, error) {
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

	convertedClaims := &types.JWTClaims{
		ID:            claims["user_id"].(float64),
		Username:      claims["username"].(string),
		Email:         claims["email"].(string),
		EmailVerified: claims["email_verified"].(bool),
	}

	return convertedClaims, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
