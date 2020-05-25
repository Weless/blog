package dao

import "blog/model"

func GetAllLabels() []model.Label {
	var labels []model.Label
	DB.Find(&labels)
	return labels
}

func GetArticlesByLabel(labelId int) []model.ArticleInfo {
	var articles []model.ArticleInfo
	sql := `SELECT
	article_id,
	title,
	create_time,
	comment_count,
	view_count,
	summary,
	category_id,
	label_id,
	categories.name AS name 
FROM
	articles
	JOIN categories ON categories.c_id = articles.category_id 
WHERE
	label_id = ?
ORDER BY
	article_id DESC`
	DB.Raw(sql,labelId).Scan(&articles)
	return articles
}