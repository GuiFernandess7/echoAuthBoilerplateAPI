package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	log "github.com/GuiFernandess7/echoAuthBoilerplate/pkg/logger"
)

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
	SSLMode  string
}

func InitDB(config DBConfig, logger *log.Logger) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
		config.SSLMode,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("%w", err)
	}

	logger.Print("Database connection established successfully!")
	return db, nil
}
