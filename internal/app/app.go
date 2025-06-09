// internal/app/app.go
package app

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"

	"github.com/GuiFernandess7/echoAuthBoilerplate/internal/config"
	"github.com/GuiFernandess7/echoAuthBoilerplate/internal/features/auth"
	"github.com/GuiFernandess7/echoAuthBoilerplate/internal/features/health"
	"github.com/GuiFernandess7/echoAuthBoilerplate/pkg/database"
	log "github.com/GuiFernandess7/echoAuthBoilerplate/pkg/logger"
)

type Application struct {
	Config *config.AppConfig
	Logger *log.Logger
	DB     *sql.DB
	Echo   *echo.Echo
}

func NewApplication(cfg *config.AppConfig) (*Application, error) {
	logLevelMap := map[string]log.LogLevel{
		"DEBUG": log.DEBUG,
		"INFO":  log.INFO,
		"WARN":  log.WARN,
		"ERROR": log.ERROR,
		"FATAL": log.FATAL,
	}
	minLevel, ok := logLevelMap[cfg.LogLevel]
	if !ok {
		minLevel = log.INFO
		fmt.Fprintf(os.Stderr, "Log level'%s' unknows. Using INFO.\n", cfg.LogLevel)
	}

	logDir := "logs"
	logPath := filepath.Join(logDir, "app.log")
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to created log folder '%s': %w", logDir, err)
	}

	appLogger, err := log.NewLogger(logPath, minLevel)
	if err != nil {
		return nil, fmt.Errorf("failed to configure logger: %w", err)
	}
	appLogger.Print("Logger initialized successfull.")

	db, err := database.InitDB(cfg.DB, appLogger)
	if err != nil {
		appLogger.Fatal("Failed to initialized database: %v", err)
		return nil, err
	}
	appLogger.Info("Database initialized successfully.")

	e := echo.New()

	app := &Application{
		Config: cfg,
		Logger: appLogger,
		DB:     db,
		Echo:   e,
	}
	app.setupRoutes()

	return app, nil
}

func (a *Application) setupRoutes() {
	health.RegisterRoutes(a.Echo)
	auth.RegisterRoutes(a.Echo, a.DB, a.Logger)

	a.Echo.GET("/", func(c echo.Context) error {
		a.Logger.Info("Request GET / received in the root.")
		return c.String(http.StatusOK, "Hello World from app.go!")
	})
}

func (a *Application) Run() {
	a.Logger.Info("Echo server initializing in port :%s", a.Config.ServerPort)
	a.Logger.Fatal("Server failed: %v", a.Echo.Start(fmt.Sprintf(":%s", a.Config.ServerPort)))
}

func (a *Application) Close() {
	if a.DB != nil {
		if err := a.DB.Close(); err != nil {
			a.Logger.Error("Error closing database: %v", err)
		} else {
			a.Logger.Info("Database connection closed.")
		}
	}
}
