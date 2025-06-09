package auth

import (
	"database/sql"
	"net/http"

	"github.com/GuiFernandess7/echoAuthBoilerplate/pkg/logger"
	"github.com/labstack/echo/v4"
)

func HandleLogin(db *sql.DB, logger *logger.Logger) echo.HandlerFunc {
	repo := NewAuthRepository(db)
	service := NewAuthService(repo, logger)

	return func(c echo.Context) error {
		var loginRequest LoginDTO

		if err := c.Bind(&loginRequest); err != nil {
			logger.Error("Error binding login request: %v", err)
			return c.JSON(echo.ErrBadRequest.Code, map[string]string{
				"message": "Request body invalid or malformed.",
				"error":   err.Error(),
			})
		}

		logger.Print("Login request received: Email=%s", loginRequest.Email)

		user, err := service.Login(loginRequest)
		if err != nil {
			logger.Error("Login error: %v", err)
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Login failed",
				"error":   err.Error(),
			})
		}

		token, err := generateToken(&loginRequest, user.Email)
		if err != nil {
			logger.Error("Error generating token: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": "Failed to generate token",
			})
		}

		logger.Printf("Login successful for user: %s", user.Email)
		return c.JSON(http.StatusOK, echo.Map{
			"access_token": token,
			"token_type":   "Bearer",
			"expires_in":   3600,
		})
	}
}

func HandleSignup(db *sql.DB, logger *logger.Logger) echo.HandlerFunc {
	repo := NewAuthRepository(db)
	service := NewAuthService(repo, logger)

	return func(c echo.Context) error {
		var signupRequest RegisterDTO

		if err := c.Bind(&signupRequest); err != nil {
			logger.Error("Error binding signup request: %v", err)
			return c.JSON(echo.ErrBadRequest.Code, map[string]string{
				"message": "Request body invalid or malformed.",
				"error":   err.Error(),
			})
		}

		logger.Print("Signup request received: Email=%s", signupRequest.Email)

		user, err := service.Signup(signupRequest)
		if err != nil {
			logger.Error("Signup error: %v", err)
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "Signup failed",
				"error":   err.Error(),
			})
		}

		logger.Printf("Signup successful for user: %s", user.Email)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Signup successful!",
			"user": map[string]interface{}{
				"id":    user.ID,
				"email": user.Email,
				"name":  user.Name,
			},
		})
	}
}
