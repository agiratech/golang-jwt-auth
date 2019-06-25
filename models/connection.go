package models
import (
  "os"
  "fmt"

  "github.com/joho/godotenv"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

// DB connection
func init() {
  var uname, password, dbName, dbHost, port, sslMode string
  if (os.Getenv("DB_MODE") != "production") {
    err := godotenv.Load()
    if err != nil {
      fmt.Println("err: ", err)
    }
  }
  uname = os.Getenv("DB_USER")
  password = os.Getenv("DB_PASS")
  dbName = os.Getenv("DB_NAME")
  dbHost = os.Getenv("DB_HOST")
  port = os.Getenv("DB_PORT")
  sslMode = os.Getenv("SSL_MODE")
  dbInfo := fmt.Sprintf("host=%s port= %s user=%s dbname=%s sslmode=%s password=%s", dbHost, port, uname, dbName, sslMode, password)

  conn, err := gorm.Open("postgres", dbInfo)
  if err != nil {
    fmt.Println("err: ", err)
  }
  fmt.Println("DB connected...")
  db = conn

  autoDebug()
}

// DB migration
func autoDebug() {
  db.Debug().AutoMigrate(&User{}, &Post{}, &Vote{}, &Comment{})
}

// DB connection info
func GetDB() *gorm.DB {
  return db
}