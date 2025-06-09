package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func generateToken(req *LoginDTO, email string) (string, error) {
	claims := jwt.MapClaims{
		"sub": req.Email,
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", fmt.Errorf("failed to generate token")
	}
	return tokenString, nil
}
