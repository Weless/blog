package dto

import (
	"blog/dao"
	"blog/model"
)

func TestGetCategoryList() []model.Category {
	return dao.GetCategoryList()
}
