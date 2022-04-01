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

func CreateUser(c *gin.Context) {
	if c.Request.Header.Get("Content-Type") == "application/json" {
		err := c.Request.ParseForm()
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		userData := &users.User{
			ID:        0,
			FirstName: "",
			LastName:  "",
			Email:     "",
			CreatedAt: "",
		}
		decoder := json.NewDecoder(c.Request.Body)
		err = decoder.Decode(userData)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		_, restErr := services.CreateUser(userData)
		if restErr != nil {
			c.JSON(restErr.Code, restErr)
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "The User is Created!",
		})
		return
	} else {
		c.JSON(http.StatusUnsupportedMediaType, gin.H{
			"message": "Unsupported Media Type!",
		})
	}
}
