package server

import (
	

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

var app *fiber.App

func Setup() {
	// Default config
	app = fiber.New()

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

	sellers := app.Group("/sellers")
	sellers.Post("", createSellerHandler)
	sellers.Get("", getManySellersHandler)
	sellers.Get("/:id", getOneSellerHandler)
	sellers.Patch("/:id", updateSellerHandler)
	sellers.Delete("/:id", deleteSellerHandler)

	sellersServices := app.Group("/sellersServices")
	sellersServices.Post("", createSellersServicesHandler)
	sellersServices.Get("", getManySellersServicesHandler)
	sellersServices.Get("/:id", getOneSellersServicesHandler)
	sellersServices.Delete("/:id", deleteSellersServicesHandler)

	services := app.Group("/services")
	services.Post("", createServiceHandler)
	services.Get("", getManyServicesHandler)
	services.Get("/:id", getOneServiceHandler)
	services.Patch("/:id", updateServiceHandler)
	services.Delete("/:id", deleteServiceHandler)

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
