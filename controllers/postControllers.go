package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sankrohan/go-api/initializers"
	"github.com/sankrohan/go-api/models"
)

func PostsCreate(c *gin.Context) {

	// GET DATA FROM REQUEST BODY

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// CREATING A POST

	// static values
	// post := models.Post{Title: "Some Title", Body: "Some Body"}

	// dynamic values
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post) // pass pointer of data to Create

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Error creating post",
		})
	}

	// return the post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {

	// GET THE POST
	var posts []models.Post
	initializers.DB.Find(&posts)

	// RESPOND WITH THEM
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {

	// GET ID FROM URL
	id := c.Param("id")

	// GET THE POST
	var post []models.Post
	initializers.DB.First(&post, id)

	// RESPOND WITH THEM
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {

	// GET ID FROM URL
	id := c.Param("id")

	// GET DATA FROM REQUEST BODY
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	// FIND THE POST THAT NEEDS TO BE UPDATED
	var post models.Post
	initializers.DB.First(&post, id)

	// UPDATE THE POST
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})
	// RESPOND WITH THEM
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")

	// DELETE THE POST
	initializers.DB.Delete(&models.Post{}, id)

	// RESPOND WITH THEM
	c.Status(200)
}
