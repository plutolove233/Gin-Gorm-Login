package router

import (
	"Login/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func StartRouter()(router *gin.Engine){
	router = gin.Default()
	router.Static("/static","./static")
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",nil)
	})
	router.POST("/login",func(c *gin.Context){
		ok := models.Query(c.PostForm("name"),c.PostForm("pasw"),c.PostForm("email"),c.PostForm("phone"))
		if ok{
			c.HTML(200,"login.html",gin.H{
				"Name":c.PostForm("name"),
				"Status":"成功",
				"Pasw":c.PostForm("pasw"),
				"Email":c.PostForm("email"),
				"Phone":c.PostForm("phone"),
			})
		}else{
			c.HTML(200,"login.html",gin.H{
				"Name":c.PostForm("name"),
				"Status":"失败",
				"Pasw":c.PostForm("pasw"),
				"Email":c.PostForm("email"),
				"Phone":c.PostForm("phone"),
			})
		}
	})
	return
}
