package controllers

import (
	"user-auth/models"

	"github.com/gin-gonic/gin"
)

// Upvote or Downvote the post
var VoteThePost = func(server *gin.Context) {
	var vote models.Vote
	server.BindJSON(&vote)
	models.VoteThePost(vote)
	server.JSON(200, gin.H{"status": "success", "message": "Upvoted/Downvoted successfully"})
}
