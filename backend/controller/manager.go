package controller

import (
	"blog/dto"
	"github.com/gin-gonic/gin"
	"strconv"
)

func DeleteHandler(c *gin.Context)  {
	articleId, _ := strconv.Atoi(c.PostForm("articleId"))
	if articleId <= 0 {
		c.JSON(400,gin.H{
			"code":400,
			"msg":"request failed",
		})
		return
	}

	if !dto.SearchId(articleId){
		c.JSON(400,gin.H{
			"code":400,
			"msg":"request failed",
		})
		return
	}
	dto.DeleteArticle(articleId)

	c.JSON(200,gin.H{
		"code":200,
		"msg":"success",
	})
	return
}

func InsertHandler(c *gin.Context)  {
	title := c.PostForm("title")
	content := c.PostForm("content")
	summary := c.PostForm("summary")
	categoryId, _ := strconv.Atoi(c.PostForm("categoryId"))
	labelId, _ := strconv.Atoi(c.PostForm("labelId"))
	if categoryId <= 0 || labelId <= 0{
		c.JSON(400,gin.H{
			"code":400,
			"msg":"request failed",
		})
		return
	}
	dto.InsertArticle(title,content,summary,categoryId,labelId)

	c.JSON(200,gin.H{
		"code":200,
		"msg":"success",
	})
	return
}

func UpdateHandler(c *gin.Context)  {
	articleId,_ := strconv.Atoi(c.PostForm("articleId"))
	title := c.PostForm("title")
	content := c.PostForm("content")
	summary := c.PostForm("summary")
	categoryId, _ := strconv.Atoi(c.PostForm("categoryId"))
	labelId, _ := strconv.Atoi(c.PostForm("labelId"))

	if !dto.SearchId(articleId){
		c.JSON(400,gin.H{
			"code":400,
			"msg":"request failed",
		})
		return
	}

	dto.UpdateArticle(articleId, title,content,summary,categoryId,labelId)

	c.JSON(200,gin.H{
		"code":200,
		"msg":"success",
	})
	return
}