package dao

import (
	"blog/model"
	"time"
)

func GetArticleList(pageNum int) []model.ArticleInfo {
	var articles []model.ArticleInfo
	sql := `SELECT 
	article_id,
    title,
    create_time,
    comment_count,
    view_count,
    summary,
    category_id,
    categories.name
    from articles 
    join categories 
    on categories.c_id = articles.category_id
	order by article_id desc
	limit 5 offset ? 
	`
	DB.Raw(sql, 5*(pageNum-1)).Scan(&articles)
	return articles
}

// 管理端用
func GetArticlesByPageAndPageSize(page, pageSize int) []model.ArticleInfo {
	var articles []model.ArticleInfo
	sql := `SELECT 
	article_id,
    title,
    create_time,
	updated_at,
    comment_count,
    view_count,
    summary,
    category_id,
    categories.name
    from articles 
    join categories 
    on categories.c_id = articles.category_id
	order by article_id desc
	limit ? offset ? 
	`
	DB.Raw(sql, pageSize, pageSize*(page-1)).Scan(&articles)
	return articles
}

func GetTotalLength() int {
	var count int
	DB.Table("articles").Count(&count)
	return count
}

func GetArticleById(id int) []model.Article {
	var article []model.Article
	DB.First(&article, id)
	return article
}

// 管理端用
func GetArticleByIdWithCategoryName(articleId int) model.ArticleInfo {
	var article model.ArticleInfo
	sql := `SELECT
	article_id,
	title,
	create_time,
	content,
	updated_at,
	comment_count,
	view_count,
	summary,
	category_id,
	categories.name AS name 
FROM
	articles
	JOIN categories ON categories.c_id = articles.category_id 
WHERE
	article_id = ?
	`
	DB.Raw(sql, articleId).Scan(&article)
	return article
}

func GetLatestArticles(num int) []model.Article {
	var articles []model.Article
	DB.Order("article_id desc").Limit(num).Select("article_id,title").Find(&articles)
	return articles
}

func InsertArticle(title, content, summary string, categoryId int) {
	var article model.Article
	article.Title = title
	article.Content = content
	article.Summary = summary
	article.CategoryId = categoryId
	article.CreateTime = time.Now()
	article.UpdatedAt = time.Now()
	DB.Create(&article)
}

func DeleteArticle(id int) {
	var article model.Article
	article.ArticleId = id
	DB.Delete(&article)
}

func UpdateArticle(id int, title, content, summary string, categoryId int) {
	var article model.Article
	DB.Where("article_id = ?", id).Take(&article)

	var updatedMap = make(map[string]interface{}, 5)

	if title != "" {
		updatedMap["title"] = title
	}
	if content != "" {
		updatedMap["content"] = content
	}
	if summary != "" {
		updatedMap["summary"] = summary
	}
	if categoryId != 0 {
		// fmt.Println(1)
		updatedMap["category_id"] = categoryId
	}
	updatedMap["updated_at"] = time.Now()
	DB.Model(&article).Updates(updatedMap)
}

func GetLatestArticle() model.Article {
	var article model.Article
	DB.Last(&article)
	return article
}

func IsIdExisted(id int) bool {
	var article model.Article
	DB.Where("article_id = ?", id).First(&article)
	if article.ArticleId != 0 {
		return true
	} else {
		return false
	}
}
