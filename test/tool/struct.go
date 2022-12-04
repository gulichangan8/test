package tool

import "test/model"

func CreateUser(username string, password string) model.User {
	var U model.User
	U.Username = username
	U.Password = password
	return U
}

func CreateQues(userName string, trueName string, likeFood string, age int) model.Ques {
	var Q model.Ques
	Q.UserName = userName
	Q.TrueName = trueName
	Q.LikeFood = likeFood
	Q.Age = age
	return Q
}

func CreateLogin(username string, login bool) model.Login {
	var L model.Login
	L.Username = username
	L.Login = login
	return L
}
