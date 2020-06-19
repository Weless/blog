package dao

import (
	"blog/utils"
	"fmt"
	"testing"
)

func TestCreateUser(t *testing.T) {
	name := "joey"
	password := "a12345"
	err := CreateUser(name, password)
	if err != nil {
		fmt.Println(err)
	}
	user := GetUserByName(name)
	t.Logf("user:%+v\n", user)
}

func TestGetUserByName(t *testing.T) {
	name := "joey"
	user := GetUserByName(name)
	t.Logf("user:%+v\n", user)
}

func TestGetUserById(t *testing.T) {
	var id uint
	id = 1
	user := GetUserById(id)
	t.Logf("user:%+v\n", user)
}

func TestDeleteUserById(t *testing.T) {
	var id uint
	id = 100
	err := DeleteUserById(id)
	if err != nil {
		t.Log("err:", err)
		return
	}
}

func TestUpdateUser(t *testing.T) {
	name := "joey"
	newPassword := "qwe123"
	err := UpdateUser(name, newPassword)
	if err != nil {
		t.Log("err:", err)
		return
	}
	user := GetUserByName(name)
	t.Logf("password:%s\n", utils.Encrypt(newPassword))
	t.Logf("newPassWord:%s\n", user.Password)

}
