package main

import (
	"github.com/fentezi/session-auth/config"
	"github.com/fentezi/session-auth/internal/controllers"
	"github.com/fentezi/session-auth/internal/middlewares"
	repostirories "github.com/fentezi/session-auth/internal/repositories"
	"github.com/fentezi/session-auth/internal/service"
	"github.com/gin-gonic/gin"
)

func init() {
	config.MustInitEnv()
	config.MustConnectToSQLDb()
	config.MustConnectToRedis()
}

func main() {
	config, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	repo := repostirories.NewRepositories()
	serv := service.NewService(repo)
	controller := controllers.NewAuthorizedController(serv)
	r := gin.Default()
	r.POST("/signup", controller.SignUp)
	r.POST("/signin", middlewares.SignInMiddleware(repo), controller.SignIn)
	r.GET("/", middlewares.SessionMiddleware(repo), controller.Home)

	r.Run(config.Server.Host + ":" + config.Server.Port)
}
