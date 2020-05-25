package dto

import (
	"blog/dao"
	"blog/model"
)

func GetArticleList(pageNum int ) []model.ArticleInfo{
	articleList := dao.GetArticleList(pageNum)
	return articleList
}

func GetTotalLength() int {
	return dao.GetTotalLength()
}

func GetLatestArticles(num int) []model.Article {
	return dao.GetLatestArticles(num)
}

func GetArticlesByTime(year,month int) []model.Article{
	return dao.GetArticlesByTime(year,month)
}

func GerArticlesByCategory(categoryId int) []model.ArticleInfo{
	return dao.GetArticlesByCategory(categoryId)
}

func GetArticlesByLabel(labelId int ) []model.ArticleInfo {
	return dao.GetArticlesByLabel(labelId)
}

func GetArticleById(id int) []model.Article {
	return dao.GetArticleById(id)
}

func InsertArticle(title,content,summary string, categoryId, labelId int)  {
	dao.InsertArticle(title,content,summary,categoryId,labelId)
}

func UpdateArticle(id int ,title,content,summary string, categoryId,labelId int)  {
	dao.UpdateArticle(id,title,content,summary , categoryId,labelId)
}

func DeleteArticle(id int)  {
	dao.DeleteArticle(id)
}

func SearchId(id int) bool {
	return dao.IsIdExisted(id)

}