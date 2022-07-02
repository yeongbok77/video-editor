package route

import (
	"github.com/gin-gonic/gin"
	"github.com/yeongbok77/video-editor/controller"
	"net/http"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLFiles("D:\\GO_WORK\\src\\video-editor\\route\\editor.html")

	r.GET("/video-editor/editor", func(c *gin.Context) {
		c.HTML(http.StatusOK, "editor.html", nil)
	})

	r.POST("/video-editor/editor", controller.VideoEditorHandler)

	r.GET("/video-editor/editor/progress", controller.VideoEditorProgress)

	return r
}
