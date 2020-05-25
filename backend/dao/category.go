package dao

import "blog/model"

func GetAllCategories() []model.Category {
	var categories []model.Category
	sql := ` select c_id,name,count(name) as cnt from categories 
    		inner join articles 
    		on articles.category_id = categories.c_id 
			group by c_id`
	DB.Raw(sql).Scan(&categories)
	return categories
}

func GetArticlesByCategory(categoryId int) []model.ArticleInfo {
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
	category_id = ?
ORDER BY
	article_id DESC`
	DB.Raw(sql,categoryId).Scan(&articles)
	return articles
}

