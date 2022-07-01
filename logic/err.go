package logic

import (
	"errors"
)

var (
	InvalidVideoURL   = errors.New("视频URL参数不合法")
	VideoEncodingErr  = errors.New("视频编码错误")
	VideoEditErr      = errors.New("视频剪辑错误")
	UploadVideoErr    = errors.New("视频上传错误")
	RecordEditDataErr = errors.New("剪辑记录写入MySQL错误")
)
