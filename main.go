package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"blog/config"
	"blog/controllers"
	"blog/models"
	"blog/routes"
)

func main() {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	if err := config.Connect(dbUser, dbPass, dbHost, dbName); err != nil {
		log.Fatal(err)
	}

	// auto migrate
	config.DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Category{})

	// create default admin if not exists
	var count int64
	config.DB.Model(&models.User{}).Count(&count)
	if count == 0 {
		hash, _ := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
		config.DB.Create(&models.User{Username: "admin", PasswordHash: string(hash)})
		log.Println("created default admin user: admin/admin")
	}

	r := gin.Default()
	controllers.SetupSession("secret", r)
	r.LoadHTMLGlob("templates/**/*")
	r.Static("/public", "public")
	r.Static("/uploads", "uploads")

	routes.Register(r)

	log.Println("Server started at :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
