package interfaces

import (
	"github.com/HunnTeRUS/bookstore_users-api/domain/users"
	"github.com/HunnTeRUS/bookstore_users-api/utils/errors"
)

type User interface {
	UserMethods
}

type UserMethods interface {
	CreateUser(user users.User) (*users.User, *errors.RestErr)
}
