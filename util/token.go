package util

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"me/pickside/data"
	"os"
	"time"
)

func GenerateToken(user data.User, tokenType string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	expTime := make(map[string]interface{})
	expTime["refreshToken"] = time.Now().AddDate(1, 0, 0).Unix()
	expTime["accessToken"] = time.Now().Add(time.Minute * 5).Unix()

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = expTime[tokenType]

	tokenString, err := token.SignedString(os.Getenv("TOKEN_SECRET"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return os.Getenv("TOKEN_SECRET"), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	return token, nil
}
