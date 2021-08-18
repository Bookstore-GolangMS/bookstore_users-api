package app

import (
	"github.com/HunnTeRUS/bookstore_users-api/controllers/ping"
	"github.com/HunnTeRUS/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
	router.PUT("/users/:user_id", users.UpdateUser)
	router.PATCH("/users/:user_id", users.UpdateUser)
	router.DELETE("/users/:user_id", users.DeleteUser)
	router.DELETE("/internal/users/search", users.Search)
}
