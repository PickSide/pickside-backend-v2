package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"pickside/service/data"
	"pickside/service/util"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
		return
	}

	parsedToken, err := util.ExtractClaims(refreshToken)
	if err != nil {
		g.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	user, err := data.GetMe(parsedToken.ID)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	g.JSON(http.StatusOK, gin.H{"result": user})
	return
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

	user, err := data.MatchUser(authRequest.Username, authRequest.Password)
	if err != nil {
		g.JSON(http.StatusNotFound, err.Error())
		return
	}

	generateTokens(g, user)

	g.JSON(http.StatusOK, gin.H{
		"result": user,
	})
}

type LoginWithGoogleRequest struct {
	Email         string `json:"email" binding:"required"`
	Locale        string `json:"locale" binding:"required"`
	Name          string `json:"name" binding:"required"`
	Picture       string `json:"picture" binding:"omitempty"`
	VerifiedEmail bool   `json:"verifiedEmail" binding:"required"`
}

func HandleLoginWithGoogle(g *gin.Context) {
	var req LoginWithGoogleRequest

	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": errors.New("bad payload").Error()})
		return
	}

	user, err := data.MatchUserByEmail(req.Email, req.Locale, req.Name, req.Picture, req.VerifiedEmail)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	generateTokens(g, user)

	g.JSON(http.StatusOK, gin.H{
		"result": user,
	})
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
	g.JSON(http.StatusOK, gin.H{"message": "logged out successfully"})
}

type CreateUserRequest struct {
	FullName      string `binding:"required"`
	Username      string `binding:"required"`
	Email         string `binding:"required" validate:"required,email"`
	Password      string `binding:"required" validate:"required,min=8"`
	Phone         string `binding:"required"`
	Sexe          string `binding:"required" validate:"required,oneof=Male Female"`
	AgreedToTerms bool   `binding:"required"`
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

func HandleCreateMe(g *gin.Context) {
	var user_req CreateUserRequest

	if err := g.ShouldBindJSON(&user_req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := validateStruct(user_req)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data.InsertUser(
		user_req.FullName,
		user_req.Username,
		user_req.Email,
		user_req.Password,
		user_req.Phone,
		user_req.Sexe,
		user_req.AgreedToTerms,
		false,
	)

	g.JSON(http.StatusCreated, "user_req")
}

func HandleGetFavorites(g *gin.Context) {
	userIdString := g.Params.ByName("activityId")

	userId, err := strconv.ParseUint(userIdString, 10, 64)
	if err != nil {
		g.JSON(http.StatusNotFound, err.Error())
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
		"message": result,
	})
	return
}

func generateTokens(g *gin.Context, user *data.User) {
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
