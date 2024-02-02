package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"log"
	"me/pickside/types"
	"me/pickside/util"
	"net/http"
)

type UserMetadata struct {
	types.AccountType `json:"accountType"`
}

type JWTClaims struct {
	jwt.Claims
	UserMetadata UserMetadata `json:"user_metadata"`
	ID           uuid.UUID    `json:"sub"`
	Email        string       `json:"email"`
}

func WithToken() gin.HandlerFunc {
	return func(g *gin.Context) {
		cookies := g.Request.Cookies()

		var refreshToken, accessToken *http.Cookie

		for _, cookie := range cookies {
			if cookie.Name == "refreshToken" {
				refreshToken = cookie
			}
			if cookie.Name == "accessToken" {
				accessToken = cookie
			}
		}

		if refreshToken == nil {
			g.JSON(http.StatusUnauthorized, "Token not found, please login")
		}
		if refreshToken != nil && accessToken == nil {
			token, err := util.VerifyToken(refreshToken.Value)
			if err != nil {
				log.Fatal("error")
			}
			fmt.Printf("token %v", token)
		}
		if refreshToken != nil && accessToken != nil {
			token, err := util.VerifyToken(accessToken.Value)
			if err != nil {
				log.Fatal("error")
			}
			fmt.Printf("token %v", token)
		}
	}
}
