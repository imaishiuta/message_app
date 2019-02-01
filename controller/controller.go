package controller

import(
  "github.com/gin-gonic/gin"
)

func IndexRouter(c *gin.Context) {
  c.HTML(200, "index.html", nil)
}

func UserRouter(c *gin.Context) {
  c.HTML(200, "user.html", nil)
}

func ChatRouter(c *gin.Context) {
  c.HTML(200, "chat.html", nil)
}
