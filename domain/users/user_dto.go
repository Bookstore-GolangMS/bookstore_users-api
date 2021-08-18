package users

import (
	"strings"

	"github.com/HunnTeRUS/bookstore_users-api/utils/errors"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Status      string `json:"status"`
	Password    string `json:"password"`
	DateCreated string `json:"date_created"`
}

type Users []User

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.LastName = strings.TrimSpace(strings.ToLower(user.LastName))
	user.FirstName = strings.TrimSpace(strings.ToLower(user.FirstName))

	if user.Email == "" {
		return errors.NewBadRequestError("Invalid email address")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.NewBadRequestError("Invalid password")
	}

	return nil
}
