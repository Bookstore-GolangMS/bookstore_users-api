package users_services

import (
	"github.com/Bookstore-GolangMS/bookstore_utils-go/errors"
	"github.com/HunnTeRUS/bookstore_users-api/domain/users"
	"github.com/HunnTeRUS/bookstore_users-api/utils/crypto_utils"
	dateutils "github.com/HunnTeRUS/bookstore_users-api/utils/date_utils"
)

type usersService struct {
	usersServiceInterface
}

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestErr)
	Search(string) (users.Users, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	GetUser(int64) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	LoginUser(users.LoginRequest) (*users.User, *errors.RestErr)
}

func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.DateCreated = dateutils.GetNowDBString()

	user.Password = crypto_utils.GetMd5(user.Password)

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *usersService) LoginUser(login_request users.LoginRequest) (*users.User, *errors.RestErr) {
	user := &users.User{
		Email:    login_request.Email,
		Password: crypto_utils.GetMd5(login_request.Password),
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.FindByEmailAndPassword(); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *usersService) Search(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.Search(status)
}

func (s *usersService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := users.User{Id: userId}

	if err := result.Get(); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := UsersService.GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.Email = user.Email
		current.LastName = user.LastName
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func (s *usersService) DeleteUser(userId int64) *errors.RestErr {
	user := users.User{Id: userId}

	return user.Delete()
}
