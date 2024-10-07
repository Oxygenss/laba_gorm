package handler

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")

	router.GET("/", GetAll)
	router.POST("/CreateStudent", CreateStudent)
	router.POST("/DeleteStudent/:id", DeleteStudent)
	return router
}
