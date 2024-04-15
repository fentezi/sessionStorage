package controllers

import (
	"net/http"

	"github.com/fentezi/session-auth/internal/models"
	"github.com/fentezi/session-auth/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthorizedController struct {
	serv *service.Service
}

func NewAuthorizedController(serv *service.Service) *AuthorizedController {
	return &AuthorizedController{
		serv: serv,
	}
}

func (a *AuthorizedController) SignUp(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID, err := a.serv.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"userID": userID,
	})
}

func (a *AuthorizedController) Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the home page",
	})
}

func (a *AuthorizedController) SignIn(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	uuid, err := a.serv.SignIn(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.SetCookie(
		"session_id",
		uuid,
		600,
		"/",
		"localhost",
		false,
		true,
	)
	c.Status(http.StatusOK)
}
