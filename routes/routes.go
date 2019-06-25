package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"

	"user-auth/controllers"
	"user-auth/services"
)

func Load() {
	route := gin.Default()
	// checking authentication token in header
	route.Use(services.Authorization())
	route.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE, OPTIONS",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "*",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	v1 := route.Group("/v1")
	{
		// routes for users
		v1.POST("/user/new", controllers.CreateUser)
		v1.POST("/user/login", controllers.LoginUser)
		v1.DELETE("/user/logout", controllers.LogoutUser)
		// routes for posts
		v1.POST("/posts/new", controllers.CreatePost)
		v1.PUT("/posts/:id", controllers.UpdatePost)
		v1.DELETE("/posts/:id", controllers.DeletePost)
		v1.GET("/posts", controllers.ListPosts)
		v1.GET("/posts/:id", controllers.GetPost)
		// route for vote the post
		v1.POST("/vote", controllers.VoteThePost)
		// routes for comments
		v1.POST("/comment", controllers.CreateComment)
		v1.PUT("/comment/:id", controllers.UpdateComment)
		v1.DELETE("comment/:id", controllers.DeleteComment)
		v1.GET("/comments/:id", controllers.ListComments)
	}
	route.NoRoute(func(server *gin.Context) {
		server.JSON(404, gin.H{"code": 404, "description": "Api endpoint not found"})
	})
	route.Run(":8000")
}
