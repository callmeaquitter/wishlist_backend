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

	//

	selection := app.Group("/selection")
	selection.Post("", createSelectionHandler)
	//Route: DELETE /selection/:id
	//DELETE /selection/selection_cneq8k9u9g5j3m6ft0v0
	selection.Delete("/:id", deleteSelectionHandler)

	selection.Get("", getManySelectionsHandler)

	selection.Get("/:id", getOneSelectionHandler)

	selection.Patch("/:id", updateSelectionHandler)

	//
	giftToSelection := app.Group("/giftToSelection")
	giftToSelection.Post("", createGiftToSelectionHandler)

	giftToSelection.Delete("/:id", deleteGiftToSelectionHandler)

	giftToSelection.Get("", findGiftToSelectionHandler)

	giftToSelection.Patch("/:id", updateGiftToSelectionHandler)

	//
	SelectionCategory := app.Group("/SelectionCategory")
	SelectionCategory.Post("", createSelectionCategoryHandler)

	SelectionCategory.Patch("/:id", updatedSelectionCategoryHandler)

	SelectionCategory.Get("", findSelectionCategoryHandler)

	SelectionCategory.Delete("/:id", deleteSelectionCategoryHandler)

	//
	LikeToSelection := app.Group("/LikeToSelection")
	LikeToSelection.Post("", createLikeToSelectionHandler)

	LikeToSelection.Get("", getLikesCountToSelectionHandler)

	LikeToSelection.Delete("/:id", deleteLikeToSelectionHandler)

	//
	CommentToSelection := app.Group("/CommentToSelection")
	CommentToSelection.Post("", createCommentToSelectionHandler)

	CommentToSelection.Patch("/:id", updateCommentToSelectionHandler)

	CommentToSelection.Get("", getCommentsToSelectionHandler)

	CommentToSelection.Delete("/:id", deleteCommentToSelectionHandler)

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
