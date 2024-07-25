package db

import (
	"database/sql"
	"os"

	"github.com/josevitorrodriguess/e-learning/internal/config/logger"
	_ "github.com/lib/pq"
)

var (
	DB_URL = "DB_URL"
)

func InitConnection() {

	logger.Info("Trying to init  connection with database")

	connStr := os.Getenv(DB_URL)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Error("Error to connect to database", err)
	}

	logger.Info("Success to init connection")
	defer db.Close()
}
