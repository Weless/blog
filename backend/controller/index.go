package controller

import (
	"blog/dto"
	"blog/response"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func IndexHandler(c *gin.Context) {
	response.Success(c, "success", nil)
}

func ArticleHandler(c *gin.Context) {
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	if pageNum <= 0 {
		log.Printf("invalid parameter, page_num:%d", pageNum)
		response.Fail(c, "no data", nil)
		return
	}
	articles := dto.GetArticleList(pageNum)

	if len(articles) == 0 {
		response.Fail(c, "no data", nil)
		return
	}

	l := dto.GetTotalLength()

	response.Success(c, "success", gin.H{"articleList": articles,
		"total": l})
}

func ArticleListHandler(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	if page <= 0 || pageSize <= 0 {
		response.Fail(c, "no data", nil)
		return
	}
	articleList := dto.GetArticlesByPageAndPageSize(page, pageSize)
	if len(articleList) == 0 {
		response.Fail(c, "no data", nil)
		return
	}

	l := dto.GetTotalLength()
	response.Success(c, "success", gin.H{"articleList": articleList,
		"total": l})
}

func GetTotalLengthHandler(c *gin.Context) {
	l := dto.GetTotalLength()
	response.Success(c, "success", gin.H{"total": l})
}

func LatestArticleHandler(c *gin.Context) {
	articles := dto.GetLatestArticles(5)
	response.Success(c, "success", gin.H{"articleList": articles})
}

func CategoriesHandler(c *gin.Context) {
	categories := dto.GetAllCategories()
	response.Success(c, "success", gin.H{"categoryList": categories})
}

func ArchiveHandler(c *gin.Context) {
	archives := dto.GetAllArchives()
	response.Success(c, "success", gin.H{"archiveList": archives})

}

func ArchiveArticleHandler(c *gin.Context) {
	year, _ := strconv.Atoi(c.Param("year"))
	month, _ := strconv.Atoi(c.Param("month"))
	if year <= 0 || month <= 0 || month > 12 {
		response.Fail(c, "no data", nil)
		return
	}

	articles := dto.GetArticlesByTime(year, month)
	if len(articles) == 0 {
		response.Fail(c, "no data", nil)
		return
	}
	response.Success(c, "success", gin.H{"articleList": articles})
}

func CategoryArticleHandler(c *gin.Context) {
	categoryId, _ := strconv.Atoi(c.Param("categoryId"))
	if categoryId <= 0 {
		response.Fail(c, "no data", nil)
		return
	}
	articles := dto.GerArticlesByCategory(categoryId)
	if len(articles) == 0 {
		response.Fail(c, "no data", nil)
		return
	}
	response.Success(c, "success", gin.H{"articleList": articles})

}
