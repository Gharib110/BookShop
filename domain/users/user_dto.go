package users

import (
	"github.com/Gharib110/bookstore_users_api/utils/errors"
	"strings"
)

type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func (user *User) ValidateUser() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	return nil
}

const (
	InsertQuery  = "INSERT INTO users(first_name, last_name, email, created_at) VALUES (?, ?, ?, ?)"
	GetUserQuery = "SELECT id, first_name, last_name, email, created_at FROM users WHERE id=?"
)
