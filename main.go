package main

import (
  "./data"
  "./controller"
  "github.com/gin-gonic/gin"
  "github.com/gin-gonic/contrib/sessions"
  _"github.com/go-sql-driver/mysql"
)

func main() {
  data.InitDB()
  server := gin.Default()
  store := sessions.NewCookieStore([]byte("secret"))
  server.Use(sessions.Sessions("SessionName", store))

  server.LoadHTMLGlob("templates/*")
  server.Static("/assets", "./assets")
  server.GET("/", controller.UserRouter)
  server.GET("/signup", controller.SignupRouter)
  server.POST("/user/signup", controller.PostSignupRouter)
  server.GET("/signin", controller.SigninRouter)
  server.POST("/user/signin", controller.PostSigninRouter)
  server.GET("/chatrooms", controller.ChatsRouter)
  server.GET("/chatroom/:id", controller.ChatRouter)
  server.POST("/message", controller.PostMessage)
  server.GET("/group/edit", controller.GourpRouter)
  server.POST("/group/create", controller.CreateGroupRouter)
  server.Run()
}
