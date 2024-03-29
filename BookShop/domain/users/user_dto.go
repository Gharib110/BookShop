package users

import (
	"github.com/Gharib110/BookShop/utils/errors"
	"strings"
)

const (
	StatusActive = "active"
)

type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	Status    string `json:"status"`
	Password  string `json:"password"`
}

type Users []User

func (user *User) ValidateUser() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	return nil
}

const (
	InsertQuery                 = "INSERT INTO users(first_name, last_name, email, created_at) VALUES (?, ?, ?, ?)"
	GetUserQuery                = "SELECT id, first_name, last_name, email, created_at FROM users WHERE id=?"
	UpdateUserQuery             = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?"
	DeleteUserQuery             = "DELETE FROM users WHERE id=?"
	queryFindByStatus           = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email=? AND password=? AND status=?"
)
