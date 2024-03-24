package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"pickside/service/data"
	"pickside/service/types"
	"pickside/service/util"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

func HandleGetAllUsers(g *gin.Context) {
	results, err := data.AllUsers()
	if err != nil {
		g.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"results": results,
	})
}

type CreateUserRequest struct {
	AgreedToTerms bool   `json:"agreedToTerms" binding:"required"`
	Email         string `json:"email" binding:"required" validate:"required,email"`
	FullName      string `json:"fullName" binding:"required"`
	Password      string `json:"password" binding:"required" validate:"required,min=8"`
	Phone         string `json:"phone" binding:"required"`
}

type Me struct {
	ID   uuid.UUID `json:"id"`
	User data.User `json:"user"`
}

func validateStruct(req CreateUserRequest) error {
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		var errorMessages []string
		for _, err := range err.(validator.ValidationErrors) {
			errorMessages = append(errorMessages, fmt.Sprintf("Field '%s' validation failed on tag '%s'", err.Field(), err.Tag()))
		}
		return fmt.Errorf("Validation errors:\n%s", errorMessages)
	}
	return nil
}

func HandleCreateUser(g *gin.Context) {
	var req CreateUserRequest

	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := validateStruct(req)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser := data.CreateUserStruct{
		AccountType:   types.DEFAULT,
		AgreedToTerms: true,
		Email:         req.Email,
		FullName:      req.FullName,
		Password:      []byte(req.Password),
		Phone:         req.Phone,
	}

	user, err := data.NewUser(newUser)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusCreated, gin.H{
		"result":      user,
		"redirectUri": "/",
	})
}

func HandleMe(g *gin.Context) {
	refreshToken, err := g.Cookie("refreshToken")
	if err != nil {
		g.JSON(http.StatusUnauthorized, err)
		return
	}

	parsedToken, err := util.ExtractClaims(refreshToken)
	if err != nil {
		g.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	user, err := data.MatchId(parsedToken.ID)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	g.JSON(http.StatusOK, gin.H{"result": user})
	return
}

type AuthenticationRequest struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RememberMe bool   `json:"rememberMe" binding:"omitempty"`
}

func HandleLogin(g *gin.Context) {
	var authRequest AuthenticationRequest

	var user *data.User
	var err error

	if err = g.ShouldBindJSON(&authRequest); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": errors.New("bad payload").Error()})
		return
	}

	isEmail := util.IsEmail(authRequest.Username)

	if isEmail {
		email := authRequest.Username
		user, err = data.MatchEmail(email, authRequest.Password)
	} else {
		user, err = data.MatchUsername(authRequest.Username, authRequest.Password)
	}
	if err != nil {
		g.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	generateTokens(g, user)

	g.JSON(http.StatusOK, gin.H{
		"result":      user,
		"redirectUri": "/",
	})
}

type LoginWithGoogleRequest struct {
	Email         string `json:"email" binding:"required"`
	GoogleID      string `json:"id" binding:"required"`
	Locale        string `json:"locale" binding:"required"`
	Name          string `json:"name" binding:"required"`
	Picture       string `json:"picture" binding:"omitempty"`
	VerifiedEmail bool   `json:"verified_email" binding:"required"`
}

func HandleLoginWithGoogle(g *gin.Context) {
	var req LoginWithGoogleRequest

	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": errors.New("bad payload").Error()})
		return
	}

	log.Println("GoogleID", req.GoogleID)

	user, err := data.MatchExternalId(req.GoogleID)
	if err != nil && err == sql.ErrNoRows {
		user, err = data.NewUser(data.CreateUserStruct{
			AccountType:   types.GOOGLE,
			Email:         req.Email,
			EmailVerified: req.VerifiedEmail,
			ExternalID:    req.GoogleID,
			FullName:      req.Name,
			Picture:       req.Picture,
		})
	}
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	generateTokens(g, user)

	g.JSON(http.StatusOK, gin.H{
		"result":      user,
		"redirectUri": "/",
	})
	return
}

func HandleLogout(g *gin.Context) {
	g.SetCookie(
		"accessToken",
		"",
		-1,
		"/api/"+os.Getenv("API_VERSION"),
		g.Request.Host,
		util.IsSecure(),
		true,
	)
	g.SetCookie(
		"refreshToken",
		"",
		-1,
		"/api/"+os.Getenv("API_VERSION"),
		g.Request.Host,
		util.IsSecure(),
		true,
	)
	g.JSON(http.StatusOK, gin.H{
		"message":     "logged out successfully",
		"redirectUri": "/login",
	})
	return
}

func HandleGetFavorites(g *gin.Context) {
	userId, err := strconv.ParseUint(g.Params.ByName("userId"), 10, 64)
	if err != nil {
		g.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	results, err := data.GetFavorites(userId)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	g.JSON(http.StatusOK, gin.H{"results": results})
	return
}

func HandleUpdateFavorites(g *gin.Context) {
	userIdString := g.Params.ByName("userId")

	activityIdString := g.Params.ByName("activityId")

	userId, err := strconv.ParseUint(userIdString, 10, 64)
	if err != nil {
		g.JSON(http.StatusNotFound, err.Error())
		return
	}

	activityId, err := strconv.ParseUint(activityIdString, 10, 64)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	result, err := data.UpdateFavorites(userId, activityId)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"result": result,
	})
	return
}

func HandleUpdateUser(g *gin.Context) {
	userIdString := g.Params.ByName("userId")

	userId, err := strconv.ParseUint(userIdString, 10, 64)
	if err != nil {
		g.JSON(http.StatusNotFound, err.Error())
		return
	}

	var settings map[string]interface{}

	if err := g.ShouldBindJSON(&settings); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = data.UpdateSettings(userId, settings)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Println("settings", settings)

	g.JSON(http.StatusOK, gin.H{
		"message": "succesfully updated",
		"result":  settings,
	})
	return
}

func generateTokens(g *gin.Context, user *data.User) {
	log.Println(user)
	refreshToken, err := util.GenerateRefresh(user.ID, user.Username, user.Email, user.EmailVerified)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
		return
	}

	// err = data.InsertNewToken(refreshToken, user.ID)
	// if err != nil {
	// 	g.JSON(http.StatusInternalServerError, err)
	// 	return
	// }

	accessToken, err := util.GenerateAccess(user.ID, user.Username, user.Email, user.EmailVerified)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
		return
	}

	// err = data.InsertNewToken(accessToken, user.ID)
	// if err != nil {
	// 	g.JSON(http.StatusInternalServerError, err)
	// 	return
	// }

	g.SetCookie(
		"refreshToken",
		refreshToken,
		3.154e10,
		"/api/"+os.Getenv("API_VERSION"),
		g.Request.Host,
		util.IsSecure(),
		true,
	)

	g.SetCookie(
		"accessToken",
		accessToken,
		300000,
		"/api/"+os.Getenv("API_VERSION"),
		g.Request.Host,
		util.IsSecure(),
		true,
	)

	return
}
