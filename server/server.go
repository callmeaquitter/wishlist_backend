package server

import (
	"regexp"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var validate = validator.New()

var app *fiber.App

// Custom validation — sees if IDs the user provides abides by the template "tag_{base32 random id}"
func ValidateIDFormat(tag_ string) validator.Func {
	return func (fl validator.FieldLevel) bool {
		text := fl.Field().String()
		return regexp.MustCompile(tag_ + `[a-z0-9]{20}$`).MatchString(text) 
	}
}

func Setup() {
	// Default config
	app = fiber.New()

	validate.RegisterValidation("seller_", ValidateIDFormat("seller_"))
	validate.RegisterValidation("service_", ValidateIDFormat("service_"))
	validate.RegisterValidation("user_", ValidateIDFormat("user_"))

	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
	    AllowOrigins: "*",
	    AllowHeaders: "Origin, Content-Type, Accept",
	}))

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

	sellers := app.Group("/sellers")
	sellers.Get("/:id", getOneSellerHandler)
	sellers.Get("", getManySellersHandler)
	sellers.Post("", createSellerHandler)
	sellers.Patch("/:id", updateSellerHandler)
	sellers.Delete("/:id", deleteSellerHandler)

	serviceReviews := app.Group("/serviceReviews")
	serviceReviews.Get("/:id", getOneServiceReviewHandler)
	serviceReviews.Get("/service/:service_id", getSingleServiceReviewHandler)
	serviceReviews.Get("", getManyServiceReviewsHandler)
	serviceReviews.Post("", createServiceReviewHandler)
	serviceReviews.Patch("/:id", updateServiceReviewHandler)
	serviceReviews.Delete("/:id", deleteServiceReviewHandler)

	sellersServices := app.Group("/sellerToService")
	sellersServices.Get("/:id", getOneSellerToServiceHandler)
	sellersServices.Get("", getManySellerToServiceHandler)
	sellersServices.Post("", createSellerToServiceHandler)
	sellersServices.Delete("/:id", deleteSellerToServiceHandler)

	services := app.Group("/services")
	services.Get("/:id", getOneServiceHandler)
	services.Get("", getManyServicesHandler)
	services.Post("", createServiceHandler)
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

	wishlists := app.Group("/wishlists", authMiddleware)
	wishlists.Get("", FindManyWishlistsHandler)
	wishlists.Post("", CreateWishlistHandler)
	wishlists.Put("/:id", UpdateWishlist)
	wishlists.Delete("/:id/:gift_id/:user_id", DeleteWishlistHandler)

	wishes := app.Group("/wishes", authMiddleware)
	wishes.Get("/:wishlist_id", FindManyWishlistsHandler)
	wishes.Post("/:gift_id/:wishlist_id", AddWishHandler)
	wishes.Delete("/:wishlist_id/:gift_id", DeleteWishHandler)

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

	// user := app.Group("/users")
	// user.Post("", CreateUserHandler)

	app.Post("/register", registerHandler)
	app.Post("/login", loginHandler)
}

func Start() {
	app.Listen(":7777")
}
