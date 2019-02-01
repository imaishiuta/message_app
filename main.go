package main

import (
  "./controller"
  "github.com/gin-gonic/gin"
  "github.com/gin-gonic/contrib/sessions"
  _"github.com/go-sql-driver/mysql"
)

func main() {
  server := gin.Default()
  store := sessions.NewCookieStore([]byte("secret"))
  server.Use(sessions.Sessions("SessionName", store))

  server.LoadHTMLGlob("templates/*")
  server.Static("/assets", "./assets")
  server.GET("/", controller.UserRouter)
  server.GET("/chatroom", controller.ChatRouter)
  server.Run()
}
