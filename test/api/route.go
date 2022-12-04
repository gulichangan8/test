package api

import "github.com/gin-gonic/gin"

func InitEngine() {
	// 用户接口
	r := gin.Default()
	user := r.Group("/user")
	{
		user.POST("/register", Register)
		user.POST("/login", Login)
		user.POST("/question", Question)
		user.POST("/answer", AnswerQuestion)
		user.PUT("/password", ChangePassword)
	}
}
