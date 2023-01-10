package main

import (
	"example.com/sarang-apis/app"
	"example.com/sarang-apis/configs"
	"example.com/sarang-apis/repository"
	"example.com/sarang-apis/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	appRoute := fiber.New()
	appRoute.Use(cors.New())

	configs.ConnectDB()
	dbClient := configs.GetCollection(configs.DB, "users")

	UserRepositoryDb := repository.NewUserRepositoryDb(dbClient)
	td := app.UserHandler{Service: services.NewUserService(UserRepositoryDb)}

	appRoute.Post("/api/user", td.CreateUser)
	appRoute.Get("/api/users", td.GetAllUser)
	appRoute.Delete("/api/user/:id", td.DeleteUser)
	appRoute.Put("/api/user/:id", td.UpdateUser)

	appRoute.Listen(":8080")

}
