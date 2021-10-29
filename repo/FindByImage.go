package repo

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func FindByImage(c *gin.Context) {
	image := c.Param("image")
	log.Println(image)
	for _, f := range ListFood {
		if f.Image == image {
			c.IndentedJSON(http.StatusOK,f)
			return
		}
	}
}