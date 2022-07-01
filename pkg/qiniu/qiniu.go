package qiniu

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"strconv"
)

var (
	AccessKey        = "PMOpO_-mep9f8MOY-WJyp8qyJLpFCAJIahoJ7VXR"
	SerectKey        = "av1aSz1oxNLe4M2BSOizs5awbhHvYFVCpP-HZLAf"
	Bucket           = "videoclipstore"               // bucket name
	StorageUrl       = "re8twfm01.hd-bkt.clouddn.com" // 域名
	ErrorQiniuFailed = errors.New("七牛：视频上传失败")
)

func UploadVideo(videoPath, fileName string, userId int64) (newVideoURL string, err error) {
	// 鉴权
	mac := qbox.NewMac(AccessKey, SerectKey)

	input := []byte("videoclipstore:ClipVideo/" + strconv.FormatInt(userId, 10) + "/" + fileName)

	// 演示base64编码
	encodeString := base64.StdEncoding.EncodeToString(input)

	// 上传策略
	putPolicy := storage.PutPolicy{
		Scope:         Bucket,
		Expires:       7200,
		PersistentOps: "avthumb/mp4|saveas/" + encodeString,
	}

	// 获取上传token
	upToken := putPolicy.UploadToken(mac)

	// 上传Config对象
	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuadong //指定上传的区域
	cfg.UseHTTPS = false            // 是否使用https域名
	cfg.UseCdnDomains = false       //是否使用CDN上传加速

	// 需要上传的文件
	localFile := videoPath + fileName

	// 七牛key
	qiniuKey := "ClipVideo/" + strconv.FormatInt(userId, 10) + "/" + fileName

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 上传文件
	err = formUploader.PutFile(context.Background(), &ret, upToken, qiniuKey, localFile, nil)
	if err != nil {
		fmt.Println("上传文件失败,原因:", err)
		return
	}
	fmt.Println("上传成功,key为:", ret.Key)
	// 返回上传后的文件访问路径
	newVideoURL = "http://" + StorageUrl + "/" + ret.Key
	fmt.Println("视频访问路径为：", newVideoURL)
	return newVideoURL, err
}
