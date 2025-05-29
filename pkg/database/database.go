package database

import (
	"database/sql"
	"fmt"

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
		return nil, fmt.Errorf("erro ao abrir o banco de dados: %w", err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("erro ao conectar ao banco de dados: %w", err)
	}

	logger.Print("Conexão com o banco de dados estabelecida com sucesso!")
	return db, nil
}
