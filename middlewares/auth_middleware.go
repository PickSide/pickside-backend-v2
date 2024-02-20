package middlewares

import (
	"net/http"
	"os"
	"pickside/service/util"
	"time"

	"github.com/gin-gonic/gin"
)

func WithToken() gin.HandlerFunc {
	return func(g *gin.Context) {

		refreshToken, err := g.Cookie("refreshToken")
		if err != nil {
			g.JSON(http.StatusNotFound, "not found please login")
			g.Abort()
			return
		}

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
				"/api/"+os.Getenv("API_VERSION"),
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
				"/api/"+os.Getenv("API_VERSION"),
				g.Request.Host,
				util.IsSecure(),
				true,
			)
		}
		return
	}
}
