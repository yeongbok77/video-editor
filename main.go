package main

import "video-editor/route"

func main() {
	// 初始化MySQL

	// 初始化日志库

	// 初始化配置文件

	// 启动服务
	r := route.SetUpRouter()

	r.Run(":8080")
}
