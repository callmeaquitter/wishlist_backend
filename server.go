package main

import "github.com/gofiber/fiber/v2"

var app *fiber.App

func serverSetup() {
	// Default config
	app = fiber.New()

	//https://docs.stripe.com/api/charges
	//LIFEHACK: Good artist copy, great artist steal

	// common - http - db (code)
	// Create - POST - Insert (Create)
	// Read - GET - Select (Retrieve)
	// Update - PUT - Update (Update)
	// Delete - DELETE - Delete (Delete)

	//Route: POST /gifts
	app.Post("/gifts", createGiftHandler)
	//Route: DELETE /gifts/:id
	//DELETE /gifts/gift_cneq8k9u9g5j3m6ft0v0
	app.Delete("/gifts/:id", deleteGiftHandler)
}

func serverStart() {
	app.Listen(":7777")
}
