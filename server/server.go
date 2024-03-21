package server

import (
	"regexp"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var validate = validator.New()

var app *fiber.App

// Custom validation — sees if IDs the user provides abides by the template "tag_{base32 random id}"
func ValidateIDFormat(tag_ string) validator.Func {
	return func(fl validator.FieldLevel) bool {
		text := fl.Field().String()
		return regexp.MustCompile(tag_ + `[a-z0-9]{20}$`).MatchString(text)
	}
}

// Custom password format validation
func ValidatePasswordFormat(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	if len(password) < 6 {
		return false
	}
	if !regexp.MustCompile(`[A-z0-9]`).MatchString(password) {
		return false
	}
	if !regexp.MustCompile(`[!@#$%^*()_+{}\[\]:;<>,.?~\-\\]`).MatchString(password) {
		return false
	}
	return regexp.MustCompile(`[A-Z]`).MatchString(password) 
}

func Setup() {
	// Default config
	app = fiber.New()
	validate.RegisterValidation("seller_", ValidateIDFormat("seller_"))
	validate.RegisterValidation("service_", ValidateIDFormat("service_"))
	validate.RegisterValidation("user_", ValidateIDFormat("user_"))
	validate.RegisterValidation("password", ValidatePasswordFormat)
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
	gifts.Post("", createGiftHandler, adminMiddleware)
	//Route: DELETE /gifts/:id
	//DELETE /gifts/gift_cneq8k9u9g5j3m6ft0v0
	gifts.Delete("/:id", adminMiddleware, deleteGiftHandler)

	gifts.Get("", getManyGiftsHandler)

	gifts.Get("/:id", getOneGiftHandler)

	gifts.Patch("/:id", updateGiftHandler, authMiddleware)

	// ??
	bookedGift := app.Group("/booked_gifts", authMiddleware)
	bookedGift.Post("", createBookedGiftInWishlist)

	//!!!!
	bookedGift.Delete("/:gift_id", deleteBookedGiftInWishlist)

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

	serviceReviews := app.Group("/serviceReviews")
	serviceReviews.Get("/:id", getOneServiceReviewHandler)
	serviceReviews.Get("/service/:service_id", getSingleServiceReviewHandler)
	serviceReviews.Get("", getManyServiceReviewsHandler)
	serviceReviews.Post("", authMiddleware, createServiceReviewHandler)
	serviceReviews.Patch("/:id", authMiddleware, updateServiceReviewHandler)
	serviceReviews.Delete("/:id", authMiddleware, deleteServiceReviewHandler)

	sellersServices := app.Group("/sellerToService")
	sellersServices.Get("/:id", getOneSellerToServiceHandler)
	sellersServices.Get("", getManySellerToServiceHandler)
	sellersServices.Post("", createSellerToServiceHandler)
	sellersServices.Delete("/:id", deleteSellerToServiceHandler)

	services := app.Group("/services")
	services.Get("/:id", getOneServiceHandler)
	services.Get("", getManyServicesHandler)
	services.Post("", createServiceHandler, sellerAuthMiddleware)
	services.Patch("/:id", updateServiceHandler, sellerAuthMiddleware)
	services.Delete("/:id", deleteServiceHandler, sellerAuthMiddleware)

	//Route: POST /quest
	quest := app.Group("/quest")
	quest.Post("", adminMiddleware, createQuestHandler)

	quest.Delete("/:id", adminMiddleware, deleteQuestHandler)

	gifts.Patch("/:id", adminMiddleware, updateQuestHandler)

	//Route: POST /subquest
	subquest := app.Group("/subquest")
	subquest.Post("", adminMiddleware, createSubquestHandler)

	subquest.Delete("/:id", adminMiddleware, deleteSubquestHandler)

	subquest.Get("", getManySubquestHandler)

	subquest.Get("/:id", getOneSubquestHandler)

	//Route: POST /tasks
	tasks := app.Group("/tasks")
	tasks.Post("", adminMiddleware, createTasksHandler)

	tasks.Delete("/:id", adminMiddleware, deleteTasksHandler)

	tasks.Get("", getManyTasksHandler)

	tasks.Get("/:id", getOneTasksHandler)

	tasks.Patch("/:id", adminMiddleware, updateTasksHandler)

	//Route: POST /offlineshops
	offlineshops := app.Group("/offlineshops")

	offlineshops.Post("", adminMiddleware, createOfflineShopsHandler)

	offlineshops.Delete("/:id", adminMiddleware, deleteOfflineShopsHandler)

	offlineshops.Patch("/:id", adminMiddleware, updateOfflineShopsHandler)

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
	selection.Post("", authMiddleware, createSelectionHandler)
	//Route: DELETE /selection/:id
	//DELETE /selection/selection_cneq8k9u9g5j3m6ft0v0
	selection.Delete("/:id", adminMiddleware, deleteSelectionHandler)

	selection.Get("", getManySelectionsHandler)

	selection.Get("/:id", getOneSelectionHandler)

	selection.Patch("/:id", adminMiddleware, updateSelectionHandler)

	//
	giftToSelection := app.Group("/giftToSelection", authMiddleware)
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
	CommentToSelection := app.Group("/CommentToSelection", authMiddleware)
	CommentToSelection.Post("", createCommentToSelectionHandler)

	CommentToSelection.Patch("/:id", updateCommentToSelectionHandler)

	CommentToSelection.Get("", getCommentsToSelectionHandler)

	CommentToSelection.Delete("/:id", deleteCommentToSelectionHandler)

	// user := app.Group("/users")
	// user.Post("", CreateUserHandler)

	app.Post("/register", registerHandler)
	app.Post("/login", loginHandler)

	app.Post("/registerSeller", registerSellerHandler)
	app.Post("/loginSeller", loginSellerHandler)
}

func Start() {
	app.Listen(":7777")
}
