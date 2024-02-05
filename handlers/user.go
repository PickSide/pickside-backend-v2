package handlers

import (
	"errors"
	"fmt"
	"me/pickside/data"
	"me/pickside/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Me struct {
	ID   uuid.UUID `json:"id"`
	User data.User `json:"user"`
}

func HandleMe(g *gin.Context) {
	refreshToken, err := g.Cookie("refreshToken")
	if err != nil {
		g.JSON(http.StatusUnauthorized, err)
	}

	parsedToken, err := util.ExtractClaims(refreshToken)
	if err != nil {
		g.JSON(http.StatusUnauthorized, err)
	}

	user, err := data.GetMe(uint64(parsedToken.ID))
	if err != nil {
		fmt.Printf("parsedToken.Id %v", parsedToken.ID)
		g.JSON(http.StatusInternalServerError, err)
	}

	g.JSON(http.StatusOK, gin.H{"result": user})
}

type AuthenticationRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func HandleLogin(g *gin.Context) {
	var authRequest AuthenticationRequest

	if err := g.ShouldBindJSON(&authRequest); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": errors.New("bad payload").Error()})
		return
	}

	user, err := data.UserMatch(authRequest.Username, authRequest.Password)
	if err != nil {
		g.JSON(http.StatusNotFound, err)
		return
	}

	refreshToken, err := util.GenerateRefresh(user.ID, user.Username, user.Email, user.EmailVerified)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
		return
	}

	err = data.InsertNewToken(refreshToken, user.ID)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
		return
	}

	accessToken, err := util.GenerateAccess(user.ID, user.Username, user.Email, user.EmailVerified)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
		return
	}

	err = data.InsertNewToken(accessToken, user.ID)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
		return
	}

	g.SetCookie(
		"refreshToken",
		refreshToken,
		3.154e10,
		"/api/v1",
		g.Request.Host,
		util.IsSecure(),
		true,
	)

	g.SetCookie(
		"accessToken",
		accessToken,
		300000,
		"/api/v1",
		g.Request.Host,
		util.IsSecure(),
		true,
	)

	g.JSON(http.StatusOK, gin.H{
		"result": user,
	})
}

func HandleLogout(g *gin.Context) {
	g.SetCookie(
		"accessToken",
		"",
		-1,
		"/api/v1",
		g.Request.Host,
		util.IsSecure(),
		true,
	)
	g.SetCookie(
		"refreshToken",
		"",
		-1,
		"/api/v1",
		g.Request.Host,
		util.IsSecure(),
		true,
	)
	g.JSON(http.StatusOK, gin.H{"message": "logged out successfully"})
}

func HandleCreateMe(c *gin.Context) {
	//	dbInstance := db.GetDB()
	//	var user_req UserRequest
	//
	//	if err := c.ShouldBindJSON(&user_req); err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//		return
	//	}
	//
	//	log.Println(user_req.Username)
	//	log.Println(user_req.Password)
	//
	//	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user_req.Password), 10)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	rows, err := dbInstance.Query(queries.InsertUser, "tonya", "tonyown11@gmail.com", user_req.Username, hashedPassword)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	defer rows.Close()
	//
	//	c.JSON(http.StatusCreated, user_req)
	//
}
