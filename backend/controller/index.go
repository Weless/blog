package controller

import (
	"blog/dto"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func IndexHandler(c *gin.Context)  {
	
	c.JSON(200,gin.H{
		"code":200,
		"msg":"success",
	})
}

func ArticleHandler(c *gin.Context)  {
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	if pageNum <= 0 {
		log.Printf("invalid parameter, page_num:%d", pageNum)
		c.JSON(404,gin.H{
			"code":404,
			"msg":"no data",
		})
		return
	}
	articles := dto.GetArticleList(pageNum)

	if len(articles) == 0{
		c.JSON(404,gin.H{
			"code":404,
			"msg":"no data",
		})
		return
	}

	l := dto.GetTotalLength()

	c.JSON(200,gin.H{
		"code":200,
		"msg":"success",
		"articleList":articles,
		"total":l,
	})
}

func GetTotalLengthHandler(c *gin.Context)  {
	l := dto.GetTotalLength()
	c.JSON(200,gin.H{
		"code":200,
		"msg":"success",
		"total":l,
	})
}

func LatestArticleHandler(c *gin.Context)  {
	articles := dto.GetLatestArticles(5)
	c.JSON(200,gin.H{
		"code":200,
		"msg":"success",
		"articleList":articles,
	})
}

func CategoriesHandler(c *gin.Context){
	categories := dto.GetAllCategories()
	c.JSON(200,gin.H{
		"code":200,
		"msg":"success",
		"categoryList":categories,
	})
}

func LabelHandler(c *gin.Context)  {
	labels := dto.GetAllLabels()
	c.JSON(200,gin.H{
		"code":200,
		"msg":"success",
		"labelList":labels,
	})
}

func ArchiveHandler(c *gin.Context)  {
	archives := dto.GetAllArchives()
	c.JSON(200,gin.H{
		"code":200,
		"msg":"success",
		"archiveList":archives,
	})
}

func ArchiveArticleHandler(c *gin.Context)  {
	year, _:= strconv.Atoi(c.Param("year"))
	month, _:= strconv.Atoi(c.Param("month"))
	if year <= 0 || month <= 0 || month > 12{
		c.JSON(404,gin.H{
			"code":404,
			"msg":"no data",
		})
		return
	}

	articles := dto.GetArticlesByTime(year,month)
	if len(articles) == 0{
		c.JSON(404,gin.H{
			"code":404,
			"msg":"no data",
		})
		return
	}
	c.JSON(200,gin.H{
		"code":200,
		"msg":"success",
		"articleList":articles,
	})
}

func CategoryArticleHandler(c *gin.Context)  {
	categoryId, _:= strconv.Atoi(c.Param("categoryId"))
	if categoryId <= 0{
		c.JSON(404,gin.H{
			"code":404,
			"msg":"no data",
		})
		return
	}
	articles := dto.GerArticlesByCategory(categoryId)
	if len(articles) == 0{
		c.JSON(404,gin.H{
			"code":404,
			"msg":"no data",
		})
		return
	}
	c.JSON(200,gin.H{
		"code":200,
		"msg":"success",
		"articleList":articles,
	})
}

func LabelArticleHandler(c *gin.Context)  {
	labelId, _:= strconv.Atoi(c.Param("labelId"))
	if labelId <= 0{
		c.JSON(404,gin.H{
			"code":404,
			"msg":"no data",
		})
		return
	}
	articles := dto.GetArticlesByLabel(labelId)
	if len(articles) == 0{
		c.JSON(404,gin.H{
			"code":404,
			"msg":"no data",
		})
		return
	}
	c.JSON(200,gin.H{
		"code":200,
		"msg":"success",
		"articleList":articles,
	})
}