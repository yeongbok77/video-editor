package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetUpRouter2() *gin.Engine {
	r := gin.New()

	r.GET("/video-editor/editor/progress", VideoEditorHandler)

	return r
}

func Test_EditorProgressHandler(t *testing.T) {
	r := SetUpRouter2()

	//向注册的路有发起请求
	req, err := http.NewRequest("GET", "http://127.0.0.1:8080/video-editor/editor/progress?fileName=WeChat_20220527235010.mp4", nil)
	w := httptest.NewRecorder()
	//模拟http服务处理请求
	r.ServeHTTP(w, req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
}
