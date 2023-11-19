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

	followRepo := repositories.NewFollowRepository(db)
	followUsecase := usecases.NewFollowUsecase(followRepo)
	followHandler := handlers.NewFollowHandler(*followUsecase)

	tweetRepo := repositories.NewTweetRepository(db)
	tweetUsecase := usecases.NewTweetUsecase(tweetRepo)
	tweetHandler := handlers.NewTweetHandler(*tweetUsecase)

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
		api.GET("/sessionInfo", userHandler.GetSessionInfo)
		api.GET("/userList", userHandler.GetUserList)
		api.GET("/followingUserList/:userID", userHandler.GetFollowingUserList)
		api.GET("/followerUserList/:userID", userHandler.GetFollowerUserList)
		api.GET("/userInfo/:userID", userHandler.GetUserInfo)
		api.GET("/userIcon/:iconPath", userHandler.GetUserIcon)
		api.POST("/editProfile", userHandler.EditProfile)

		api.POST("/follow/:userID", followHandler.FollowUser)
		api.POST("/unfollow/:userID", followHandler.UnfollowUser)
		api.GET("/followerIDs", followHandler.GetFollowerIDs)
		api.GET("/followerIDs/:userID", followHandler.GetFollowerIDs)
		api.GET("/followingIDs", followHandler.GetFollowingIDs)
		api.GET("/followingIDs/:userID", followHandler.GetFollowingIDs)

		api.GET("/tweetList", tweetHandler.GetTweetList)
		api.GET("/tweetList/:userID", tweetHandler.GetTweetListByUserID)
		api.POST("/createTweet", tweetHandler.CreateTweet)
	}

	return &Application{
		Router: router,
	}
}

func (app *Application) Run(addr string) error {
	return app.Router.Run(addr)
}
