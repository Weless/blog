package model

type CategoryInfo struct {
	C_id int    `gorm:"primary_key;AUTO_INCREMENT"`
	Name string `gorm:"unique"`
	Cnt  int
}

type Category struct {
	C_id int `gorm:"primary_key;AUTO_INCREMENT"`
	Name string
}
