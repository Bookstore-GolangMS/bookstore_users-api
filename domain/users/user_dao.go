package users

import (
	"fmt"

	"github.com/HunnTeRUS/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Save() *errors.RestErr {
	current := usersDB[user.Id]

	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("User %d already exists", user.Id))
	}

	usersDB[user.Id] = user
	return nil
}

func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]

	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.DateCreated = result.DateCreated
	user.Email = result.Email
	user.LastName = result.LastName

	return nil
}
