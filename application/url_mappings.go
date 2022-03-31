package application

import (
	"github.com/Gharib110/bookstore_users_api/controllers/ping"
	"github.com/Gharib110/bookstore_users_api/controllers/users"
)

func mapURLS() {
	// Ping Controllers
	router.GET("ping-string", ping.PingString)
	router.GET("/ping-json", ping.PingJSON)

	// Users Controllers
	router.GET("/users/:user_id", users.GetUser)
	router.GET("/users/search", users.FindUser)
	router.POST("/users", users.CreateUser)
}
