package auth

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/GuiFernandess7/echoAuthBoilerplate/pkg/logger"
	"github.com/labstack/echo/v4"
)

func HandleLogin(db *sql.DB, logger *logger.Logger) echo.HandlerFunc {
	repo := NewAuthRepository(db)

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
		user, err := repo.FindByEmail(loginRequest.Email)
		if err != nil {
			logger.Error("Error finding user: %v", err)
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Invalid credentials",
			})
		}

		// TODO: verify password

		logger.Printf("Login successful for user: %s", user.Email)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Login successful!",
			"user": map[string]interface{}{
				"id":    user.ID,
				"email": user.Email,
				"name":  user.Name,
			},
		})
	}
}

func HandleSignup(db *sql.DB, logger *logger.Logger) echo.HandlerFunc {
	repo := NewAuthRepository(db)

	return func(c echo.Context) error {
		var signupRequest RegisterDTO

		if err := c.Bind(&signupRequest); err != nil {
			logger.Error("Error binding signup request: %v", err)
			return c.JSON(echo.ErrBadRequest.Code, map[string]string{
				"message": "Request body invalid or malformed.",
				"error":   err.Error(),
			})
		}

		if err := ValidateEmail(signupRequest.Email); err != nil {
			logger.Error("Invalid email format: %s", signupRequest.Email)
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "Invalid email format",
				"error":   fmt.Sprintf("%v", err),
			})
		}

		if err := ValidatePassword(signupRequest.Password); err != nil {
			logger.Error("Invalid password: %s", signupRequest.Password)
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "Invalid password",
				"error":   fmt.Sprintf("%v", err),
			})
		}

		passwordHashed, err := hashPassword(signupRequest.Password)
		if err != nil {
			logger.Error("Error hashing password: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": "Error processing password",
				"error":   err.Error(),
			})
		}

		logger.Print("Signup request received: Email=%s", signupRequest.Email)
		user := &User{
			Email:    signupRequest.Email,
			Password: passwordHashed,
			Name:     signupRequest.Name,
		}

		if err := repo.Create(user); err != nil {
			logger.Error("Error creating user: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": "Error creating user",
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
