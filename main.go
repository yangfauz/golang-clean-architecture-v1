package main

import (
	"golang-clean-architecture-v1/app/api/user"
	"golang-clean-architecture-v1/app/handler"
	"golang-clean-architecture-v1/config"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Setup Configuration
	configuration := config.New()
	database := config.NewPostgreDatabase(configuration)

	// Setup Repository
	userRepository := user.NewRepository(database)

	// Setup Service
	userService := user.NewService(userRepository)

	// Setup Handler
	userHandler := handler.NewUserHandler(userService)

	//Setup Gin Route
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "server running..")
	})

	api := router.Group("/api/v1")
	api.GET("/users", userHandler.GetAll)
	api.GET("/users/:id", userHandler.GetByID)
	api.POST("/users", userHandler.CreateUser)
	api.PUT("/users/:id", userHandler.UpdateUser)
	api.DELETE("/users/:id", userHandler.DeleteUser)

	router.Run()
}
