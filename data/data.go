package data

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/gin-gonic/contrib/sessions"
  "github.com/jinzhu/gorm"
  "strconv"
)

type User struct {
  gorm.Model
  Name string
  Email string
  Password string
  ConfirmPassword string
  FriendId int
  Messages []Message
  Groups []Group `gorm:"many2many:members;"`
}

type Message struct {
  gorm.Model
  Text string `json:"name"`
  Image string `json:"name"`
  UserId uint
  GroupId uint
  Group Group
  User User
}

type Group struct {
  gorm.Model
  Name string
  Image string
  Messages []Message
  Users []User `gorm:"many2many:members;"`
}

func InitDB() {
  db, err := gorm.Open("mysql", "root@/messageapp?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()
  db.AutoMigrate(&User{}, &Message{}, &Group{})
}

func Find_User(email string, password string) User {
  db, err := gorm.Open("mysql", "root@/messageapp?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()
  var user User
  db.Where("email = ? AND password = ?", email, password).Find(&user)
  return user
}

func Create_User(name string, email string, password string, ConfirmPassword string) {
  db, err := gorm.Open("mysql", "root@/messageapp?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()
  user := User{
    Name: name,
    Email: email,
    Password: password,
    ConfirmPassword: ConfirmPassword,
  }
  db.Create(&user)
}

func Create_Message(c *gin.Context, text string) {
  db, err := gorm.Open("mysql", "root@/messageapp?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()
  group_id := c.Param("id")
  int_group_id, err := strconv.ParseUint(group_id, 10, 0)
    if err != nil {
      fmt.Println(err)
    }
  group_uint := uint(int_group_id)
  session := sessions.Default(c)
  user_id := session.Get("userID")
  message := Message{
    Text: text,
    UserId: user_id.(uint),
    GroupId: group_uint,
  }
  db.Create(&message)
}

func Get_Group_Data(c *gin.Context) ([]User, []Message) {
  db, err := gorm.Open("mysql", "root@/messageapp?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()
  var users []User
  var messages []Message
  var group Group
  group_id := c.Param("id")
  int_group_id, err := strconv.ParseUint(group_id, 10, 0)
    if err != nil {
      fmt.Println(err)
    }
  db.Find(&group, int_group_id)
  db.Model(&group).Association("Users").Find(&users)
  db.LogMode(true)
  db.Preload("User").Order("created_at desc").Where("group_id = ?", &group.ID).Find(&messages)
  fmt.Println(messages)
  return users, messages
}

func Get_Current_User(c *gin.Context) User {
  db, err := gorm.Open("mysql", "root@/messageapp?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()
  var current_user User
  var messages []Message
  session := sessions.Default(c)
  user_id := session.Get("userID")
  db.Model(&current_user).Related(&messages)
  db.Preload("User").Find(&messages)
  db.Where("id = ?", user_id).Find(&current_user)
  return current_user
}

func Get_All_User() []User{
  db, err := gorm.Open("mysql", "root@/messageapp?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()
  var user []User
  db.Find(&user)
  return user
}

func Create_Group(name string, user_ids []string) {
  db, err := gorm.Open("mysql", "root@/messageapp?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()
  group := Group {
    Name: name,
  }
  db.Create(&group)
  for i := range user_ids {
    var user User
    user_id := user_ids[i]
    int_user_id, err := strconv.ParseUint(user_id, 10, 0)
    if err != nil {
      fmt.Println(err)
    }
    db.Preload("Groups").First(&user, int_user_id)
    db.Model(&user).Association("Groups").Append(&user, &group)
  }
}

func Find_Group(c *gin.Context) []Group {
  db, err := gorm.Open("mysql", "root@/messageapp?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()
  var group []Group
  var user User
  session := sessions.Default(c)
  user_id := session.Get("userID")
  db.First(&user, user_id)
  db.Model(&user).Association("Groups").Find(&group)
  return group
}
