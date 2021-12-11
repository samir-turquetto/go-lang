package http

import (
	"facec/blog/internal/handler"
	"facec/blog/internal/middleware"

	"github.com/gin-gonic/gin"
)

func StartServer() error {
	engine := gin.Default()
	gin.EnableJsonDecoderDisallowUnknownFields()
	engine.Use(middleware.TokenValidate())

	engine.GET("/health", handler.HealthCheck)
	engine.POST("/users", handler.CreateUser)
	engine.POST("/login", handler.Login)

	posts := engine.Group("/posts")
	posts.GET("/posts", handler.GetPosts)
	posts.POST("/posts", handler.CreatePost)
	posts.GET("/posts/:id", handler.GetPost)
	posts.PUT("/posts/:id", handler.UpdatePost)
	posts.PATCH("/posts/:id", handler.PartialUpdatePost)
	posts.DELETE("/posts/:id", handler.DeletePost)
	posts.POST("/posts/:id/comments", handler.CreateComment)
	posts.GET("/posts/:id/comments", handler.GetCommentsByPost)
	posts.GET("/posts/:id/comments/:commentId", handler.GetCommentByCommentAndPost)

	return engine.Run()
}
