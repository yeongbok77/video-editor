package main

import (
	"github.com/yeongbok77/video-editor/dao/mysql"
	"github.com/yeongbok77/video-editor/route"
	"github.com/yeongbok77/video-editor/settings"
)

func main() {
	// 初始化配置文件
	settings.Init()
	// 初始化MySQL
	mysql.Init(settings.Conf.MySQLConfig)

	r := route.SetUpRouter()
	// 启动服务
	r.Run(":8080")
}
