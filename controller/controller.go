package controller

import(
  "../session"
  "../data"
  "github.com/gin-gonic/contrib/sessions"
  "github.com/gin-gonic/gin"
  "fmt"
)

func UserSigninRedirect(c *gin.Context) {
  session := sessions.Default(c)
  user_id := session.Get("userID")
  if user_id == nil {
    c.Redirect(301, "/signin")
  }
}

func IndexRouter(c *gin.Context) {
  c.HTML(200, "index.html", nil)
}

func UserRouter(c *gin.Context) {
  UserSigninRedirect(c)
  c.HTML(200, "user.html", nil)
}

func ChatRouter(c *gin.Context) {
  UserSigninRedirect(c)
  users, messages := data.Get_Group_Data(c)
  current_user := data.Get_Current_User(c)
  c.HTML(200, "chat.html", gin.H{
    "Users": users,
    "Messages": messages,
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
  c.Redirect(301, "/chatrooms")
}

func SigninRouter(c *gin.Context) {
  c.HTML(200, "signin.html", nil)
}

func PostSigninRouter(c *gin.Context) {
  email := c.PostForm("email")
  password := c.PostForm("password")
  user := data.Find_User(email, password)
  session.Login(c, user)
  c.Redirect(301, "/chatrooms")
}

func PostMessage(c *gin.Context) {
  text := c.PostForm("text")
  data.Create_Message(c, text)
  c.Redirect(301, "/chatroom")
}

func GourpRouter(c *gin.Context) {
  UserSigninRedirect(c)
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
  UserSigninRedirect(c)
  group := data.Find_Group(c)
  current_user := data.Get_Current_User(c)
  c.HTML(200, "user.html", gin.H{
    "Group": group,
    "Current_user": current_user,
    })
}

func EditUserRouter(c *gin.Context) {
  UserSigninRedirect(c)
  current_user := data.Get_Current_User(c)
  c.HTML(200, "edit_user.html", gin.H{
    "Current_user": current_user,
    })
}

func UpdateUserRouter(c *gin.Context) {
  new_name := c.PostForm("name")
  new_status_message := c.PostForm("status_message")
  friend_id := c.PostForm("friend_id")

  email := c.PostForm("email")
  fmt.Println("loaded")
  user := data.Current_User(c)
  data.Update_User(user, new_name, new_status_message, friend_id, email)
  c.Redirect(301, "/chatrooms")
}

func AddUserRouter(c *gin.Context) {
  UserSigninRedirect(c)
  c.HTML(200, "add_user.html", nil)
}

func SearchUserRouter(c *gin.Context) {
  keyword := c.PostForm("keyword")
  fmt.Println(keyword)
  user := data.Search_User(keyword)
  fmt.Println(user)
  c.JSON(200, gin.H{
    "User": user,
    })
}
