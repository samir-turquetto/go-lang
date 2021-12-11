package handler

import (
	"facec/blog/internal/commons"
	"facec/blog/internal/container"
	"facec/blog/pkg"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var userRequest pkg.User
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(400, pkg.NewResponseError("invalid user", err))
		return
	}

	user := container.UserRepository.GetUser(userRequest.Login)

	if user == nil || userRequest.Password != user.Password {
		c.JSON(401, pkg.ResponseError{Cause: "user unauthorized"})
		return
	}

	token, err := commons.GenerateToken(user.Login)
	if err != nil {
		c.JSON(500, pkg.NewResponseError("fail to generate toke", err))
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}
