package router

import (
	"github.com/dingdinglz/test-blog/config"
	"github.com/dingdinglz/test-blog/handlers"
	"github.com/dingdinglz/test-blog/middleware"
	"github.com/gin-gonic/gin"
)

// SetupRouter 配置路由
func SetupRouter() *gin.Engine {
	// 设置Gin模式
	gin.SetMode(config.AppConfig.Server.Mode)

	// 创建路由引擎
	r := gin.New()

	// 使用全局中间件
	r.Use(middleware.CORS())
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	// API路由组
	api := r.Group("/api")
	{
		// 公开路由 - 不需要认证
		api.POST("/register", handlers.Register)
		api.POST("/login", handlers.Login)

		// 公开的文章查询接口
		api.GET("/articles", handlers.GetAllArticles)
		api.GET("/articles/user/:user_id", handlers.GetArticlesByUser)

		// 需要认证的路由
		auth := api.Group("")
		auth.Use(middleware.AuthMiddleware())
		{
			// 用户相关
			auth.GET("/user/info", handlers.GetInfo)

			// 文章相关
			auth.POST("/articles", handlers.CreateArticle)
			auth.PUT("/articles/:id", handlers.UpdateArticle)
			auth.DELETE("/articles/:id", handlers.DeleteArticle)
		}
	}

	return r
}
