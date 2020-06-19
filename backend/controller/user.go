package controller

import (
	"blog/common"
	"blog/dto"
	"blog/model"
	"blog/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func RegisterHandler(c *gin.Context) {
	// 获取参数
	var requestUser = model.User{}
	// json.NewDecoder(c.Request.Body).Decode(&requestUser)
	c.Bind(&requestUser)
	name := requestUser.Name
	password := requestUser.Password
	// 数据验证
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "password should not be less than 6 bit")
		return
	}
	if IsUserExit(name) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "user has already been created")
		return
	}
	// 创建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) // 密码加密
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "encrypt password failed")
		return
	}

	err = dto.CreateUser(name, string(hasedPassword))
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "create user failed")
		return
	}

	// 注册成功直接发放jwtToken
	user := dto.GetUserByName(name)
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error: %v\n", err)
		return
	}

	// 返回结果
	response.Success(c, "register success", gin.H{"token": token})
}

func LoginHandler(c *gin.Context) {
	// name := c.PostForm("name")
	// password := c.PostForm("password")
	var requestUser = model.User{}
	c.Bind(&requestUser)
	name := requestUser.Name
	password := requestUser.Password
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "password should not be less than 6 bit")
		return
	}
	// 判断用户是否存在
	user := dto.GetUserByName(name)
	if user.Id == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "user doesn't exit")
		return
	}
	// 判断密码是否正确
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) // 第一个参数为加密后的密码，第二个为密码
	if err != nil {
		response.Fail(c, "wrong password", nil)
		return
	}

	// 发放jwtToken
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error: %v\n", err)
		return
	}
	// 返回结果
	response.Success(c, "login success", gin.H{"token": token})
}

func UserInfoHandler(c *gin.Context) {
	user, _ := c.Get("user") // 从上下文获取用户信息
	response.Success(c, "success", gin.H{"user": dto.ToUserDTO(user.(model.User))})
}

func IsUserExit(name string) bool {
	user := dto.GetUserByName(name)
	if user.Id != 0 {
		return true
	}
	return false
}
