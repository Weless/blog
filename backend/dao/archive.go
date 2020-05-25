package dao

import "blog/model"

func GetAllArchives() []model.Archive {
	var archives []model.Archive
	sql := `select DISTINCT DATE_FORMAT(create_time, '%Y年%m月') as pub_time, count(*) as cnt from articles 
			GROUP BY DATE_FORMAT(create_time, '%Y年%m月') order by DATE_FORMAT(create_time, '%Y年%m月') desc`
	DB.Raw(sql).Scan(&archives)
	return archives
}

func GetArticlesByTime(year,month int)[]model.Article {
	var articles []model.Article
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
	YEAR ( create_time ) = ?
	AND MONTH ( create_time ) = ?
ORDER BY
	article_id DESC`
	DB.Raw(sql,year,month).Scan(&articles)
	return articles
}