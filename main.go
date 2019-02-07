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

  server.GET("/", controller.ChatListRouter)
  server.POST("/user/signup", controller.PostSignupRouter)
  server.POST("/user/signin", controller.PostSigninRouter)
  server.GET("/chatrooms", controller.ChatListRouter)
  server.GET("/group/chatroom/:id", controller.GroupChatRouter)
  server.GET("/user/chatroom/:id", controller.UserChatRouter)
  server.POST("/message", controller.PostMessage)
  server.GET("/group/edit", controller.GourpRouter)
  server.POST("/group/create", controller.CreateGroupRouter)
  server.GET("/user/edit/:id", controller.EditUserRouter)
  server.POST("/user/update/:id", controller.UpdateUserRouter)
  server.GET("/add/users", controller.AddUserRouter)
  server.POST("/users/search", controller.SearchUserRouter)
  server.GET("/signinsignup", controller.UserSigninSignoutRouter)
  server.POST("/users/add", controller.AddFriendRouter)
  server.Run()
}
