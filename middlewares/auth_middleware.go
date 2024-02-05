package middlewares

import (
	"me/pickside/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func WithToken() gin.HandlerFunc {
	return func(g *gin.Context) {
		// Check refreshToken
		refreshToken, err := g.Cookie("refreshToken")
		if err != nil {
			g.JSON(http.StatusNotFound, "not found please login")
			g.Abort()
			return
		}

		// Check valid
		valid := util.IsTokenValid(refreshToken)
		if !valid {
			g.JSON(http.StatusForbidden, "token expired please login")
			g.Abort()
			return
		}

		userClaims, err := util.ExtractClaims(refreshToken)
		if err != nil {
			g.JSON(http.StatusInternalServerError, err.Error())
			g.Abort()
			return
		}

		accessToken, err := g.Cookie("accessToken")
		if err != nil {
			token, err := util.GenerateAccess(uint64(userClaims.ID), userClaims.Username, userClaims.Email, userClaims.EmailVerified)
			if err != nil {
				g.JSON(http.StatusInternalServerError, err.Error())
				g.Abort()
				return
			}

			g.SetCookie(
				"accessToken",
				token,
				int(time.Now().Add(time.Second*5).Unix()),
				"/api/v1",
				g.Request.Host,
				util.IsSecure(),
				true,
			)
		}

		if !util.IsTokenValid(accessToken) {
			token, err := util.GenerateAccess(uint64(userClaims.ID), userClaims.Username, userClaims.Email, userClaims.EmailVerified)
			if err != nil {
				g.JSON(http.StatusInternalServerError, err.Error())
				g.Abort()
				return
			}

			g.SetCookie(
				"accessToken",
				token,
				int(time.Now().Add(time.Second*5).Unix()),
				"/api/v1",
				g.Request.Host,
				util.IsSecure(),
				true,
			)
		}
		return
	}
}
