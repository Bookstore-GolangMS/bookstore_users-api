package users

import (
	"fmt"
	"strings"

	usersdb "github.com/HunnTeRUS/bookstore_users-api/datasources/mysql/users_db"
	"github.com/HunnTeRUS/bookstore_users-api/logger"
	"github.com/HunnTeRUS/bookstore_users-api/utils/errors"
	"github.com/HunnTeRUS/bookstore_users-api/utils/mysql_utils"
)

const (
	queryUpdateUser                 = "UPDATE users SET first_name = ?, last_name = ?, email = ? WHERE id = ?"
	queryInsertUser                 = "INSERT INTO users(first_name, last_name, email, date_created, password, status) VALUES(?, ?, ?, ?, ?, ?)"
	queryGetUser                    = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id = ?"
	queryDeleteUser                 = "DELETE FROM users WHERE id = ?"
	queryFindUserByStatus           = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status = ?"
	queryFindUserByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email = ? AND password = ? AND status = ?;"
)

func (user *User) Save() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("Error trying to prepare insert statement", err)
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	result, saveErr := usersdb.Client.Exec(queryInsertUser, user.FirstName, user.LastName, user.Email,
		user.DateCreated, user.Password, user.Status)

	if saveErr != nil {
		logger.Error("Error trying to save user", err)
		return mysql_utils.ParseError(saveErr)
	}

	userId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error trying to get last inserted id", err)
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to get last inserted id: %s", err.Error()))
	}

	user.Id = userId

	return nil
}

func (user *User) Get() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("Error trying to prepare insert statement", err)
		return errors.NewInternalServerError("database error")
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email,
		&user.DateCreated, &user.Status); getErr != nil {
		logger.Error("Error trying to scan result from database", err)
		return mysql_utils.ParseError(getErr)
	}

	return nil
}

func (user *User) FindByEmailAndPassword() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryFindUserByEmailAndPassword)
	if err != nil {
		logger.Error("Error trying to prepare get user by email and password statement", err)
		return errors.NewInternalServerError("database error")
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.Email, user.Password, StatusActive)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email,
		&user.DateCreated, &user.Status); getErr != nil {

		if strings.Contains(getErr.Error(), mysql_utils.ErrorNoRows) {
			return errors.NewNotFoundError("invalid user credentials")
		}
		logger.Error("Error trying to scan result from database when getting user by email and password", err)
		return mysql_utils.ParseError(getErr)
	}

	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("Error trying to prepare update statement", err)
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	_, updateErr := usersdb.Client.Exec(queryUpdateUser, user.FirstName, user.LastName, user.Email, user.Id)
	if updateErr != nil {
		logger.Error("Error trying to update user on database", err)
		return mysql_utils.ParseError(updateErr)
	}

	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("Error trying to prepare delete statement", err)
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	_, deleteErr := usersdb.Client.Exec(queryDeleteUser, user.Id)
	if deleteErr != nil {
		logger.Error("Error trying to delete user from database", err)
		return mysql_utils.ParseError(deleteErr)
	}

	return nil
}

func (user *User) Search(status string) ([]User, *errors.RestErr) {
	stmt, err := usersdb.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("Error trying to prepare search statement", err)
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := usersdb.Client.Query(status)
	if err != nil {
		logger.Error("Error trying to search user in database based on status", err)
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName,
			&user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("Error trying to scan result from database", err)
			return nil, errors.NewInternalServerError(err.Error())
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}

	return results, nil
}
