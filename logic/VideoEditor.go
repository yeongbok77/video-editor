package logic

import (
	"encoding/json"
	"fmt"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"github.com/yeongbok77/video-editor/dao/mysql"
	"github.com/yeongbok77/video-editor/pkg/qiniu"
	"log"
	"math/rand"
	"net"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	tmpVideoPATH string = "D:\\GO_WORK\\src\\video-editor\\public\\tmp\\"
	newVideoPATH string = "D:\\GO_WORK\\src\\video-editor\\public\\new\\"
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
	//a, err := ffmpeg.Probe(newVideoPATH + fileName)
	//if err != nil {
	//	log.Fatalf("ffmpeg.Probe err: %v", err)
	//	return "", err
	//}
	//totalDuration, err := probeDuration(a)
	//if err != nil {
	//	log.Fatalf("probeDuration err: %v", err)
	//	return "", err
	//}

	err = ffmpeg.Input(tmpVideoPATH+fileName, ffmpeg.KwArgs{}).
		Output(newVideoPATH+fileName, ffmpeg.KwArgs{"ss": StartTime, "to": EndTime, "c": "copy", "progress": "D:\\GO_WORK\\src\\video-editor\\public\\progress\\progress.txt"}).
		//GlobalArgs("-progress", TempSock(totalDuration)).
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

func probeDuration(a string) (float64, error) {
	pd := probeData{}
	err := json.Unmarshal([]byte(a), &pd)
	if err != nil {
		return 0, err
	}
	f, err := strconv.ParseFloat(pd.Format.Duration, 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}

type probeFormat struct {
	Duration string `json:"duration"`
}

type probeData struct {
	Format probeFormat `json:"format"`
}

func TempSock(totalDuration float64) string {
	// serve

	rand.Seed(time.Now().Unix())
	sockFileName := path.Join(os.TempDir(), fmt.Sprintf("%d_sock", rand.Int()))
	l, err := net.Listen("unix", sockFileName)
	if err != nil {
		panic(err)
	}

	go func() {
		re := regexp.MustCompile(`out_time_ms=(\d+)`)
		fd, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}
		buf := make([]byte, 16)
		data := ""
		progress := ""
		for {
			_, err := fd.Read(buf)
			if err != nil {
				return
			}
			data += string(buf)
			a := re.FindAllStringSubmatch(data, -1)
			cp := ""
			if len(a) > 0 && len(a[len(a)-1]) > 0 {
				c, _ := strconv.Atoi(a[len(a)-1][len(a[len(a)-1])-1])
				cp = fmt.Sprintf("%.2f", float64(c)/totalDuration/1000000)
			}
			if strings.Contains(data, "progress=end") {
				cp = "done"
			}
			if cp == "" {
				cp = ".0"
			}
			if cp != progress {
				progress = cp
				fmt.Println("progress: ", progress)
			}
		}
	}()

	return sockFileName
}
