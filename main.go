package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sankrohan/go-api/controllers"
	"github.com/sankrohan/go-api/initializers"
)

func init() {
	initializers.LoadEnvVairables()
	initializers.ConnectToDB()

}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)       // Creates a post
	r.GET("/posts", controllers.PostsIndex)         // Gets all post
	r.GET("/posts/:id", controllers.PostsShow)      // Gets post by ID
	r.PUT("/posts/:id", controllers.PostsUpdate)    // Updates post by ID
	r.DELETE("/posts/:id", controllers.PostsDelete) // Delete post by ID

	r.Run() // router listens on configured port

}
