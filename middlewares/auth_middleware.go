package middlewares

import (
	"github.com/gin-gonic/gin"
	"me/pickside/util"
	"net/http"
)

func WithToken() gin.HandlerFunc {
	return func(g *gin.Context) {
		refreshToken, err := g.Cookie("refreshToken")
		if err != nil {
			g.JSON(http.StatusNotFound, "not found please login")
			g.Abort()
			return
		}

		claims, err := util.GetTokenClaims(refreshToken)
		if err != nil {
			g.JSON(http.StatusForbidden, err)
			g.Abort()
			return
		}

		err = util.VerifyToken(refreshToken)
		if err != nil {
			g.JSON(http.StatusForbidden, err)
			g.Abort()
			return
		}

		accessToken, err := g.Cookie("accessToken")
		if err != nil {
			g.JSON(http.StatusForbidden, err)
			g.Abort()
			return
		}

		err = util.VerifyToken(accessToken)
		if err != nil {
			newAt, err := util.GenerateToken(
				claims.ID,
				claims.Username,
				claims.Email,
				claims.EmailVerified,
				"accessToken",
			)
			if err != nil {
				g.JSON(http.StatusInternalServerError, err)
				g.Abort()
				return
			}

			g.SetCookie(
				"accessToken",
				newAt,
				300000,
				"/api/v1",
				g.Request.Host,
				util.IsSecure(),
				true,
			)
		}
		return
	}
}
