package controllers

import (
	"user-auth/models"

	"github.com/gin-gonic/gin"
)

// Create the comment
var CreateComment = func(server *gin.Context) {
	var cmnt models.Comment
	server.BindJSON(&cmnt)
	resp, err := cmnt.CreateComment()
	if err == nil {
		server.JSON(200, gin.H{"status": "success", "message": "Comment created successfully", "data": resp})
	} else {
		server.JSON(400, gin.H{"status": "failed", "error": err.Error()})
	}
}

// Update the comment
var UpdateComment = func(server *gin.Context) {
	var cmnt models.Comment
	server.BindJSON(&cmnt)
	resp, err := cmnt.UpdateComment(server.Param("id"))
	if err == nil {
		server.JSON(200, gin.H{"status": "success", "message": "Comment updated successfully", "data": resp})
	} else {
		server.JSON(400, gin.H{"status": "failed", "error": err.Error()})
	}
}

// Delete the comment
var DeleteComment = func(server *gin.Context) {
	err := models.DeleteComment(server.Param("id"))
	if err == nil {
		server.JSON(200, gin.H{"status": "success", "message": "Comment deleted successfully"})
	} else {
		server.JSON(400, gin.H{"status": "failed", "error": err.Error()})
	}
}

// List the comments for the post
var ListComments = func(server *gin.Context) {
	resp, err := models.ListComments(server.Param("id"))
	if err == nil {
		server.JSON(200, gin.H{"status": "success", "data": resp})
	} else {
		server.JSON(400, gin.H{"status": "failed", "error": err.Error()})
	}
}
