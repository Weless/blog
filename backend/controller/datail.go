package controller

import (
	"blog/dao"
	"blog/dto"
	"blog/response"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func ArticleDetailController(c *gin.Context) {
	articleKey := "article:" + c.Param("articleId")
	exist, _ := dao.RDB.Exists(articleKey).Result()

	if exist == 0 {
		articleId, _ := strconv.Atoi(c.Param("articleId"))
		if articleId <= 0 {
			response.Fail(c, "no data", nil)
			return
		}
		article, err := dto.GetArticleById(articleId)
		if err != nil {
			response.Fail(c, "no data", nil)
			return
		}

		data := make(map[string]interface{})
		// 时间转化为字符串，要不会出现乱码
		createTime := article.CreateTime.Format("2006-01-02 15:04:05")
		updateTime := article.UpdatedAt.Format("2006-01-02 15:04:05")
		data["ArticleId"] = article.ArticleId
		data["Title"] = article.Title
		data["CreateTime"] = createTime
		data["UpdatedAt"] = updateTime
		data["CommentCount"] = article.CommentCount
		data["ViewCount"] = article.ViewCount
		data["Content"] = article.Content
		data["Summary"] = article.Summary

		_, err = dao.RDB.HMSet(articleKey, data).Result()
		if err != nil {
			response.Fail(c, "set cache failed", nil)
			return
		}
		_, err = dao.RDB.Expire(articleKey, 4*3600*time.Second).Result() // 缓存4小时
		if err != nil {
			response.Fail(c, "set expire failed", nil)
			return
		}
		response.Success(c, "success", gin.H{"article": data})
	} else {
		article, err := dao.RDB.HGetAll(articleKey).Result()
		if err != nil {
			response.Fail(c, "get cache failed", nil)
			return
		}
		response.Success(c, "success", gin.H{"article": article, "status": "from cache"})
	}
}

func ArticleDetailController2(c *gin.Context) {
	articleId, _ := strconv.Atoi(c.Param("articleId"))
	if articleId <= 0 {
		response.Fail(c, "no data", nil)
		return
	}
	article, err := dto.GetArticleById(articleId)
	if err != nil {
		response.Fail(c, "no data", nil)
		return
	}
	response.Success(c, "success", gin.H{"article": article})
}
