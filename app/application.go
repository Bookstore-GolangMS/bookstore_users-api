package app

import (
	"github.com/HunnTeRUS/bookstore_users-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	logger.Info("abount to start the application...")
	router.Run(":8080")
}
