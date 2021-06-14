package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	_ = godotenv.Load()
	r := gin.Default()

	InitCovid()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Running.")
	})

	r.GET("/covid", func(c *gin.Context) {
		c.JSON(200, GetCovidData())
	})

	err := r.Run(":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("Error starting server.\n%s", err.Error())
	}
	log.Println("Started server on port " + os.Getenv("PORT"))
}
