package auth

import (
	"database/sql"

	"github.com/GuiFernandess7/echoAuthBoilerplate/pkg/logger"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, db *sql.DB, logger *logger.Logger) {
	auth := e.Group("/auth")
	{
		auth.POST("/login", HandleLogin(db, logger))
		auth.POST("/signup", HandleSignup(db, logger))
	}
}
