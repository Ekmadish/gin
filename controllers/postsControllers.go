package controllers

import (
	"fmt"

	"gin-test.com/test/inits"
	"gin-test.com/test/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	//get data of req body

	var content struct {
		Body  string
		Title string
	}
	c.Bind(&content)

	//reate post

	post := models.Post{Title: content.Title, Body: content.Body}
	result := inits.GetDB().Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	//return it

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {

	//get the post id

	var posts []models.Post
	inits.GetDB().Find(&posts)

	//reposnce with  them
	c.JSON(200, gin.H{
		"posts": posts,
	})

}

func PostShow(c *gin.Context) {

	//Get id off url
	id := c.Param("id")
	fmt.Println(id)

	//Get post
	var post models.Post
	inits.DB.First(&post, id)

	c.JSON(200, gin.H{
		"posts": post,
	})

}

func PostUpdate(c *gin.Context) {

	// get tje data of Body
	var content struct {
		Title string
		Body  string
	}
	c.Bind(&content)
	//Get id off url
	id := c.Param("id")
	//Get post
	var post models.Post
	inits.DB.First(&post, id)
	//update it
	inits.DB.Model(&post).Updates(models.Post{
		Title: content.Title,
		Body:  content.Body,
	})

	//reposce with it

	c.JSON(200, gin.H{
		"postsss": post,
	})

}

func PostDelete(c *gin.Context) {
	//Get id off url
	id := c.Param("id")

	//delete the post

	inits.DB.Delete(&models.Post{}, id)

	//responce with it
	c.Status(200)

}
