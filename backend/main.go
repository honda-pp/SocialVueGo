package main

import (
	"github.com/honda-pp/SocialVueGo/backend/app"
	"github.com/honda-pp/SocialVueGo/backend/infrastructure/database"
	"github.com/honda-pp/SocialVueGo/backend/infrastructure/logger"
)

func main() {
	log := logger.InitLogger()
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	application := app.NewApplication(db)

	err = application.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
