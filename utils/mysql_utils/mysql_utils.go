package mysql_utils

import (
	"errors"
	"strings"

	"github.com/Kento75/bookstore_utils-go/rest_errors"
	"github.com/go-sql-driver/mysql"
)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) *rest_errors.RestErr {

	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return rest_errors.NewNotFoundError("no record matching given id")
		}
		return rest_errors.NewInternalServerError("error parsing database response", err)
	}
	switch sqlErr.Number {
	// 一意制約違反
	case 1062:
		return rest_errors.NewBadRequestError("invalid data")
	}

	// 原因不明のDBエラー(上で定義していないやつ)
	return rest_errors.NewInternalServerError("error processing request", errors.New("database error"))
}
