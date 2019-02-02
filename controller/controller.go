package controller

import(
  "fmt"
  "../session"
  "../data"
  "github.com/gin-gonic/gin"
)

func IndexRouter(c *gin.Context) {
  c.HTML(200, "index.html", nil)
}

func UserRouter(c *gin.Context) {
  c.HTML(200, "user.html", nil)
}

func ChatRouter(c *gin.Context) {
  messages, current_user := data.Find_Another_User_Messages(c)
  fmt.Println(current_user.ID)
  c.HTML(200, "chat.html", gin.H{
    "messages": messages,
    "current_user": current_user,
    })
}

func SignupRouter(c *gin.Context) {
  c.HTML(200, "signup.html", nil)
}

func PostSignupRouter(c *gin.Context) {
  name := c.PostForm("name")
  email := c.PostForm("email")
  password := c.PostForm("password")
  confirm_password := c.PostForm("confirm_password")
  data.Create_User(name, email, password, confirm_password)
  c.HTML(200, "signin.html", nil)
}

func SigninRouter(c *gin.Context) {
  c.HTML(200, "signin.html", nil)
}

func PostSigninRouter(c *gin.Context) {
  email := c.PostForm("email")
  password := c.PostForm("password")
  user := data.Find_User(email, password)
  session.Login(c, user)
  c.HTML(200, "chat.html", nil)
}

func PostMessage(c *gin.Context) {
  text := c.PostForm("text")
  data.Create_Message(c, text)
  c.HTML(200, "chat.html", nil)
}
