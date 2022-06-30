package logic

import (
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
)

func VideoEditor(videoURL, StartTime, EndTime string) (ResultVideoURL string, err error) {
	//
	videoURL = "http://re8twfm01.hd-bkt.clouddn.com/original/WeChat_20220527235010.mp4"
	StartTime = "00:00:05"
	EndTime = "00:00:11"

	//
	err = ffmpeg.Input(videoURL, ffmpeg.KwArgs{}).
		Output("D:\\GO_WORK\\src\\video-editor\\public\\tmp\\tmpWeChat_20220527235010.mp4", ffmpeg.KwArgs{"strict": -2, "qscale": 0, "intra": ""}).OverWriteOutput().Run()

	// 剪辑
	err = ffmpeg.Input("D:\\GO_WORK\\src\\video-editor\\public\\tmp\\tmpWeChat_20220527235010.mp4", ffmpeg.KwArgs{}).
		Output("D:\\GO_WORK\\src\\video-editor\\public\\new\\newWeChat_20220527235010.mp4", ffmpeg.KwArgs{"ss": StartTime, "to": EndTime, "c": "copy"}).OverWriteOutput().Run()
	if err != nil {
		log.Println("剪辑错误", err)

	}

	return "", err
}
