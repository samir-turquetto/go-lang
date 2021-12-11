package handler

import (
	"facec/blog/internal/container"
	"facec/blog/pkg"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user pkg.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, pkg.NewResponseError("invalid user", err))
		return
	}

	container.UserRepository.CreateUser(&user)

	c.JSON(201, pkg.ResponseUser{
		Login: user.Login,
	})
}
