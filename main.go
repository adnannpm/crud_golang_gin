package main

import (
	"tugas3/controller"
	"tugas3/helper"
	"tugas3/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db, _ := helper.GetConnection()

	userRepository := repository.NewUserRepository(db)
	userController := controller.NewUserController(userRepository)

	publisherRepository := repository.NewPublisherRepository(db)
	publisherController := controller.NewPublisherController(publisherRepository)

	public := router.Group("/api")
	{
		public.POST("/register", userController.Register)
		public.POST("/login", userController.Login)

		public.GET("/publisher", publisherController.GetAllData)
		public.GET("/publisher/:id", publisherController.GetData)
		public.POST("/publisher", publisherController.AddPublisher)
		public.PATCH("/publisher/:id", publisherController.Update)
		public.DELETE("/publisher/:id", publisherController.Delete)
	}

	router.Run(":8080")
}