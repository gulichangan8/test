package service

import (
	"test/dao"
	"test/tool"
)

func CheckUsername(username string) bool {
	ok1 := IsShortName(username)
	ok2 := IsNameErr(username)
	ok3 := IsRepeatUsername(username)
	if ok1 && ok2 {
		if ok3 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

// CheckPassword 检查密码格式
func CheckPassword(password string) bool {
	ok1 := IsPasswordErr(password)
	ok2 := IsShortPassword(password)
	if ok1 && ok2 {
		return true
	} else {
		return false
	}
}

// CheckLogin 登录账号检查
func CheckLogin(username string, password string) bool {
	U := dao.TakeUserDate()
	ok := false
	for _, date := range U {
		if date.Username == username && date.Password == password {
			ok = true
		} else {
			continue
		}
	}
	if ok {
		return true
	} else {
		return false
	}
}

// PasswordProtect 密保问题是否回答正确
func PasswordProtect(username string, trueName string, likeFood string, age int) bool {
	Q := dao.TakeQuestionDate()
	ok := false
	for _, date := range Q {
		if date.UserName == username && date.TrueName == trueName && date.LikeFood == likeFood && date.Age == age {
			ok = true
		} else {
			continue
		}
	}
	if ok {
		return true
	} else {
		return false
	}
}

func BringUserDate(username string, P string) {
	U := tool.CreateUser(username, P)
	dao.BringUserDate(U)
}

func BringLoginDate(username string) {
	l := tool.CreateLogin(username, true)
	dao.BringLoginDate(l)
}

func BringQuestionDate(username string, trueName string, likeFood string, age int) {
	q := tool.CreateQues(username, trueName, likeFood, age)
	dao.BringQuestionDate(q)
}

func ChangePasswordDate(username string, password string) {
	p := Hmac("a", password)
	dao.ChangePasswordDate(username, p)
}
