package dao

import "testing"

func init()  {
	InitConfig()
	InitDB()
}

func TestGetAllArchives(t *testing.T) {
	archives := GetAllArchives()
	t.Logf("archives: %+v\n",archives)
}

func TestGetArticlesByTime(t *testing.T) {
	articles := GetArticlesByTime(2020, 5)
	t.Logf("articles: %+v\n",articles)
}
