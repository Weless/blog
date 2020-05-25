package dto

import (
	"blog/dao"
	"blog/model"
)

func GetAllCategories()[]model.Category {
	return dao.GetAllCategories()
}

func GetAllLabels() []model.Label {
	return dao.GetAllLabels()
}

func GetAllArchives() []model.Archive{
	return dao.GetAllArchives()
}


