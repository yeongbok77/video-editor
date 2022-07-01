package model

type VideoEditData struct {
	Id             int64  `gorm:"primaryKey; type:bigint(20) AUTO_INCREMENT"`
	UserId         int64  `gorm:"uniqueIndex:idx_userId; type:bigint(20) UNIQUE NOT NULL"`
	FileName       string `gorm:"type:varchar(255) NOT NULL"`
	ResultVideoURL string `gorm:"type:varchar(255) NOT NULL"`
	CreatedAt      JSONTime
}
