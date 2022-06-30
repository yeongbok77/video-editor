package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"video-editor/logic"
)

//	VideoEditorHandler 视频剪辑接口
func VideoEditorHandler(c *gin.Context) {
	// 获取参数
	videoURL := c.Query("videoURL")
	StartTime := c.Query("StartTime")
	EndTime := c.Query("EndTime")

	_, err := logic.VideoEditor(videoURL, StartTime, EndTime)
	if err != nil {
		fmt.Println("剪辑错误", err)
	}

}
