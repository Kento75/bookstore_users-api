package app

import "github.com/Kento75/bookstore_users-api/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
