package handler

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")

	r.GET("/", GetAll)
	//r.POST("/FilterStudents", FilterStudents)
	r.POST("/CreateStudent", CreateStudent)
	r.POST("/DeleteStudent/:id", DeleteStudent)
	return r
}
