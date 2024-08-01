package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/josevitorrodriguess/e-learning/internal/controllers"
	"github.com/josevitorrodriguess/e-learning/internal/repository"
	"github.com/josevitorrodriguess/e-learning/internal/usecase"
)

func InitRoutes(server *gin.Engine, conn *sqlx.DB) {

	userRepo := repository.NewUserRepository(conn)
	useUC := usecase.NewUserUseCase(userRepo)
	controller := controllers.NewUserController(useUC)

	server.GET("/getUsers")
	server.GET("/getUser/{id}")
	server.POST("/createUser", controller.CreateUser)
	server.PUT("/updateUser")
	server.DELETE("/deleteUser/{id}")
}
