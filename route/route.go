package route

import (
	"github.com/gin-gonic/gin"
	"github.com/yeongbok77/video-editor/controller"
)

func SetUpRouter() *gin.Engine {
	r := gin.New()

	r.GET("/video-editor/editor", controller.VideoEditorHandler)

	return r
}
