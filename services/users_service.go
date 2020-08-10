package services

import (
	"github.com/Kento75/bookstore_users-api/domain/users"
	"github.com/Kento75/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return nil, nil

	// return &user, nil

	// return &user, &errors.RestErr{
	// 	Status: http.StatusInternalServerError,
	// }
}

func GetUser() {

}

func FindUser() {

}
