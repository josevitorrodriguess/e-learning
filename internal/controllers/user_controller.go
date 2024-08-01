package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/josevitorrodriguess/e-learning/internal/config/logger"
	"github.com/josevitorrodriguess/e-learning/internal/models"
	"github.com/josevitorrodriguess/e-learning/internal/usecase"
)

type userController struct {
	userUc usecase.UserUseCase
}

func NewUserController(userUc usecase.UserUseCase) userController {
	return userController{
		userUc: userUc,
	}
}

func (uc *userController) CreateUser(ctx *gin.Context) {
	var user models.User

	err := ctx.BindJSON(&user)
	if err != nil {
		logger.Error("error to binding user in JSON", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := uc.userUc.Create(user)
	if err != nil {
		logger.Error("fail to validate user informations", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdUser)
}
