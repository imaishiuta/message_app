package controller

import(
  "../session"
  "../data"
  "github.com/gin-gonic/gin"
  "fmt"
)

func IndexRouter(c *gin.Context) {
  c.HTML(200, "index.html", nil)
}

func UserRouter(c *gin.Context) {
  c.HTML(200, "user.html", nil)
}

func ChatRouter(c *gin.Context) {
  users := data.Get_Group_Data(c)
  current_user := data.Get_Current_User(c)
  fmt.Println(users)
  c.HTML(200, "chat.html", gin.H{
    "Users": users,
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
  user := data.Find_User(email, password)
  session.Login(c, user)
  c.Redirect(301, "/chatroom")
}

func SigninRouter(c *gin.Context) {
  c.HTML(200, "signin.html", nil)
}

func PostSigninRouter(c *gin.Context) {
  email := c.PostForm("email")
  password := c.PostForm("password")
  user := data.Find_User(email, password)
  session.Login(c, user)
  c.Redirect(301, "/chatroom")
}

func PostMessage(c *gin.Context) {
  text := c.PostForm("text")
  data.Create_Message(c, text)
  c.Redirect(301, "/chatroom")
}

func GourpRouter(c *gin.Context) {
  user := data.Get_All_User()
  c.HTML(200, "group.html", gin.H {
    "user": user,
    })
}

func CreateGroupRouter(c *gin.Context) {
  user_id := c.PostFormArray("user_ids[]")
  name := c.PostForm("name")
  data.Create_Group(name, user_id)
  c.Redirect(301, "/chatroom")
}

func ChatListRouter(c *gin.Context) {
  group := data.Find_Group(c)
  c.HTML(200, "user.html", gin.H{
    "Group": group,
    })
}
