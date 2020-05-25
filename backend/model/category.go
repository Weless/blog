package model


type Category struct {
	C_id int `gorm:"primary_key;AUTO_INCREMENT"`
	Name string
	Cnt int
}

