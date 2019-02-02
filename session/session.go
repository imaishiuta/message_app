package session

import (
  "github.com/gin-gonic/gin"
  "github.com/gin-gonic/contrib/sessions"
  "../data"
)

type SessionInfo struct {
  UserID         interface{}
  Email       interface{}
  IsSessionAlive bool
}

func Login(c *gin.Context, user data.User) {
  session := sessions.Default(c)
  session.Set("alive", true)
  session.Set("userID", user.ID)
  session.Set("uemail", user.Email)
  session.Save()
}

func Logout(c *gin.Context) {
  session := sessions.Default(c)
  session.Clear()
  session.Save()
}

func GetSessionInfo(c *gin.Context) SessionInfo {
  var info SessionInfo
  session := sessions.Default(c)
  user_id := session.Get("userID")
  email := session.Get("uemail")
  alive := session.Get("alive")
  info = SessionInfo{
    UserID:         user_id.(uint),
    Email:          email.(string),
    IsSessionAlive: alive.(bool),
    }
  return info
}
