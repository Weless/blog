package dto

import (
	"blog/dao"
	"blog/model"
)

type UserDTO struct {
	Name string `json:"name"`
}

func ToUserDTO(user model.User) UserDTO {
	return UserDTO{
		Name: user.Name,
	}
}

func CreateUser(name string, password string) error {
	err := dao.CreateUser(name, password)
	return err
}

func UpdateUser(name, password string) error {
	err := dao.UpdateUser(name, password)
	return err
}

func DeleteUserById(id uint) error {
	err := dao.DeleteUserById(id)
	return err
}

func GetUserById(id uint) model.User {
	user := dao.GetUserById(id)
	return user
}

func GetUserByName(name string) model.User {
	user := dao.GetUserByName(name)
	return user
}
