package http

import (
	"api/golang/internal/handler"

	"github.com/gin-gonic/gin"
)

func StartServer() error {
	engine := gin.Default()
	gin.EnableJsonDecoderDisallowUnknownFields()

	engine.GET("/posts", handler.HealthCheck)
	engine.GET("/posts", handler.GetPosts)
	engine.POST("/posts", handler.CreatePost)
	engine.GET("/posts/:id", handler.GetPost)
	engine.PUT("/posts/:id", handler.UpdatePost)
	engine.PATCH("/posts/:id", handler.PartialUpdatePost)
	engine.DELETE("/posts/:id", handler.DeletePost)

	engine.POST("/posts/:postId/comments", handler.CreateComment)
	//engine.GET("/posts/:id/comments")
	//engine.GET("/posts/:id/comments/:id")

	return engine.Run("")
}
