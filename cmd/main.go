package main

import (
	"fmt"
	"log"
	"os"

	"gogogo/internal/api"
	"gogogo/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()
	fmt.Println("Hello, Go!")

	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&model.User{})

	r := gin.Default()
	r.POST("/login", api.GinLoginHandlerGorm(db))
	r.POST("/users", api.GinCreateUserHandlerGorm(db))

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
