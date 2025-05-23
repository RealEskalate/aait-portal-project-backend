package main

import (
	"log"
	"os"

	database "github.com/Elizabethyonas/A2SV-Portal-Project/cmd/infrastructure/database"
	"github.com/Elizabethyonas/A2SV-Portal-Project/cmd/router"
	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/infrastructure"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.NewDatabase()
	if err != nil {
		log.Printf("Failed to connect to database: %s", err.Error())
		return
	}

	// Run migrations
	if err := infrastructure.Migrate(db); err != nil {
		log.Printf("Failed to run migrations: %s", err.Error())
		return
	}
	ensureUploadDirExists()

	app := gin.Default()
	app.Static("/uploads", "./uploads")
	router.SetupRoutes(app, db)

	port := ":8080"
	log.Printf("Server running on http://localhost%s", port)
	err = app.Run(port)
	if err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}

func ensureUploadDirExists() {
	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		os.Mkdir("uploads", os.ModePerm)
	}
}
