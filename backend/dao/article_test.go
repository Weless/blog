package dao

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"testing"
)

func init() {
	InitConfig()
	InitDB()
}

func InitConfig() {
	workDir, err := os.Getwd()
	workDir = workDir[:len(workDir)-4] // 去掉末尾的 /dao
	if err != nil {
		log.Println("failed to get work dir, err:", err)
		return
	}
	viper.SetConfigName("application")       // 设置读取文件的名字
	viper.SetConfigType("yml")               // 设置读取文件的类型
	viper.AddConfigPath(workDir + "/config") // 设置文件路径
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func TestGetArticleList(t *testing.T) {
	articles := GetArticleList(2)
	t.Logf("length: %v\n", len(articles))
	for _, article := range articles {
		t.Logf("article id:%d\n", article.ArticleId)
	}

}

func TestGetTotalLength(t *testing.T) {
	l := GetTotalLength()
	t.Logf("length: %d\n", l)
}

func TestGetLatestArticles(t *testing.T) {
	articles := GetLatestArticles(5)
	t.Logf("articles: %+v\n", articles)
}

func TestGetArticleById(t *testing.T) {
	articles := GetArticleById(1)
	t.Logf("articles: %+v\n", articles)
}

func TestGetLatestArticle(t *testing.T) {
	article := GetLatestArticle()
	t.Logf("articles: %+v\n", article)
}

func TestInsertArticle(t *testing.T) {
	var title = "第一次插入"
	var content = "试试看"
	var summary = "试试看"
	var categoryId = 1
	InsertArticle(title, content, summary, categoryId)
	article := GetLatestArticle()
	t.Logf("articles: %+v\n", article)
}

func TestDeleteArticle(t *testing.T) {
	var id = 18
	DeleteArticle(id)
	article := GetLatestArticle()
	t.Logf("articles: %+v\n", article)
}

func TestUpdateArticle(t *testing.T) {
	title := ""
	content := ""
	summary := "哈哈哈"
	categoryId := 6
	id := 15
	UpdateArticle(id, title, content, summary, categoryId)
	article := GetArticleById(id)
	t.Logf("articles: %+v\n", article)
}

func TestIsIdExisted(t *testing.T) {
	exist := IsIdExisted(15)
	if exist {
		t.Logf("id exists")
	} else {
		t.Logf("id does not exists")
	}
}

func TestGetArticlesByPageAndPageSize(t *testing.T) {
	articles := GetArticlesByPageAndPageSize(2, 3)
	t.Logf("articles: %+v\n", articles)
}

func TestGetArticleByIdWithCategoryName(t *testing.T) {
	article := GetArticleByIdWithCategoryName(1)
	t.Logf("articles: %+v\n", article)
}
