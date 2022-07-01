package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yeongbok77/video-editor/logic"
	"net/http"
)

//	VideoEditorHandler 视频剪辑接口
func VideoEditorHandler(c *gin.Context) {
	// 获取参数
	videoURL := c.Query("videoURL")
	StartTime := c.Query("StartTime")
	EndTime := c.Query("EndTime")

	ResultVideoURL, err := logic.VideoEditor(videoURL, StartTime, EndTime)
	if err != nil {
		fmt.Println("剪辑错误", err)
		c.JSON(http.StatusOK, "视频剪辑失败,请重试！")
	}
	c.JSON(http.StatusOK, ResultVideoURL)
}
