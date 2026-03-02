package main

import (
	"rest-api-gin-go/internal/database"

	"github.com/gin-gonic/gin"
)

func (app *application) GetUserFromContext(c *gin.Context) *database.User {
	// Extract user from context
	contextUser, exist := c.Get("user")
	if !exist {
		return &database.User{}
	}
	
	user, ok := contextUser.(*database.User)
	if !ok {
		return &database.User{}
	}
	return user
}