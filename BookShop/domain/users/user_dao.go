package users

import (
	"context"
	"fmt"
	"github.com/Gharib110/BookShop/database/mysql"
	date_utils "github.com/Gharib110/BookShop/utils/date"
	"github.com/Gharib110/BookShop/utils/errors"
	"log"
	"net/http"
	"strings"
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

func (user *User) Update() *errors.RestErr {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err := mysql2.UserDB.ExecContext(ctx, UpdateUserQuery,
		user.FirstName, user.LastName, user.Email)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

func (user *User) Delete() *errors.RestErr {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err := mysql2.UserDB.ExecContext(ctx, DeleteUserQuery, user.ID)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := mysql2.UserDB.Prepare(queryFindByStatus)
	if err != nil {
		log.Println("error when trying to prepare find users by status statement", err)
		return nil, errors.NewInternalServerError("error when tying to get user")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		log.Println("error when trying to find users by status", err)
		return nil, errors.NewInternalServerError("error when tying to get user")
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.Status); err != nil {
			log.Println("error when scan user row into user struct", err)
			return nil, errors.NewInternalServerError("error when tying to get user")
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil
}

func (user *User) FindByEmailAndPassword() *errors.RestErr {
	stmt, err := mysql2.UserDB.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		log.Println("error when trying to prepare get user by email and password statement", err)
		return errors.NewInternalServerError("error when tying to find user")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email, user.Password, StatusActive)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.Status); getErr != nil {
		if strings.Contains(getErr.Error(), "no rows in result set") {
			return errors.NewNotFoundError("invalid user credentials")
		}
		log.Println("error when trying to get user by email and password", getErr)
		return errors.NewInternalServerError("error when tying to find user" + "database error")
	}
	return nil
}
