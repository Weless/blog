package dao

import (
	"blog/model"
	"log"
)

func GetAllCategories() []model.CategoryInfo {
	var categories []model.CategoryInfo
	sql := ` select c_id,name,count(name) as cnt from categories 
    		inner join articles 
    		on articles.category_id = categories.c_id 
			group by c_id`
	DB.Raw(sql).Scan(&categories)
	return categories
}

func GetCategoryList() []model.Category {
	var categoryList []model.Category
	DB.Find(&categoryList)
	return categoryList
}

func GetArticlesByCategory(categoryId int) []model.ArticleInfo {
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
	categories.name AS name 
FROM
	articles
	JOIN categories ON categories.c_id = articles.category_id 
WHERE
	category_id = ?
ORDER BY
	article_id DESC`
	DB.Raw(sql, categoryId).Scan(&articles)
	return articles
}

func InsertCategory(categoryName string) error {
	var newCategory model.Category
	newCategory.Name = categoryName
	err := DB.Create(&newCategory).Error
	if err != nil {
		log.Println("insert new category failed, err:", err)
		return err
	}
	return nil
}

func DeleteCategory(categoryId int) error {
	var category model.Category
	category.C_id = categoryId
	err := DB.Delete(&category).Error
	if err != nil {
		log.Println("delete category failed, err:", err)
		return err
	}
	return nil
}
