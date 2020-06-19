package dao

import (
	"testing"
)

func init() {
	InitConfig()
	InitDB()
}

func TestGetAllCategories(t *testing.T) {
	categories := GetAllCategories()
	t.Logf("categoryList: %+v\n", categories)
}

func TestGetArticlesByCategory(t *testing.T) {
	articles := GetArticlesByCategory(1)
	t.Logf("articles: %+v\n", articles)
}

func TestInsertCategory(t *testing.T) {
	err := InsertCategory("hahah")
	if err != nil {
		t.Error("TestInsertCategory failed")
		return
	}
	t.Logf("insert newCategory success!")
}

func TestGetCategoryList(t *testing.T) {
	categoryList := GetCategoryList()
	t.Logf("categoryList: %+v\n", categoryList)
}

func TestDeleteCategory(t *testing.T) {
	categoryId := 9
	err := DeleteCategory(categoryId)
	if err != nil {
		t.Error("Test DeleteCategory failed")
		return
	}
	t.Logf("Test DeleCategory success!")
}
