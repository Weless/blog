package main

import (
	"blog/controller"
	"blog/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.GET("/api/", controller.IndexHandler)
	r.GET("/api/articles", controller.ArticleHandler)

	// 管理端用
	r.GET("/api/totalLength", controller.GetTotalLengthHandler)
	r.GET("/api/articles/detail/:articleId", controller.ArticleDetailController)

	r.GET("/api/categories", controller.CategoriesHandler)
	r.GET("/api/categories/:categoryId", controller.CategoryArticleHandler)

	r.GET("/api/latest", controller.LatestArticleHandler)

	r.GET("/api/archive", controller.ArchiveHandler)
	r.GET("/api/archive/:year/:month", controller.ArchiveArticleHandler)
	// 平台接口

	// 获取文章详情
	r.GET("/api/articleDetail", controller.ArticleDetailForManagementHandler)
	// 获取文章列表
	r.GET("/api/articleList", controller.ArticleListHandler)
	// 获取分类列表
	r.GET("/api/categoryList", controller.GetCategoryListHandler)
	// 删除文章
	r.POST("/api/deleteArticle", controller.DeleteHandler)
	// 插入文章
	r.POST("/api/insertArticle", controller.InsertHandler)
	// 修改文章
	r.POST("/api/updateArticle", controller.UpdateHandler)
	// 插入分类
	r.POST("/api/insertCategory", controller.InsertCategoryHandler)

	// 创建用户
	r.POST("/api/auth/register", controller.RegisterHandler)
	// 登录
	r.POST("/api/auth/login", controller.LoginHandler)

	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.UserInfoHandler)
	return r
}
