package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	// :8080のみだとダイアログ出てウザいので127.0.0.1:8080で指定
	router.Run("127.0.0.1:8080")
}
