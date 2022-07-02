package logic

import (
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"github.com/yeongbok77/video-editor/dao/mysql"
	"github.com/yeongbok77/video-editor/pkg/qiniu"
	"log"
	"strings"
)

var (
	tmpVideoPATH      string = "D:\\GO_WORK\\src\\video-editor\\public\\tmp\\"
	newVideoPATH      string = "D:\\GO_WORK\\src\\video-editor\\public\\new\\"
	progressPATH      string = "D:\\GO_WORK\\src\\video-editor\\public\\progress\\"
	progressShellPath string = "D:\\GO_WORK\\src\\video-editor\\controller\\progress.sh"
)

func VideoEditor(videoURL, StartTime, EndTime string, userId int64) (ResultVideoURL string, err error) {
	// 获取url中的文件名称
	splitURL := strings.Split(videoURL, "/")
	fileName := splitURL[len(splitURL)-1]

	// 视频中 StartTime 和 EndTime 的位置，可能不是关键帧，
	// 视频剪辑完后可能会出现起始和终止位置不准确、丢帧等情况。
	// 为了保证视频剪辑无误差，需要将所有的帧的编码方式转为帧内编码。
	//
	// 转为帧内编码的操作
	err = ffmpeg.Input(videoURL, ffmpeg.KwArgs{}).
		Output(tmpVideoPATH+fileName, ffmpeg.KwArgs{"strict": -2, "qscale": 0, "intra": ""}).OverWriteOutput().Run()
	if err != nil {
		log.Fatalf("转帧内编码操作错误：%v", err)
		return "", VideoEncodingErr
	}

	// 根据起始与终止时间，进行视频剪辑。

	err = ffmpeg.Input(tmpVideoPATH+fileName, ffmpeg.KwArgs{}).
		Output(newVideoPATH+fileName, ffmpeg.KwArgs{"ss": StartTime, "to": EndTime, "c": "copy", "progress": progressPATH + fileName + ".txt"}).
		OverWriteOutput().Run()
	if err != nil {
		log.Fatalf("视频剪辑错误: %v", err)
		return "", VideoEditErr
	}

	// 上传七牛云
	ResultVideoURL, err = qiniu.UploadVideo(newVideoPATH, fileName, userId)
	if err != nil {
		log.Fatalf("视频上传错误: %v", err)
		return "", UploadVideoErr
	}

	// 将剪辑记录写入MySQL
	err = mysql.RecordEditData(userId, fileName, ResultVideoURL)
	if err != nil {
		log.Fatalf("剪辑记录写入MySQL错误: %v", err)
		return "", RecordEditDataErr
	}

	return ResultVideoURL, err
}
