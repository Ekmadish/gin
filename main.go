package main

import (
	"gin-test.com/test/controllers"
	"gin-test.com/test/inits"
	"github.com/gin-gonic/gin"
)

func init() {
	inits.LoadEnvVar()
	inits.ConnecToDB()

}
func main() {

	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/post/:id", controllers.PostUpdate)
	r.GET("/all", controllers.PostsIndex)
	r.GET("/post/:id", controllers.PostShow)
	r.DELETE("/post/:id", controllers.PostDelete)
	r.Run()
}
