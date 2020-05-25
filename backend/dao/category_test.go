package dao

import (
	"testing"
)

func init()  {
	InitConfig()
	InitDB()
}


func TestGetAllCategories(t *testing.T) {
	categories := GetAllCategories()
	t.Logf("categoryList: %+v\n",categories)
}

func TestGetArticlesByCategory(t *testing.T) {
	articles := GetArticlesByCategory(1)
	t.Logf("articles: %+v\n",articles)
}