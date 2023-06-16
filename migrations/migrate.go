package main

import (
	"log"

	"github.com/sankrohan/go-api/initializers"
	"github.com/sankrohan/go-api/models"
)

func init() {
	initializers.LoadEnvVairables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
	log.Fatal("Migration Applied Successfully!")
}
