package api

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"test/respond"
	"test/service"
)

// Register 注册接口
func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	ok1 := service.CheckUsername(username)
	ok2 := service.CheckPassword(password)
	P := service.Hmac("a", password)
	if ok1 && ok2 {
		respond.RegisterTrue(c)
		service.BringUserDate(username, P)
	} else {
		respond.RegisterErr(c)
	}
}

// Login 登录接口
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	P := service.Hmac("a", password)
	ok := service.CheckLogin(username, P)
	if ok {
		respond.LoginTrue(c)
		service.BringLoginDate(username)
	} else {
		respond.LoginErr(c)
	}
}

// Question 设置密保接口
func Question(c *gin.Context) {
	username := c.PostForm("username")
	trueName := c.PostForm("true_name")
	likeFood := c.PostForm("like_food")
	age, _ := strconv.Atoi(c.PostForm("age"))
	service.BringQuestionDate(username, trueName, likeFood, age)
	respond.QuestionTrue(c)
}

// AnswerQuestion 回答密保接口
func AnswerQuestion(c *gin.Context) {
	username := c.PostForm("username")
	trueName := c.PostForm("true_name")
	likeFood := c.PostForm("like_food")
	age, _ := strconv.Atoi(c.PostForm("age"))
	ok := service.PasswordProtect(username, trueName, likeFood, age)
	if ok {
		respond.AnswerQuestionTrue(c)
	} else {
		respond.AnswerQuestionErr(c)
	}
}

// ChangePassword 修改密码接口
func ChangePassword(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	ok := service.CheckPassword(password)
	if ok {
		service.ChangePasswordDate(username, password)
		respond.ChangePasswordTrue(c)
	} else {
		respond.ChangePasswordErr(c)
	}
}
