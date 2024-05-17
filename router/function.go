package router

import (
	"github.com/gin-gonic/gin"
	"github.com/robbymaul/golang-swagger.git/config"
	controller "github.com/robbymaul/golang-swagger.git/controller/user"
	"github.com/robbymaul/golang-swagger.git/middleware"
	repository "github.com/robbymaul/golang-swagger.git/repository/user"
	service "github.com/robbymaul/golang-swagger.git/service/user"
)

func InitRoute() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		auth := v1.Group("/auth", middleware.BearerToken())
		{
			repo := repository.NewUserRepository(config.DB)
			service := service.NewUserServiceImpl(repo)
			controller := controller.NewUserController(service)

			auth.POST("/register", controller.Register)
			auth.POST("/login")
		}
	}

	return router
}
