package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jigmetnamgyal/globalsearch-go/controllers"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	r := gin.Default()

	r.GET("/search", controllers.GlobalSearch)

	r.POST("/create-index", controllers.CreateIndex)

	r.Run(":3000")
}
