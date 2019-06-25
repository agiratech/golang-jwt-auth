package controllers

import (
	"user-auth/models"

	"github.com/gin-gonic/gin"
)

var CreateUser = func(server *gin.Context) {
	var user models.User
	server.BindJSON(&user)
	resp, err := user.CreateUser()
	if err == nil {
		server.JSON(200, gin.H{"status": "success", "message": "User created successfully", "data": resp})
	} else {
		server.JSON(400, gin.H{"status": "failed", "error": err.Error()})
	}
}

var LoginUser = func(server *gin.Context) {
	var user models.User
	server.BindJSON(&user)
	resp, err := user.LoginUser()
	if err == nil {
		server.JSON(200, gin.H{"status": "success", "message": "User logged in successfully", "data": resp})
	} else {
		server.JSON(400, gin.H{"status": "failed", "error": err.Error()})
	}
}

var LogoutUser = func(server *gin.Context) {
	email := server.Request.Header.Get("X-User-Email")
	err := models.LogoutUser(email)
	if err == nil {
		server.JSON(200, gin.H{"status": "success", "message": "User logged out successfully"})
	} else {
		server.JSON(400, gin.H{"status": "failed", "error": err.Error()})
	}
}
