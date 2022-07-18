package users

import (
	"encoding/json"
	"github.com/Gharib110/bookstore_users_api/domain/users"
	"github.com/Gharib110/bookstore_users_api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

func FindUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
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
