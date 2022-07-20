package users

import (
	"encoding/json"
	"github.com/Gharib110/bookstore_users_api/domain/users"
	"github.com/Gharib110/bookstore_users_api/services"
	"github.com/Gharib110/bookstore_users_api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("invalid user id")
		c.JSON(err.Code, err)
		return
	}
	result, saveErr := services.GetUser(userId)
	if saveErr != nil {
		c.JSON(saveErr.Code, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
	return
}

func UpdateUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("invalid user id")
		c.JSON(err.Code, err)
		return
	}

	var user users.User
	if err := c.BindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid JSON body")
		c.JSON(restErr.Code, restErr)
		return
	}

	user.ID = userId

	result, err := services.UpdateUser(&user)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, result)
	return
}

// CreateUser Use for create new user it is the Controller
func CreateUser(c *gin.Context) {
	if c.Request.Header.Get("Content-Type") == "application/json" {
		userData := &users.User{
			ID:        0,
			FirstName: "",
			LastName:  "",
			Email:     "",
			CreatedAt: "",
		}
		decoder := json.NewDecoder(c.Request.Body)
		err := decoder.Decode(userData)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		result, restErr := services.CreateUser(userData)
		if restErr != nil {
			c.JSON(restErr.Code, restErr)
			return
		}

		c.JSON(http.StatusCreated, result)
		return
	} else {
		c.JSON(http.StatusUnsupportedMediaType, gin.H{
			"message": "Unsupported Media Type!",
		})
	}
}

func DeleteUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("invalid user id")
		c.JSON(err.Code, err)
		return
	}

	if err := services.DeleteUser(userId); err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
	return
}
