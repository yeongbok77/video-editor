package logic

import (
	"fmt"
	"testing"
)

func TestEditorProgress(t *testing.T) {
	res, err := EditorProgress("WeChat_20220527235010.mp4")
	if err != nil {
		fmt.Println("函数运行错误")
		t.Error(err)
	}
	if res != "100%" {
		t.Error("结果错误")
	}
}
