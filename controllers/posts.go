package controllers

import (
	"user-auth/models"

	"github.com/gin-gonic/gin"
)

// Create the post
var CreatePost = func(server *gin.Context) {
	var post models.Post
	server.BindJSON(&post)
	resp, err := post.CreatePost()
	if err == nil {
		server.JSON(200, gin.H{"status": "success", "message": "Post created successfully", "data": resp})
	} else {
		server.JSON(400, gin.H{"status": "failed", "error": err.Error()})
	}
}

// Update the post
var UpdatePost = func(server *gin.Context) {
	var post models.Post
	server.BindJSON(&post)
	resp, err := post.UpdatePost(server.Param("id"))
	if err == nil {
		server.JSON(200, gin.H{"status": "success", "message": "Post updated successfully", "data": resp})
	} else {
		server.JSON(400, gin.H{"status": "failed", "error": err.Error()})
	}
}

// Delete the post
var DeletePost = func(server *gin.Context) {
	err := models.DeletePost(server.Param("id"))
	if err == nil {
		server.JSON(200, gin.H{"status": "success", "message": "Post deleted successfully"})
	} else {
		server.JSON(400, gin.H{"status": "failed", "error": err.Error()})
	}
}

// Get all posts
var ListPosts = func(server *gin.Context) {
	token, exist := server.Get("token")
	if !exist {
		token = "notExist"
	}
	data, count, err := models.ListPosts(server.Query("page"), token.(string))
	if err == nil {
		server.JSON(200, gin.H{"status": "success", "data": data, "count": count})
	} else {
		server.JSON(400, gin.H{"status": "failed", "error": err.Error()})
	}
}

// Get post details
var GetPost = func(server *gin.Context) {
	resp, err := models.GetPost(server.Param("id"))
	if err == nil {
		server.JSON(200, gin.H{"status": "success", "data": resp})
	} else {
		server.JSON(400, gin.H{"status": "failed", "error": err.Error()})
	}
}
