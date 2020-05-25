package model

import "time"

type Article struct {
	ArticleId    int `gorm:"primary_key;AUTO_INCREMENT"`
	Title        string
	CreateTime   time.Time
	CommentCount int
	ViewCount    int
	Content      string
	Summary      string
	CategoryId   int
	LabelId      int
}

type ArticleInfo struct {
	Article
	Name string
}

type LatestArticle struct {
	ArticleId    int `gorm:"primary_key;AUTO_INCREMENT"`
	Title        string
}





