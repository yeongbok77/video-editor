package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func SetUpRouter() *gin.Engine {
	r := gin.New()

	r.GET("/video-editor/editor", VideoEditorHandler)

	return r
}

func Test_helloHandler(t *testing.T) {
	videoURL := "http://re8twfm01.hd-bkt.clouddn.com/original/WeChat_20220526160454.mp4"
	StartTime := "00:00:05"
	EndTime := "00:00:10"

	userId := int64(10)
	resultURL := "http://re8twfm01.hd-bkt.clouddn.com/ClipVideo/" + strconv.FormatInt(userId, 10) + "/WeChat_20220526160454.mp4"

	r := SetUpRouter()
	//将项目中的API注册到测试使用的router
	r.GET("/recipes", VideoEditorHandler)
	//向注册的路有发起请求
	req, err := http.NewRequest("GET", "http://127.0.0.1:8080/video-editor/editor?"+
		"videoURL="+videoURL+"&"+
		"StartTime="+StartTime+"&"+
		"EndTime="+EndTime, nil)
	w := httptest.NewRecorder()
	//模拟http服务处理请求
	r.ServeHTTP(w, req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, resultURL, w.Body.String())

}
