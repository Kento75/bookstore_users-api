package mysql_utils

import (
	"strings"

	"github.com/Kento75/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {

	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NotFoundError("no record matching given id")
		}
		return errors.InternalServerError("error parsing database response")
	}
	switch sqlErr.Number {
	// 一意制約違反
	case 1062:
		return errors.BadRequestError("invalid data")
	}
	return errors.InternalServerError("error processing request")
}
