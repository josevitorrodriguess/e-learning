package db

import (
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/josevitorrodriguess/e-learning/internal/config/logger"
)


func SetupDB() (*sqlx.DB, error) {
	dsn := "root:12345@tcp(localhost:3306)/elearning?charset=utf8&parseTime=True&loc=Local"
	conn, err := sqlx.Open("mysql", dsn)
	if err != nil {
		logger.Error("fail to connect database", err)
		return nil, err
	}

	err = realizeMigrations(conn)
	if err != nil {
		logger.Error("fail to run migrations", err)
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		logger.Error("fail to ping database", err)
		return nil, err
	}

	logger.Info("success to init connection")
	return conn, nil
}

func CloseConnection(conn *sqlx.DB) {
	err := conn.Close()
	if err != nil {
		logger.Error("fail to close database connection", err)
	}
}


func realizeMigrations(conn *sqlx.DB) error {
	logger.Info("running migrations...")

	schema, err := os.ReadFile("internal/config/db/sql/schema.sql")
	if err != nil {
		logger.Error("error to read schema from db", err)
		return err
	}

	// Splitting the schema file into individual SQL statements
	statements := strings.Split(string(schema), ";")
	for _, statement := range statements {
		statement = strings.TrimSpace(statement)
		if statement == "" {
			continue
		}

		_, err = conn.Exec(statement)
		if err != nil {
			logger.Error("error executing statement", err)
			return err
		}
	}

	logger.Info("success to run migrations")
	return nil
}
