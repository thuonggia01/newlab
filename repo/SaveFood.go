package repo

import (
	"Demo2APP/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var ListFood = []model.Food{
	{Image: "https://cdn.huongnghiepaau.com/wp-content/uploads/2021/01/mon-ngon-tu-trung-cut.jpg", NameFood:"bánh" , Recipe: "xào", Detail: "gọa,..."},
	{Image: "https://cdn.daynauan.info.vn/wp-content/uploads/2019/05/mon-an-truyen-thong-viet-nam.jpg", NameFood:"phở" , Recipe: "xào", Detail: "gọa,..."},
	{Image: "https://img.tastykitchen.vn/resize/764x-/2020/09/14/mam-com-ngon-1-4751.jpg", NameFood:"cơm" , Recipe: "Gỏi", Detail: "gọa,..."},
	{Image: "https://vietchef.com.vn/ckfinder/userfiles/images/S%C6%B0%E1%BB%9Dn%20heo%20kho%20%C4%91%E1%BA%ADm%20%C4%91%C3%A0.jpg", NameFood:"Thịt xào" , Recipe: "xào", Detail: "gọa,..."},
}
func GetListFood(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Food" : ListFood,
	})
}
func SaveFood(c *gin.Context) {
	newFood := model.Food{
		Image:    c.PostForm("image"),
		NameFood: c.PostForm("nameFood"),
		Recipe:   c.PostForm("recipe"),
		Detail:   c.PostForm("detail"),
	}
	ListFood= append(ListFood,newFood)
	fmt.Println("save thanh cong ",newFood)
	GetListFood(c)
}