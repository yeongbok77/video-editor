package route

import (
	"github.com/gin-gonic/gin"
	"github.com/yeongbok77/video-editor/controller"
	middlewares "github.com/yeongbok77/video-editor/middleware"
	"net/http"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLFiles("D:\\GO_WORK\\src\\video-editor\\route\\editor.html")

	r.GET("/video-editor/editor", func(c *gin.Context) {
		c.HTML(http.StatusOK, "editor.html", nil)
	})

	r.POST("/video-editor/editor", middlewares.JWTAuthMiddleware(), controller.VideoEditorHandler)

	r.GET("/video-editor/editor/progress", middlewares.JWTAuthMiddleware(), controller.VideoEditorProgress)

	return r
}
