package server

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	cors "github.com/gofiber/fiber/v2/middleware/cors"
)

var app *fiber.App

func Setup() {
	// Default config
	app = fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept"}))

	app.Get("/docs/*", swagger.HandlerDefault)
	//https://docs.stripe.com/api/charges
	//LIFEHACK: Good artist copy, great artist steal

	// common - http - db (code)
	// Create - POST - Insert (Create)
	// Read - GET - Select (Retrieve)
	// Update - PUT - Update (Update)
	// Delete - DELETE - Delete (Delete)

	//Route: POST /gifts
	gifts := app.Group("/gifts")
	gifts.Post("", createGiftHandler)
	//Route: DELETE /gifts/:id
	//DELETE /gifts/gift_cneq8k9u9g5j3m6ft0v0
	gifts.Delete("/:id", deleteGiftHandler)

	gifts.Get("", getManyGiftsHandler)

	gifts.Get("/:id", getOneGiftHandler)

	gifts.Patch("/:id", updateGiftHandler)

	//Route: POST /quest
	quest := app.Group("/quest")
	quest.Post("", createQuestHandler)

	quest.Delete("/:id", deleteQuestHandler)

	gifts.Patch("/:id", updateQuestHandler)

	//Route: POST /subquest
	subquest := app.Group("/subquest")
	subquest.Post("", createSubquestHandler)

	subquest.Delete("/:id", deleteSubquestHandler)

	subquest.Get("", getManySubquestHandler)

	subquest.Get("/:id", getOneSubquestHandler)

	//Route: POST /tasks
	tasks := app.Group("/tasks")
	tasks.Post("", createTasksHandler)

	tasks.Delete("/:id", deleteTasksHandler)

	tasks.Get("", getManyTasksHandler)

	tasks.Get("/:id", getOneTasksHandler)

	tasks.Patch("/:id", updateTasksHandler)

	//Route: POST /offlineshops
	offlineshops := app.Group("/offlineshops")

	offlineshops.Post("", createOfflineShopsHandler)

	offlineshops.Delete("/:id", deleteOfflineShopsHandler)

	offlineshops.Patch("/:id", updateOfflineShopsHandler)

	//
	//request -> middleware -> handler -> response
	supersecret := app.Group("/supersecret", authMiddleware)
	supersecret.Get("", superSecretHandler)
	supersecret.Get("/1", superSecretHandler)
	supersecret.Get("/2", superSecretHandler)
	supersecret.Get("/3", superSecretHandler)

	//
	// app.Post("/register", registerHandler)

}

func Start() {
	app.Listen(":7777")
}
