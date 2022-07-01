package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yeongbok77/video-editor/logic"
	"log"
	"net/http"
	"regexp"
	"strings"
)

//	VideoEditorHandler 视频剪辑接口
func VideoEditorHandler(c *gin.Context) {
	// 获取参数
	videoURL := c.Query("videoURL")
	StartTime := c.Query("StartTime")
	EndTime := c.Query("EndTime")
	//UserId, ok := c.Get("userId")
	//if !ok {
	//	log.Fatalln("userId 获取错误")
	//	c.JSON(http.StatusOK, codeMsgMap[CodeServerError])
	//	return
	//}
	var UserId int64 = 2

	// 参数校验
	re := regexp.MustCompile("(http|https):\\/\\/[\\w\\-_]+(\\.[\\w\\-_]+)+([\\w\\-\\.,@?^=%&:/~\\+#]*[\\w\\-\\@?^=%&/~\\+#])?")
	result := re.FindAllStringSubmatch(videoURL, -1)
	if result == nil {
		log.Fatalln("url不合法")
		c.JSON(http.StatusOK, codeMsgMap[CodeInvalidVideoURL])
		return
	}
	// 视频剪辑时，用户使用类似拖动进度条的形式，来选取起始和终止时间。（参考十行笔记的剪辑方式）
	// 所以时间参数的格式需在前端定义，并且在前端应把起始和终止时间控制在视频长度以内。
	// 所以后端只校验格式是否正确。
	rStart := strings.Split(StartTime, ":")
	rEnd := strings.Split(EndTime, ":")
	if len(rStart) != 3 || len(rEnd) != 3 {
		log.Fatalln("时间参数不合法")
		c.JSON(http.StatusOK, codeMsgMap[CodeInvalidTime])
		return
	}

	// 业务逻辑
	ResultVideoURL, err := logic.VideoEditor(videoURL, StartTime, EndTime, UserId)
	if err != nil {
		log.Fatalf("logic.VideoEditor 业务内部错误: %v", err)
		c.JSON(http.StatusOK, codeMsgMap[CodeEditError])
		return
	}

	// 操作成功
	c.JSON(http.StatusOK, ResultVideoURL)
}
