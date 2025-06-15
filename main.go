package main

import (
	"2BENGENHARIA7S/controller"
	"2BENGENHARIA7S/middleware"
	"2BENGENHARIA7S/model"
	"2BENGENHARIA7S/service"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=1 dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to postgres database: %v", err)
	}

	var count int64
	db.Raw("SELECT count(*) FROM pg_database WHERE datname = 'mtg_db'").Count(&count)
	if count == 0 {
		err = db.Exec("CREATE DATABASE mtg_db").Error
		if err != nil {
			log.Fatalf("Failed to create database: %v", err)
		}
		fmt.Println("Database 'mtg_db' created successfully")
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}
	sqlDB.Close()

	dsn = "host=localhost user=postgres password=1 dbname=mtg_db port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to mtg_db: %v", err)
	}

	service.InitDB(DB)

	err = DB.AutoMigrate(&model.Card{}, &model.User{})
	if err != nil {
		log.Fatalf("Failed to migrate models: %v", err)
	}
	fmt.Println("Database migration completed successfully")
}

func main() {
	InitDB()

	r := gin.Default()

	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/cards", controller.GetCards)
		protected.GET("/cards/:id", controller.GetCardByID)
		protected.POST("/cards", controller.CreateCard)
		protected.PUT("/cards/:id", controller.UpdateCard)
		protected.DELETE("/cards/:id", controller.DeleteCard)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Server is running on port %s\n", port)
	r.Run(":" + port)
}
