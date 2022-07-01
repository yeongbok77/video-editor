package mysql

import (
	"github.com/yeongbok77/video-editor/model"
	"log"
)

func RecordEditData(userId int64, fileName string, resultVideoURL string) (err error) {
	videoEditData := &model.VideoEditData{
		UserId:         userId,
		FileName:       fileName,
		ResultVideoURL: resultVideoURL,
	}
	if err = db.Select("user_id", "file_name", "result_video_url").Create(videoEditData).Error; err != nil {
		log.Fatalf("db.Select  数据插入失败: %v", err)
		return
	}

	return nil
}
