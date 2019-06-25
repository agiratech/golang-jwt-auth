package models

import (
  "os"
  "errors"
  "strings"
  "time"

  "github.com/jinzhu/gorm"
  "github.com/dgrijalva/jwt-go"
  "golang.org/x/crypto/bcrypt"
)

// Validate user data
func (user *User) validate() error {
  if !strings.Contains(user.Email, "@") {
    return errors.New("Envalid email")
  }

  if len(user.Username) < 6 {
    return errors.New("Envalid username, length should be more than 6")
  }

  if len(user.Password) < 6 {
    return errors.New("Envalid password, length should be more than 6")
  }

  existingUser := &User{}

  //Email must be unique
  //check for errors and duplicate emails
  err := GetDB().Table("users").Where("email = ?", user.Email).First(existingUser).Error
  if err != nil && err != gorm.ErrRecordNotFound {
    return errors.New("Connection error. Please retry")
  }
  if existingUser.Email != "" {
    return errors.New("Email address already in use by another user.")
  }

  //check for errors and duplicate usernames
  err = GetDB().Table("users").Where("username = ?", user.Username).First(existingUser).Error
  if err != nil && err != gorm.ErrRecordNotFound {
    return errors.New("Connection error. Please retry")
  }
  if existingUser.Username != "" {
    return errors.New("Username address already in use by another user.")
  }

  return nil
}

// Create user
func (user *User) CreateUser() (*User, error) {
  if err := user.validate(); err != nil {
    return user, err
  }
  // Salt and hash the password using the bcrypt algorithm
  hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
  user.Password = string(hashedPassword)
  //Create new JWT token for the newly registered user
  GetDB().Create(user)
  if user.ID <= 0 {
    return user, errors.New("Failed to create account, connection error")
  }

  tk := &Token{UserId: user.ID}
  token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
  tokenString, _ := token.SignedString([]byte(os.Getenv("AUTH_TOKEN")))
  user.Token = tokenString
  GetDB().Model(&user).Update("token", tokenString)
  user.Password = "" //remove password
  return user, nil
}

// User login
func (user *User) LoginUser() (*User, error) {
  existingUser := &User{}

  err := GetDB().Table("users").Where("email = ? OR username= ?", user.Username, user.Username).First(existingUser).Error
  if err != nil {
    if (err == gorm.ErrRecordNotFound) {
      return user, errors.New("Username not found")
    }
    return user, errors.New("Connection error. Please retry")
  }

  err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
  //Password does not match
  if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
    return user, errors.New("Invalid login credentials. Please try again")
  }
  existingUser.Password = "" // remove password
  tk := &Token{UserId: existingUser.ID}
  token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
  tokenString, _ := token.SignedString([]byte(os.Getenv("AUTH_TOKEN")))
  existingUser.Token = tokenString
  GetDB().Model(&existingUser).Update("token", tokenString)
  return existingUser, nil
}

func LogoutUser(email string) error {
  existingUser := &User{}
  err := GetDB().Table("users").Where("email = ?", email).First(existingUser).Error
  if err != nil {
    if (err == gorm.ErrRecordNotFound) {
      return errors.New("Username not found")
    }
    return errors.New("Connection error. Please retry")
  }
  tk := &Token{UserId: existingUser.ID}
  tk.ExpiresAt = time.Now().Unix()
  token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
  token.SignedString([]byte(os.Getenv("AUTH_TOKEN")))
  return nil
}