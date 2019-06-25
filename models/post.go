package models
import (
  "net/url"
  "errors"
  "strconv"

  "github.com/jinzhu/gorm"
)

// Validate post data
func (post *Post) validate() error {
  if post.Title == "" {
    return errors.New("Title required")
  }
  if _, err := url.ParseRequestURI(post.Url);err != nil {
    return errors.New("Envalid URL")
  }
  if post.Description == "" {
    return errors.New("Description required")
  }
  existingUser := &User{}

  // valid user or not
  err := GetDB().Table("users").Where("id = ?", post.UserId).First(existingUser).Error
  if err != nil && err != gorm.ErrRecordNotFound {
    return errors.New("Connection error. Please retry")
  }
  if existingUser.Email == "" {
    return errors.New("User not found")
  }
  post.Username = existingUser.Username
  return nil
}

func (post *Post) CreatePost() (*Post, error) {
  if err := post.validate(); err != nil {
    return post, err
  }
  GetDB().Create(post)
  if post.ID <= 0 {
    return post, errors.New("Failed to create post, connection error")
  }
  return post, nil
}

func (post *Post) UpdatePost(id string) (*Post, error) {
  if err := post.validate(); err != nil {
    return post, err
  }
  postId, err := strconv.Atoi(id)
  if err != nil {
    return post, err
  }
  var postExist Post
  db.First(&postExist, postId) // find post with id
  if postExist.ID <= 0 {
    return post, errors.New("Failed to update post, record not exist")
  }
  db.Model(&postExist).UpdateColumns(&post)
  post.ID = postExist.ID
  return post, nil
}

func DeletePost(id string) error {
  postId, err := strconv.Atoi(id)
  if err != nil {
    return err
  }
  var postExist Post
  db.First(&postExist, postId) // find post with id
  if postExist.ID <= 0 {
    return errors.New("Failed to delete post, record not exist")
  }
  GetDB().Unscoped().Where("id = ?", postId).Delete(Post{})
  return nil
}

func ListPosts(page string, token string) ([]map[string]interface{}, int, error) {
  var posts []Post
  var user User
  var count int
  result := []map[string]interface{}{}
  pNum, err := strconv.Atoi(page)
  if err != nil {
    return result, count, err
  }
  pNum -= 1
  if (pNum <= 0) { pNum = 0 }
  GetDB().Limit(10).Offset(pNum*10).Find(&posts)
  GetDB().Where("token = ?", token).First(&user)
  for _, post := range posts {
    var vote Vote
    var voted bool
    if err := db.Where("user_id = ? AND post_id = ?", user.ID, post.ID).First(&vote).Error; gorm.IsRecordNotFoundError(err) {
      voted = false
    } else {
      voted = true
    }
    m := make(map[string]interface{})
    m["id"] = post.ID
    m["title"] = post.Title
    m["url"] = post.Url
    m["description"] = post.Description
    m["user_id"] = post.UserId
    m["username"] = post.Username
    m["points"] = post.Points
    m["voted"] = voted
    result = append(result, m)
  }
  GetDB().Table("posts").Count(&count)
  return result, count, nil
}

func GetPost(id string) (Post, error) {
  var post Post
  postId, err := strconv.Atoi(id)
  if err != nil {
    return post, err
  }
  db.First(&post, postId) // find post with id
  if post.ID <= 0 {
    return post, errors.New("Failed to fetch post, record not exist")
  }
  return post, nil
}