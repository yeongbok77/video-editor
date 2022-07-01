package qiniu

import (
	"log"
	"net/http"
	"testing"
)

func TestQiniu(t *testing.T) {
	videoPath := "D:\\GO_WORK\\src\\video-editor\\public\\new\\"
	fileName := "WeChat_20220527235010.mp4"
	newVideoURL, err := UploadVideo(videoPath, fileName, int64(6))
	if err != nil {
		log.Println("上传错误", err)
		t.Error(err)
	}
	resp, err := http.Get(newVideoURL)
	if err != nil {
		log.Println("http请求错误", err)
		t.Error(err)
	}
	if resp.StatusCode != 200 {
		t.Error("test failed！")
	}
}
