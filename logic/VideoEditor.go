package logic

import (
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
	"strings"
)

var (
	tmpVideoPATH string = ""
	newVideoPATH string = ""
)

func VideoEditor(videoURL, StartTime, EndTime string) (ResultVideoURL string, err error) {

	splitURL := strings.Split(videoURL, "/")
	fileName := splitURL[len(splitURL)-1]

	err = ffmpeg.Input(videoURL, ffmpeg.KwArgs{}).
		Output(tmpVideoPATH+fileName, ffmpeg.KwArgs{"strict": -2, "qscale": 0, "intra": ""}).OverWriteOutput().Run()

	// 剪辑
	err = ffmpeg.Input(tmpVideoPATH+fileName, ffmpeg.KwArgs{}).
		Output(newVideoPATH+fileName, ffmpeg.KwArgs{"ss": StartTime, "to": EndTime, "c": "copy"}).OverWriteOutput().Run()
	if err != nil {
		log.Println("剪辑错误", err)

	}

	return "", err
}
