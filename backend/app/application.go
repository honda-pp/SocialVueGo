package app

import (
	"database/sql"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/honda-pp/SocialVueGo/backend/app/handlers"
	"github.com/honda-pp/SocialVueGo/backend/app/repositories"
	"github.com/honda-pp/SocialVueGo/backend/app/usecases"
)

type Application struct {
	Router *gin.Engine
}

func NewApplication(db *sql.DB) *Application {
	userRepo := repositories.NewUserRepository(db)
	userUsecase := usecases.NewUserUsecase(userRepo)
	userHandler := handlers.NewUserHandler(*userUsecase)

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:88"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{Domain: "localhost", Path: "/api/", MaxAge: 86400})
	router.Use(sessions.Sessions("mysession", store))

	api := router.Group("/api")
	{
		api.POST("/login", userHandler.Login)
		api.POST("/logout", userHandler.Logout)
		api.POST("/signup", userHandler.Signup)
		api.GET("/checkLoggedIn", userHandler.CheckLoggedIn)
		api.GET("/userList", userHandler.GetUserList)
	}

	return &Application{
		Router: router,
	}
}

func (app *Application) Run(addr string) error {
	return app.Router.Run(addr)
}
