package users

import (
	"net/http"
	"strconv"

	"github.com/HunnTeRUS/bookstore_users-api/domain/users"
	users_services "github.com/HunnTeRUS/bookstore_users-api/services"
	"github.com/HunnTeRUS/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError("Invalid json body"))
		return
	}

	savedUser, err1 := users_services.CreateUser(user)

	if err1 != nil {
		c.JSON(err1.Code, err1)
		return
	}

	c.JSON(http.StatusCreated, savedUser)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if userErr != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError("Invalid user id"))
		return
	}

	savedUser, err1 := users_services.GetUser(userId)

	if err1 != nil {
		c.JSON(err1.Code, err1)
		return
	}

	c.JSON(http.StatusOK, savedUser)
	return
}
