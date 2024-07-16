package controller

import (
	"github.com/gin-gonic/gin"
	"go_api/helper"
	"go_api/model"
	"net/http"
)

// ErrorResponse represents the structure of an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// JWTResponse represents the structure of a JWT response
type JWTResponse struct {
	JWT string `json:"jwt"`
}

// @Summary     Register a new user
// @Description Register a new user with a username and password
// @Accept      json
// @Produce     json
// @Param       user body     model.AuthenticationInput true "User registration data"
// @Success     201  {object} model.User                "Created user"
// @Failure     400  {object} ErrorResponse             "Error message"
// @Router      /auth/register [post]

func Register(context *gin.Context) {
	var input model.AuthenticationInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		Username: input.Username,
		Password: input.Password,
	}

	savedUser, err := user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"user": savedUser})
}

// @Summary     Login a user
// @Description Login a user and return a JWT token
// @Accept      json
// @Produce     json
// @Param       user body     model.AuthenticationInput true "User login data"
// @Success     200  {object} JWTResponse               "JWT token"
// @Failure     400  {object} ErrorResponse             "Error message"
// @Router      /auth/login [post]
func Login(context *gin.Context) {
	var input model.AuthenticationInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := model.FindUserByUsername(input.Username)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.ValidatePassword(input.Password)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwt, err := helper.GenerateJWT(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"jwt": jwt})
}
