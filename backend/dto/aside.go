package dto

import (
	"blog/dao"
	"blog/model"
)

func GetAllCategories() []model.CategoryInfo {
	return dao.GetAllCategories()
}

func GetAllArchives() []model.Archive {
	return dao.GetAllArchives()
}

func InsertCategory(categoryName string) error {
	return dao.InsertCategory(categoryName)
}
