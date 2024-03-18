package server

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	gifts.Post("", createGiftHandler, authMiddleware)
	//Route: DELETE /gifts/:id
	//DELETE /gifts/gift_cneq8k9u9g5j3m6ft0v0
	gifts.Delete("/:id", deleteGiftHandler, authMiddleware)

	gifts.Get("", getManyGiftsHandler)

	gifts.Get("/:id", getOneGiftHandler)

	gifts.Patch("/:id", updateGiftHandler, authMiddleware)

	// ??
	bookedGift := app.Group("/booked_gifts")
	bookedGift.Post("", createBookedGiftInWishlist, authMiddleware)

	//!!!!
	bookedGift.Delete("/:gift_id", deleteBookedGiftInWishlist, authMiddleware)

	//http://localhost:7777/booked_gifts/:user_id
	bookedGift.Get("/:user_id", findUserBookedGifts)

	giftCategory := app.Group("/gift_category")
	giftCategory.Post("", createGiftCategory)

	giftCategory.Delete("/:id", deleteGiftCategory)

	giftReview := app.Group("/gift_review")
	giftReview.Post("", createGiftReviwHandler, authMiddleware)

	giftReview.Delete("/:id", deleteGiftReviewHandler, authMiddleware)

	giftReview.Get("/review/:id", getGiftReviewByIDHandler)

	giftReview.Get("/gift/:gift_id", getGiftReviewsByGiftIDHandler)

	giftReview.Get("/mark/:gift_id", calculateAverageMarkByGiftIDHandler)

	//
	//request -> middleware -> handler -> response
	supersecret := app.Group("/supersecret", authMiddleware)
	supersecret.Get("", superSecretHandler)
	supersecret.Get("/1", superSecretHandler)
	supersecret.Get("/2", superSecretHandler)
	supersecret.Get("/3", superSecretHandler)

	//
	// app.Post("/register", registerHandler)

	wishlists := app.Group("/wishlists", authMiddleware)
	wishlists.Get("", FindManyWishlistsHandler)
	wishlists.Post("", CreateWishlistHandler)
	wishlists.Put("/:id", UpdateWishlist)
	wishlists.Delete("/:id/:gift_id/:user_id", DeleteWishlistHandler)

	wishes := app.Group("/wishes", authMiddleware)
	wishes.Get("/:wishlist_id", FindManyWishlistsHandler)
	wishes.Post("/:gift_id/:wishlist_id", AddWishHandler)
	wishes.Delete("/:wishlist_id/:gift_id", DeleteWishHandler)

	// user := app.Group("/users")
	// user.Post("", CreateUserHandler)

	app.Post("/register", registerHandler)
	app.Post("/login", loginHandler)
}

func Start() {
	app.Listen(":7777")
}
