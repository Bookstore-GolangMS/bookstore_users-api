package users

import (
	"fmt"

	usersdb "github.com/HunnTeRUS/bookstore_users-api/datasources/mysql/users_db"
	dateutils "github.com/HunnTeRUS/bookstore_users-api/utils/date_utils"
	"github.com/HunnTeRUS/bookstore_users-api/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?)"
)

func (user *User) Save() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()
	user.DateCreated = dateutils.GetNowString()

	result, err := usersdb.Client.Exec(queryInsertUser, user.FirstName, user.LastName, user.Email, user.DateCreated)

	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error trying to save user: %s", err.Error()))
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to get last inserted id: %s", err.Error()))
	}

	user.Id = userId

	return nil
}

func (user *User) Get() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryInsertUser)
	if err == nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	result, err := usersdb.Client.Exec(queryInsertUser, user.FirstName, user.LastName, user.Email, user.DateCreated)

	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error trying to save user: %s", err.Error()))
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to get last inserted id: %s", err.Error()))
	}

	user.Id = userId

	return nil
}
