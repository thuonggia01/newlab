package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddFood(c *gin.Context) {
	c.HTML(http.StatusOK,"add.html",gin.H{
	})
}
