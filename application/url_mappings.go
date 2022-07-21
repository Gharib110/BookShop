package application

import (
	"github.com/Gharib110/BookShop/controllers/ping"
	"github.com/Gharib110/BookShop/controllers/users"
)

func mapURLS() {
	// Ping Controllers
	router.GET("ping-string", ping.String)
	router.GET("/ping-json", ping.Json)

	// Users Controllers
	router.GET("/users/:user_id", users.GetUser)
	router.PATCH("/users/:user_id", users.UpdateUser)
	router.POST("/users", users.CreateUser)
	router.DELETE("/users/:user_id", users.DeleteUser)
	router.GET("/internal/users/search", users.SearchUser)
	router.POST("/users/login", users.Login)
}
