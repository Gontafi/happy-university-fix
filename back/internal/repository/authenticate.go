package repository

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

func (r *Repository) Authenticate(email, password string) bool {
	var storedPasswordHash string
	err := r.Db.QueryRow("SELECT password FROM users WHERE email = ?", email).Scan(&storedPasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		panic(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPasswordHash), []byte(password))
	if err != nil {
		return false
	}

	return true
}

func (r *Repository) CreateUser(name, email, role, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = r.Db.Query("INSERT INTO users (username, email, role, password) VALUES (?, ?, ?, ?)", name, email, role, string(hashedPassword))
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetRole(email string) (string, error) {
	var role string
	err := r.Db.QueryRow("SELECT role FROM users WHERE email = ?", email).Scan(&role)
	if err != nil {
		return "", err
	}

	return role, nil
}
