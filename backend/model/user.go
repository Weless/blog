package model

import "time"

type User struct {
	Id        uint	`gorm:"primary_key;AUTO_INCREMENT"`
	Name      string `gorm:"unique"`
	Password  string
	CreatedTime time.Time
}
