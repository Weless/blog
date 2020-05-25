package model

type Label struct {
	L_id int `gorm:"primary_key;AUTO_INCREMENT"`
	Name string
}