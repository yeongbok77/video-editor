package logic

import (
	"fmt"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"github.com/yeongbok77/video-editor/pkg/qiniu"
	"log"
	"strings"
)

var (
	tmpVideoPATH string = "D:\\GO_WORK\\src\\video-editor\\public\\tmp\\"
	newVideoPATH string = "D:\\GO_WORK\\src\\video-editor\\public\\new\\"
)

func VideoEditor(videoURL, StartTime, EndTime string) (ResultVideoURL string, err error) {
	// userId 方便测试
	userId := int64(10)

	// 获取url中的文件名称
	splitURL := strings.Split(videoURL, "/")
	fileName := splitURL[len(splitURL)-1]
	fmt.Println("视频文件名为：", fileName)

	// 视频中 StartTime 和 EndTime 的位置，可能不是关键帧，
	// 视频剪辑完后可能会出现起始和终止位置不准确、丢帧等情况。
	// 为了保证视频剪辑无误差，需要将所有的帧的编码方式转为帧内编码。
	//
	// 转为帧内编码的操作
	err = ffmpeg.Input(videoURL, ffmpeg.KwArgs{}).
		Output(tmpVideoPATH+fileName, ffmpeg.KwArgs{"strict": -2, "qscale": 0, "intra": ""}).OverWriteOutput().Run()
	if err != nil {
		log.Printf("转帧内编码操作错误：%v", err)
		return
	}

	// 根据起始与终止时间，进行视频剪辑。
	err = ffmpeg.Input(tmpVideoPATH+fileName, ffmpeg.KwArgs{}).
		Output(newVideoPATH+fileName, ffmpeg.KwArgs{"ss": StartTime, "to": EndTime, "c": "copy"}).OverWriteOutput().Run()
	if err != nil {
		log.Println("视频剪辑错误", err)

	}

	// 将剪辑记录写入MySQL
	//mysql.RecordEditData()
	// 上传七牛云
	ResultVideoURL, err = qiniu.UploadVideo(newVideoPATH, fileName, userId)
	// 返回剪辑后的视频URL

	return ResultVideoURL, err
}
