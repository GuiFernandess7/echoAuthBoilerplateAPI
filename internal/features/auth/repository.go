package auth

import (
	"database/sql"
	"errors"
)

type AuthRepositoryInterface interface {
	FindByEmail(email string) (*User, error)
	Create(user *User) error
}

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) FindByEmail(email string) (*User, error) {
	var user User
	err := r.db.QueryRow("SELECT id, email, password, name, created_at, updated_at FROM users WHERE email = $1", email).Scan(
		&user.ID, &user.Email, &user.Password, &user.Name, &user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *AuthRepository) Create(user *User) error {
	query := `
		INSERT INTO users (email, password, name)
		VALUES ($1, $2, $3)
		RETURNING id`

	err := r.db.QueryRow(query, user.Email, user.Password, user.Name).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}
