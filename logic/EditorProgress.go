package logic

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var (
	progressComment string = "ffprobe   -v error   -select_streams v:0   -count_packets   -show_entries stream=nb_read_packets   -of csv=p=0  D:\\\\GO_WORK\\\\src\\\\video-editor\\\\public\\\\new\\\\"
)

func EditorProgress(fileName string) (resPercent string, err error) {
	// 追加写入文件名
	file, err := os.OpenFile(progressShellPath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("progressShell 文件打开失败", err)
		return
	}
	os.Truncate(progressShellPath, 0)
	file.Seek(0, 0)
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString(progressComment + fileName)
	write.Flush()

	//-----------------------------------------------------

	// 执行脚本
	cmd := exec.Command("C:\\Program Files\\Git\\bin\\bash", progressShellPath)
	bytes, err := cmd.Output()
	if err != nil {
		fmt.Println("cmd.Output:", err)
		return
	}
	str := strings.TrimRight(string(bytes), "\r\n")

	AllFrame, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf("strconv.Atoi err: %v", err)
		return
	}

	// 做计算
	// 先打开转码时记录数据的文件
	content, err := ioutil.ReadFile(progressPATH + fileName + ".txt")
	if err != nil {
		if err != nil {
			log.Fatalf("progressShell 文件打开失败 err:%v", err)
			return
		}
	}

	// 根据每行进行分割，得出最后一个 frame
	lines := strings.Split(string(content), "\n")
	Frame := lines[len(lines)-13]

	// 将 frame=xxx  进行分割，得出数值
	strFrameSlice := strings.Split(Frame, "=")
	fmt.Println(strFrameSlice)
	strFrame := strFrameSlice[len(strFrameSlice)-1]
	nowFrame, err := strconv.Atoi(strFrame)
	if err != nil {
		log.Fatalf("strconv.Atoi err: %v", err)
		return
	}

	// 计算 progress
	progressPercent := float64(nowFrame) / float64(AllFrame) * 100
	progressPercent, err = strconv.ParseFloat(fmt.Sprintf("%.2f", progressPercent), 64)
	if err != nil {
		log.Fatalf("strconv.ParseFloat err: %v", err)
		return
	}

	resPercent = strconv.FormatFloat(progressPercent, 'f', 2, 64)
	fmt.Println(resPercent + "%")

	return resPercent, nil
}
