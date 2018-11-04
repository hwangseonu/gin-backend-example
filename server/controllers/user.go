package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/gin-backend/server"
)

func CreateUser(c *gin.Context) {
	var payload struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Nickname string `json:"nickname" binding:"required"`
		Email    string `json:"email" binding:"required"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil{
		c.AbortWithError(400, err).SetType(gin.ErrorTypeBind)
		return
	}
	server.DB.C("users").Insert(payload)
}