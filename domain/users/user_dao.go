package users

import (
	"context"
	"github.com/Gharib110/bookstore_users_api/database/mysql"
	date_utils "github.com/Gharib110/bookstore_users_api/utils/date"
	"github.com/Gharib110/bookstore_users_api/utils/errors"
	"net/http"
	"time"
)

func (user *User) Get(userId int64) *errors.RestErr {
	if err := mysql2.UserDB.Ping(); err != nil {
		return &errors.RestErr{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
			Error:   "StatusInternalServerError",
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	err := mysql2.UserDB.QueryRowContext(ctx, GetUserQuery).Scan(&user.ID,
		&user.FirstName, &user.LastName,
		&user.Email, &user.CreatedAt)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	user.CreatedAt = date_utils.GetNowString()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	result, err := mysql2.UserDB.ExecContext(ctx, InsertQuery,
		user.FirstName, user.LastName, user.Email, user.CreatedAt)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	user.ID = userId
	return nil
}
