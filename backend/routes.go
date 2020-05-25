package main

import (
	"blog/controller"
	"blog/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.GET("/",controller.IndexHandler)
	r.GET("/articles",controller.ArticleHandler)
	r.GET("/totalLength",controller.GetTotalLengthHandler)
	r.GET("/articles/detail/:articleId",controller.ArticleDetailController)

	r.GET("/categories",controller.CategoriesHandler)
	r.GET("/categories/:categoryId",controller.CategoryArticleHandler)


	r.GET("/labels",controller.LabelHandler)
	r.GET("/labels/:labelId",controller.LabelArticleHandler)


	r.GET("/latest",controller.LatestArticleHandler)

	r.GET("/archive",controller.ArchiveHandler)
	r.GET("/archive/:year/:month",controller.ArchiveArticleHandler)

	// 平台接口
	// 删除
	r.POST("/deleteArticle",controller.DeleteHandler)
	// 插入
	r.POST("/insertArticle",controller.InsertHandler)
	// 修改
	r.POST("/updateArticle",controller.UpdateHandler)
	return r
}
