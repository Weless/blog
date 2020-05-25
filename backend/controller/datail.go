package controller

import (
	"blog/dto"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ArticleDetailController(c *gin.Context){
	articleId, _:= strconv.Atoi(c.Param("articleId"))
	if articleId <= 0{
		c.JSON(404,gin.H{
			"code":404,
			"msg":"no data",
		})
		return
	}
	article := dto.GetArticleById(articleId)
	if len(article) == 0{
		c.JSON(404,gin.H{
			"code":404,
			"msg":"no data",
		})
		return
	}
	c.JSON(200,gin.H{
		"code":200,
		"msg":"success",
		"article":article,
	})
}