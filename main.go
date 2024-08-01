package main

import (
	"github.com/gin-gonic/gin"
	"github.com/josevitorrodriguess/e-learning/internal/config/db"
	"github.com/josevitorrodriguess/e-learning/internal/config/logger"
	"github.com/josevitorrodriguess/e-learning/internal/routes"
)

func main() {
	logger.Info("start program")

	r := gin.Default()

	conn, err := db.SetupDB()
	if err != nil {
		logger.Error("error to setup connection with database", err)
	}
	defer db.CloseConnection(conn)

	routes.InitRoutes(r, conn)
	r.Run()
}
