package main

import (
	"Demo2APP/handler"
	"Demo2APP/repo"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("assets","./assets")
	router.LoadHTMLGlob("./assets/**/*.html")
	router.GET("/add", handler.AddFood)
	router.POST("/save", repo.SaveFood)
	router.GET("/index", repo.GetListFood)
	router.GET("/detail/:image", repo.FindByImage)
	router.Run(":8080")
}
