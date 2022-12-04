package dao

import (
	"database/sql"
	"test/model"
)

// TakeUserDate 将user表中的数据取出
func TakeUserDate() model.Users {
	rows, _ := dB.Query("select * from ?", "user")
	var u model.User
	var U model.Users
	for rows.Next() {
		err := rows.Scan(&u.Username, &u.Password)
		if err != nil {
			panic(err)
		}
		U = append(U, u)
	}
	return U
}

// TakeQuestionDate 将question表中的数据取出
func TakeQuestionDate() model.Queses {
	rows, _ := dB.Query("select * from ?", "user")
	var q model.Ques
	var Q model.Queses
	for rows.Next() {
		err := rows.Scan(&q.UserName, &q.TrueName, &q.LikeFood, &q.Age)
		if err != nil {
			panic(err)
		}
		Q = append(Q, q)
	}
	return Q
}

// BringUserDate 将数据存入user表
func BringUserDate(u model.User) {
	_, err := dB.Exec("insert into user (username,password) value (?,?)",
		u.Username, u.Password)
	if err != nil {
		return
	}
}

// BringQuestionDate 将数据存入question表
func BringQuestionDate(q model.Ques) {
	_, err := dB.Exec("insert into question (username,true_name,like_food,age) value (?,?,?,?)",
		q.UserName, q.TrueName, q.LikeFood, q.Age)
	if err != nil {
		return
	}
}

// BringLoginDate 将数据存入login表
func BringLoginDate(l model.Login) {
	_, err := dB.Exec("insert into login (username,login) value (?,?)",
		l.Username, l.Login)
	if err != nil {
		return
	}
}

// ChangePasswordDate 修改密码
func ChangePasswordDate(username string, password string) {
	var dns = "root:040818@tcp(127.0.0.1:3306)/message_board?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := sql.Open("mysql", dns)
	_, err := db.Exec("update user set password=? where username=?", password, username)
	if err != nil {
		return
	}
}
