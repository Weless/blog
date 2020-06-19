package model

import "time"

type Session struct {
	Id        int	`gorm:"primary_key;AUTO_INCREMENT"`
	Uuid      string
	UserId    int
	CreatedTime time.Time
}