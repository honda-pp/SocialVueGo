package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/honda-pp/SocialVueGo/backend/app/handlers"
	"github.com/honda-pp/SocialVueGo/backend/app/repositories"
	"github.com/honda-pp/SocialVueGo/backend/app/usecases"
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

	userRepo := repositories.NewUserRepository(db)

	userUsecase := usecases.NewUserUsecase(*userRepo)

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:88"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{Domain: "localhost", Path: "/api/", MaxAge: 86400})
	router.Use(sessions.Sessions("mysession", store))

	userHandler := handlers.NewUserHandler(*userUsecase)

	api := router.Group("/api")
	{
		api.POST("/login", userHandler.Login)
		api.POST("/logout", userHandler.Logout)
	}

	err = router.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
