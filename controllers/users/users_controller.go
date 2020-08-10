package users

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Kento75/bookstore_users-api/domain/users"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
	bytes, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		// TODO: handler error
		return
	}

	// jsonをパース
	if err := json.Unmarshal(bytes, &user); err != nil {
		// TODO: handler error
		return
	}

	c.String(http.StatusNotImplemented, "Implement me !")
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me !")
}
