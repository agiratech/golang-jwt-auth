package models

import (
  "github.com/jinzhu/gorm"
  "github.com/dgrijalva/jwt-go"
)
// user representation structure
type User struct {
  gorm.Model
  Email string     `json:"email"`
  Username string  `json:"username"`
  Password string  `json:"password"`
  Token string     `json:"token"`
}

type Token struct {
  UserId uint
  jwt.StandardClaims
}

type Post struct {
  gorm.Model
  Title string       `json:"title"`
  Url string         `json:"url"`
  Description string `json:"description"`
  UserId int         `json:"user_id"`
  Points int         `json:"points"`
  Username string    `json:"username"`
}

type Vote struct {
  gorm.Model
  UserId int `json:"user_id"`
  PostId int `json:"post_id"`
}

type Comment struct {
  gorm.Model
  Text string   `json:"text"`
  UserId int    `json:"user_id"`
  PostId int    `json:"post_id"`
  ParentId int  `json: "parent_id"`
}