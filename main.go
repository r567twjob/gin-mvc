package main

import (
	"github.com/gin-gonic/gin"
	"github.com/your-package/controllers"
)

func main() {
	r := gin.Default()

	userController := &controllers.UserController{}

	// 設置路由
	r.GET("/", homeHandler)
	r.GET("/users", userController.GetUsers)
	r.POST("/users", userController.CreateUser)

	// 啟動服務器
	r.Run(":8080")
}

// 其他路由處理函數...
