package dao

import "testing"

func init()  {
	InitConfig()
	InitDB()
}

func TestGetAllLabels(t *testing.T) {
	labels := GetAllLabels()
	t.Logf("labels: %+v\n",labels)
}

func TestGetArticlesByLabel(t *testing.T) {
	articles := GetArticlesByLabel(1)
	t.Logf("articles: %+v\n",articles)
}