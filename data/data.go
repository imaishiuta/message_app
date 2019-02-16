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
  FriendCode string
  StatusMessage string
  Messages []Message
  Groups []Group `gorm:"many2many:members;"`
  Friends []*User `gorm:"many2many:friendships;association_jointable_foreignkey:friend_id"`
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
  fmt.Println(group_id)
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

func Get_Group_Data(c *gin.Context) ([]User, []Message, uint) {
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
  group_uint := uint(int_group_id)
  db.Find(&group, int_group_id)
  db.Model(&group).Association("Users").Find(&users)
  db.Preload("User").Order("created_at desc").Where("group_id = ?", &group.ID).Find(&messages)
  fmt.Println(messages)
  return users, messages, group_uint
}

func Get_Chat_Data(c *gin.Context, current_user User, user_id string) ([]User, []Message) {
  db, err := gorm.Open("mysql", "root@/messageapp?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()
  int_user_id, err := strconv.ParseUint(user_id, 10, 0)
    if err != nil {
      fmt.Println(err)
    }
  var users []User
  var messages []Message
  uint_user_id := uint(int_user_id)
  my_user_id := current_user.ID
  db.Where("id in (?)", []uint{my_user_id, uint_user_id}).Find(&users)
  db.Preload("User").Where("user_id in (?)", []uint{my_user_id, uint_user_id}).Find(&messages)
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

func Current_User(c *gin.Context) User {
  db, err := gorm.Open("mysql", "root@/messageapp?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()
  var user User
  session := sessions.Default(c)
  user_id := session.Get("userID")
  db.First(&user, user_id)
  return user
}

func Update_User(user User, new_name string, new_status_message string, friend_id string, email string ) {
  db, err := gorm.Open("mysql", "root@/messageapp?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()
  user.Name = new_name
  user.StatusMessage = new_status_message
  user.FriendCode = friend_id
  user.Email = email
  db.Save(&user)
}

func Search_User(keyword string) User {
  db, err := gorm.Open("mysql", "root@/messageapp?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()
  var user User
  db.Where("friend_code = ?", keyword).Find(&user)
  return user
}

func Add_User_Friend(c *gin.Context, user_id string) {
  db, err := gorm.Open("mysql", "root@/messageapp?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()
  var user User
  current_user := Current_User(c)
  int_user_id, err := strconv.ParseUint(user_id, 10, 0)
    if err != nil {
      fmt.Println(err)
    }
  db.First(&user, int_user_id)
  db.Model(&current_user).Association("Friends").Append(&user)
}

func Create_User_Chat(c *gin.Context, user_id string) {
  db, err := gorm.Open("mysql", "root@/messageapp?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()

  int_user_id, err := strconv.ParseUint(user_id, 10, 0)
    if err != nil {
      fmt.Println(err)
    }

  var users []User
  current_user := Current_User(c)
  group := Group{
    Name: "Personal",
  }

  db.Create(&group)
  uint_user_id := uint(int_user_id)
  my_user_id := current_user.ID
  db.Where("id in (?)", []uint{my_user_id, uint_user_id}).Find(&users)
  fmt.Println(users)
  fmt.Println(group)
  db.Model(&group).Association("Users").Append(&users, &group)
}
