package main

import (
	"log"

	"github.com/dingdinglz/test-blog/config"
	"github.com/dingdinglz/test-blog/database"
	"github.com/dingdinglz/test-blog/router"
)

func main() {
	// 加载配置
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("配置加载失败: %v", err)
	}

	// 初始化数据库
	if err := database.Init(); err != nil {
		log.Fatalf("数据库初始化失败: %v", err)
	}

	// 设置路由
	r := router.SetupRouter()

	// 启动服务器
	port := ":" + config.AppConfig.Server.Port
	log.Printf("服务器启动成功，监听端口: %s\n", port)
	log.Printf("访问地址: http://localhost%s/api\n", port)

	if err := r.Run(port); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
