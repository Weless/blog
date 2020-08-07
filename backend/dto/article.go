package dto

import (
	"blog/dao"
	"blog/model"
)

func GetArticleList(pageSize, pageNum int) []model.ArticleInfo {
	articleList := dao.GetArticleList(pageSize, pageNum)
	return articleList
}

func GetArticlesByPageAndPageSize(page, pageSize int) []model.ArticleInfo {
	return dao.GetArticlesByPageAndPageSize(page, pageSize)
}
func GetTotalLength() int {
	return dao.GetTotalLength()
}

func GetLatestArticles(num int) []model.Article {
	return dao.GetLatestArticles(num)
}

func GetArticlesByTime(year, month int) []model.ArticleInfo {
	return dao.GetArticlesByTime(year, month)
}

func GerArticlesByCategory(categoryId int) []model.ArticleInfo {
	return dao.GetArticlesByCategory(categoryId)
}

func GetArticleById(id int) (model.Article, error) {
	return dao.GetArticleById(id)
}

func GetArticleByIdWithCategoryName(articleId int) model.ArticleInfo {
	return dao.GetArticleByIdWithCategoryName(articleId)
}

func InsertArticle(title, content, summary string, categoryId int) {
	dao.InsertArticle(title, content, summary, categoryId)
}

func UpdateArticle(id int, title, content, summary string, categoryId int) {
	dao.UpdateArticle(id, title, content, summary, categoryId)
}

func DeleteArticle(id int) {
	dao.DeleteArticle(id)
}

func SearchId(id int) bool {
	return dao.IsIdExisted(id)

}
