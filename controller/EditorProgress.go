package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yeongbok77/video-editor/logic"
	"log"
	"net/http"
)

// VideoEditorProgress 获取视频剪辑进度
func VideoEditorProgress(c *gin.Context) {
	// 获取文件名
	fileName := c.Query("fileName")

	progressPercent, err := logic.EditorProgress(fileName)
	if err != nil {
		log.Fatalf("logic.EditorProgress 业务内部错误: %v", err)
		c.JSON(http.StatusOK, CodeMsgMap[CodeServerError])
		return
	}

	c.String(http.StatusOK, progressPercent)

}
