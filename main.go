package main

import (
	"github.com/josevitorrodriguess/e-learning/internal/config/db"
	"github.com/josevitorrodriguess/e-learning/internal/config/logger"
)

func main() {
	logger.Info("START PROGRAM")

	conn, err := db.SetupDB()
	if err != nil {
		logger.Error("error to setup connection with database", err)
	}
	defer db.CloseConnection(conn)

}
