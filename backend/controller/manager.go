package controller

import (
	"blog/dto"
	"blog/model"
	"blog/response"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func DeleteHandler(c *gin.Context) {
	// articleId, _ := strconv.Atoi(c.PostForm("articleId"))
	var article model.Article
	c.Bind(&article)
	articleId := article.ArticleId
	if articleId <= 0 {
		response.Response(c, 400, 400, nil, "request failed")
		return
	}

	if !dto.SearchId(articleId) {
		response.Response(c, 400, 400, nil, "request failed")
		return
	}
	dto.DeleteArticle(articleId)
	response.Success(c, "success", nil)
}

func InsertHandler(c *gin.Context) {
	var article model.Article
	c.Bind(&article)
	title := article.Title
	content := article.Content
	summary := article.Summary
	categoryId := article.CategoryId
	if categoryId <= 0 {
		response.Response(c, 400, 400, nil, "request failed")
		return
	}
	dto.InsertArticle(title, content, summary, categoryId)

	response.Success(c, "success", nil)
}

func UpdateHandler(c *gin.Context) {
	var article model.Article
	c.Bind(&article)
	articleId := article.ArticleId
	title := article.Title
	content := article.Content
	summary := article.Summary
	categoryId := article.CategoryId

	if !dto.SearchId(articleId) {
		response.Response(c, 400, 400, nil, "request failed")
		return
	}

	dto.UpdateArticle(articleId, title, content, summary, categoryId)

	response.Success(c, "success", nil)
}

func InsertCategoryHandler(c *gin.Context) {
	var newCategory model.Category
	c.Bind(&newCategory)
	categoryName := newCategory.Name
	if categoryName == "" {
		response.Fail(c, "request failed", nil)
		return
	}
	err := dto.InsertCategory(categoryName)
	if err != nil {
		log.Println("failed to call insertCategory function, err:", err)
		return
	}
	response.Success(c, "success", nil)
}

func GetCategoryListHandler(c *gin.Context) {
	categoryList := dto.TestGetCategoryList()
	if len(categoryList) == 0 {
		response.Fail(c, "no data", nil)
		return
	}
	response.Success(c, "success", gin.H{"categoryList": categoryList})
}

func ArticleDetailForManagementHandler(c *gin.Context) {
	id := c.Query("articleId")
	articleId, err := strconv.Atoi(id)
	if err != nil {
		log.Println("fail to call atoi, err:", err)
		return
	}
	if articleId <= 0 {
		response.Fail(c, "no data", nil)
		return
	}
	article := dto.GetArticleByIdWithCategoryName(articleId)

	response.Success(c, "success", gin.H{"article": article})
}
