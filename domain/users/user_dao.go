package users

import (
	"fmt"
	"strings"

	"github.com/Kento75/bookstore_users-api/datasources/mysql/users_db"
	"github.com/Kento75/bookstore_users-api/utils/date_utils"
	"github.com/Kento75/bookstore_users-api/utils/errors"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?,?,?,?);"
)

var (
	usersDB = make(map[int64]*User)
)

func something() {
	user := User{}
	if err := user.Get(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user.FirstName)
}

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	result := usersDB[user.Id]

	if result == nil {
		return errors.NotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	// first_name, last_name, email, date_created
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		// 一意制約違反の場合
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.BadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.InternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	userId, err := insertResult.LastInsertId()

	if err != nil {
		return errors.InternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	user.Id = userId

	return nil
}
