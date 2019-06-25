package models

// If exist, delete it. If not exist create it.
func VoteThePost(vote Vote) {
  var existedVote Vote
  GetDB().Where("user_id = ? AND post_id = ?", vote.UserId, vote.PostId).First(&existedVote)
  if existedVote.ID <= 0 {
    GetDB().Create(&vote)
    GetDB().Exec("UPDATE posts SET points = points + 1 WHERE id = ?", vote.PostId)
  } else {
    GetDB().Unscoped().Delete(&existedVote)
    GetDB().Exec("UPDATE posts SET points = points - 1 WHERE id = ?", vote.PostId)
  }
}