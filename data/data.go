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
  Messages []Message
  Groups []Group `gorm:"many2many:members;"`
}

type Message struct {
  gorm.Model
  Text string `json:"name"`
  Image string `json:"name"`
  UserId uint
  User User
}

type Members struct {

}

type Group struct {
  gorm.Model
  Name string
  Image string
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
  session := sessions.Default(c)
  user_id := session.Get("userID")
  message := Message{Text: text, UserId: user_id.(uint),}
  db.Create(&message)
}

func Find_Another_User_Messages(c *gin.Context) ([]Message, User ) {
  db, err := gorm.Open("mysql", "root@/messageapp?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()
  var messages []Message
  var user User
  var current_user User
  session := sessions.Default(c)
  user_id := session.Get("userID")
  db.Model(&user).Related(&messages)
  db.Preload("User").Find(&messages)
  db.Where("id = ?", user_id).Find(&current_user)
  return messages, current_user
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
  var user User
  group := Group {
    Name: name,
  }
  db.Create(&group)
  for i := range user_ids {
  user_id := user_ids[i]
  int_user_id, err := strconv.ParseUint(user_id, 10, 0)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(int_user_id)
  db.Find(&user)
  //db.Where("ID = ?", int_user_id).First(&user)
  db.Preload("Groups").First(&user)
  db.Model(&user).Association("Groups").Append(&user, &group)
  }
}
