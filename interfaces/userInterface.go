package interfaces

import (
	"github.com/Bookstore-GolangMS/bookstore_utils-go/errors"
	"github.com/HunnTeRUS/bookstore_users-api/domain/users"
)

type User interface {
	UserMethods
}

type UserMethods interface {
	CreateUser(user users.User) (*users.User, *errors.RestErr)
}
