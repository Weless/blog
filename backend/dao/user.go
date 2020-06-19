package dao

import (
	"blog/model"
	"blog/utils"
	"time"
)

func CreateUser(name string, password string) error {
	var user model.User
	user.Name = name
	user.Password = password
	user.CreatedTime = time.Now()
	err := DB.Create(&user).Error
	return err
}

func UpdateUser(name, password string) error {
	var user model.User
	DB.Where("name = ?", name).Take(&user)
	newPassWord := utils.Encrypt(password)
	err := DB.Model(&user).Update("password", newPassWord).Error
	return err
}

func DeleteUserById(id uint) error {
	var user model.User
	user.Id = id
	err := DB.Delete(&user).Error
	return err
}

func GetUserById(id uint) model.User {
	var user model.User
	DB.First(&user, id)
	return user
}

func GetUserByName(name string) model.User {
	var user model.User
	DB.Where("name = ?", name).First(&user)
	return user
}

// func GetUserByUUID(uuid int) model.User {
// 	var user model.User
// 	DB.Where("uuid = ?", uuid).First(&user)
// 	return user
// }
