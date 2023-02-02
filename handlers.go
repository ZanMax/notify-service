package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func notifyInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"services": []string{"mail", "sms"},
	})
}

func notifySend(c *gin.Context) {
	service := c.Param("service")
	c.JSON(http.StatusOK, gin.H{
		"service": service,
	})
}
