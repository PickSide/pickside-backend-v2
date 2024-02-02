package handlers

import (
	"errors"
	"me/pickside/data"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Me struct {
	ID   uuid.UUID `json:"id"`
	User data.User `json:"user"`
}

func HandleMe(g *gin.Context) {
	g.JSON(http.StatusOK, "OK")
}

type AuthenticationRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func HandleLogin(g *gin.Context) {
	var authRequest AuthenticationRequest

	if err := g.ShouldBindJSON(&authRequest); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": errors.New("bad payload").Error()})
		return
	}

	user, err, statusCode := data.Authenticate(authRequest.Username, authRequest.Password)
	if err != nil {
		g.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"result": user,
	})
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
