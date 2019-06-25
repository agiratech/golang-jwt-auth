package models
import (
  "errors"
  "strconv"

  "github.com/jinzhu/gorm"
)

// Validate post data
func (cmnt *Comment) validate() error {
  if cmnt.Text == "" {
    return errors.New("Comment text required")
  }
  
  existingUser := &User{}

  // valid user or not
  err := GetDB().Table("users").Where("id = ?", cmnt.UserId).First(existingUser).Error
  if err != nil && err != gorm.ErrRecordNotFound {
    return errors.New("Connection error. Please retry")
  }
  if existingUser.Email == "" {
    return errors.New("User not found")
  }

  existingPost := &Post{}

  // post exist or not to like
  err = GetDB().Table("posts").Where("id = ?", cmnt.PostId).First(existingPost).Error
  if err != nil && err != gorm.ErrRecordNotFound {
    return errors.New("Connection error. Please retry")
  }
  if existingPost.ID <= 0 {
    return errors.New("Post not found")
  }
  return nil
}

func (cmnt *Comment) CreateComment() (*Comment, error) {
  if err := cmnt.validate(); err != nil {
    return cmnt, err
  }
  GetDB().Create(cmnt)
  if cmnt.ID <= 0 {
    return cmnt, errors.New("Failed to create comment, connection error")
  }
  return cmnt, nil
}

func (cmnt *Comment) UpdateComment(id string) (*Comment, error) {
  if err := cmnt.validate(); err != nil {
    return cmnt, err
  }

  cmntId, err := strconv.Atoi(id)
  if err != nil {
    return cmnt, err
  }
  var cmntExist Comment
  db.First(&cmntExist, cmntId) // find comment with id
  if cmntExist.ID <= 0 {
    return cmnt, errors.New("Failed to update comment, record not exist")
  }
  db.Model(&cmntExist).UpdateColumns(&cmnt)
  return cmnt, nil
}

func DeleteComment(id string) error {
  cmntId, err := strconv.Atoi(id)
  if err != nil {
    return err
  }
  var cmntExist Comment
  db.First(&cmntExist, cmntId) // find comment with id
  if cmntExist.ID <= 0 {
    return errors.New("Failed to delete comment, record not exist")
  }
  GetDB().Unscoped().Where("id = ?", cmntId).Delete(Comment{})
  return nil
}

func ListComments(id string) ([]Comment, error) {
  var comments []Comment
  postId, err := strconv.Atoi(id)
  if err != nil {
    return comments, err
  }
  db.Where("post_id = ?", postId).Find(&comments)
  return comments, nil
}