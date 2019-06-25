package services

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"user-auth/models"
)

func Authorization() gin.HandlerFunc {
	return func(server *gin.Context) {
		validateToken(server)
		server.Next()
	}
}

func validateToken(server *gin.Context) {
	//List of endpoints that doesn't require auth
	notAuth := []string{"/v1/user/new", "/v1/user/login", "/v1/posts"}
	//current request path
	requestPath := server.Request.URL.Path
	requestMethod := server.Request.Method
	//check if request does not need authentication
	token := server.Request.Header.Get("Authorization")
	if token != "" {
		splitted := strings.Split(token, " ")
		if len(splitted) == 2 {
			server.Set("token", splitted[1])
		}
	}

	for _, value := range notAuth {
		if value == requestPath || requestMethod == "OPTIONS" {
			server.Next()
			return
		}
	}

	if token == "" {
		server.JSON(403, gin.H{"status": "failed", "error": "Missing auth token"})
		server.AbortWithStatus(403)
		return
	}

	splitted := strings.Split(token, " ")
	if len(splitted) != 2 {
		server.JSON(403, gin.H{"status": "failed", "error": "Invalid auth token"})
		server.AbortWithStatus(403)
		return
	}

	tokenPart := splitted[1] //Grab the token part, what we are truly interested in
	tk := &models.Token{}

	jwtToken, err := jwt.ParseWithClaims(tokenPart, tk, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("AUTH_TOKEN")), nil
	})

	//Malformed token
	if err != nil {
		server.JSON(403, gin.H{"status": "failed", "error": "Malformed authentication token"})
		server.AbortWithStatus(403)
		return
	}

	//Token is invalid
	if !jwtToken.Valid {
		server.JSON(403, gin.H{"status": "failed", "error": "Token is not valid"})
		server.AbortWithStatus(403)
		return
	}

	fmt.Sprintf("User %", tk.UserId) //Useful for monitoring
	ctx := context.WithValue(server.Request.Context(), "user", tk.UserId)
	server.Request = server.Request.WithContext(ctx)
	server.Next()
	return
}
