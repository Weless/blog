package controller

import (
	"blog/dto"
	"blog/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ArticleDetailController(c *gin.Context) {
	articleId, _ := strconv.Atoi(c.Param("articleId"))
	if articleId <= 0 {
		response.Fail(c, "no data", nil)
		return
	}
	article := dto.GetArticleById(articleId)
	if len(article) == 0 {
		response.Fail(c, "no data", nil)
		return
	}
	response.Success(c, "success", gin.H{"article": article})
}
