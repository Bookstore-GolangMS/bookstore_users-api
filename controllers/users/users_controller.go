package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Bookstore-GolangMS/bookstore_oauth-go/oauth"
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

	savedUser, err1 := users_services.UsersService.CreateUser(user)

	if err1 != nil {
		c.JSON(err1.Code, err1)
		return
	}

	c.JSON(http.StatusCreated, savedUser.Marshal(c.GetHeader("X-Public") == "true"))
}

func GetUser(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err)
		return
	}

	if callerId := oauth.GetCallerId(c.Request); callerId == 0 {
		err := errors.RestErr{
			Code:    http.StatusUnauthorized,
			Message: "resource not available",
		}
		c.JSON(err.Code, err)
		return
	}

	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if userErr != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError("Invalid user id"))
		return
	}

	user, err1 := users_services.UsersService.GetUser(userId)

	if err1 != nil {
		c.JSON(err1.Code, err1)
		return
	}

	if oauth.GetCallerId(c.Request) == user.Id {
		c.JSON(http.StatusOK, user.Marshal(false))
		return
	}

	c.JSON(http.StatusOK, user.Marshal(oauth.IsPublic(c.Request)))
}

func UpdateUser(c *gin.Context) {
	userId, userErr := getUserID(c)

	if userErr != nil {
		c.JSON(userErr.Code, userErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError("Invalid json body"))
		return
	}
	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, err := users_services.UsersService.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, result.Marshal(c.GetHeader("X-Public") == "true"))
}

func DeleteUser(c *gin.Context) {
	userId, userErr := getUserID(c)

	if userErr != nil {
		c.JSON(userErr.Code, userErr)
		return
	}

	if err := users_services.UsersService.DeleteUser(userId); err != nil {
		c.JSON(userErr.Code, userErr)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "delete"})
}

func Login(c *gin.Context) {
	login_request := &users.LoginRequest{}
	if err := c.ShouldBindJSON(&login_request); err != nil {
		func_err := errors.NewBadRequestError("json body is malformatted")
		c.JSON(func_err.Code, func_err)
		return
	}

	user, err := users_services.UsersService.LoginUser(*login_request)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, user.Marshal(c.GetHeader("X-Public") == "true"))
}

func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := users_services.UsersService.Search(status)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, users.Marshal(c.GetHeader("X-Public") == "true"))
}

func getUserID(c *gin.Context) (int64, *errors.RestErr) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError(fmt.Sprintf("user id must be a number"))
	}

	return userId, nil
}
