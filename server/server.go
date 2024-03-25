package server

import (
	"os"
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
	app = fiber.New(fiber.Config{
		BodyLimit: 5 * 1024 * 1024, // Limit file size to 5MB
	})

	validate.RegisterValidation("role_", ValidateIDFormat("role_"))
	validate.RegisterValidation("gift_", ValidateIDFormat("gift_"))
	validate.RegisterValidation("wishlist_", ValidateIDFormat("wishlist_"))
	validate.RegisterValidation("selection_", ValidateIDFormat("selection_"))
	validate.RegisterValidation("seller_", ValidateIDFormat("seller_"))
	validate.RegisterValidation("service_", ValidateIDFormat("service_"))
	validate.RegisterValidation("user_", ValidateIDFormat("user_"))
	validate.RegisterValidation("password", ValidatePasswordFormat)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, x-api-key",
	}))

	app.Static("/", "./public/gifts/")

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
	gifts.Post("", adminMiddleware, createGiftHandler)
	//Route: DELETE /gifts/:id
	//DELETE /gifts/gift_cneq8k9u9g5j3m6ft0v0
	gifts.Delete("/:id", adminMiddleware, deleteGiftHandler)

	gifts.Get("", getManyGiftsHandler)

	gifts.Get("/:id", getOneGiftHandler)

	gifts.Patch("/:id", updateGiftHandler, authMiddleware)

	// Обработчик для загрузки файлов
	app.Post("/upload", uploadHandler)

	// ??
	bookedGift := app.Group("/booked_gifts", authMiddleware)
	bookedGift.Post("", createBookedGiftInWishlist)

	//!!!!
	bookedGift.Delete("/:gift_id", deleteBookedGiftInWishlist)

	//http://localhost:7777/booked_gifts/:user_id
	bookedGift.Get("/:user_id", findUserBookedGifts)

	giftCategory := app.Group("/gift_category")
	giftCategory.Post("", adminMiddleware, createGiftCategory)

	giftCategory.Delete("/:id", adminMiddleware, deleteGiftCategory)

	giftCategory.Get("", getManyGiftsCategoryHandler)

	giftReview := app.Group("/gift_review")
	giftReview.Post("", authMiddleware, createGiftReviwHandler)

	giftReview.Delete("/:id", authMiddleware, deleteGiftReviewHandler)

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
	sellersServices.Post("", adminMiddleware, createSellerToServiceHandler)
	sellersServices.Delete("/:id", deleteSellerToServiceHandler)

	services := app.Group("/services")
	services.Get("/seller/:seller_id", getSingleServiceHandler)
	services.Get("/:id", getOneServiceHandler)
	services.Get("", getManyServicesHandler)
	services.Post("", sellerAuthMiddleware, createServiceHandler)
	services.Patch("/:id", sellerAuthMiddleware, updateServiceHandler)
	services.Delete("/:id", sellerAuthMiddleware, deleteServiceHandler)

	//Route: POST /quest
	quest := app.Group("/quest")
	quest.Post("", adminMiddleware, createQuestHandler)

	quest.Get("", getManyQuestHandler)

	quest.Delete("/:id", adminMiddleware, deleteQuestHandler)

	//Route: POST /subquest
	subquest := app.Group("/subquest")
	subquest.Post("", adminMiddleware, createSubquestHandler)

	subquest.Delete("/:id", adminMiddleware, deleteSubquestHandler)

	subquest.Get("", getManySubquestHandler)

	subquest.Get("/:id", getOneSubquestHandler)

	subquest.Patch("/:id", adminMiddleware, updateSubquestHandler)

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

	offlineshops.Get("", getManyOfflineShopsHandler)

	offlineshops.Get("/:id", getOneOfflineShopsHandler)

	offlineshops.Delete("/:/:photoid", adminMiddleware, deleteOfflineShopsHandler)

	offlineshops.Patch("/:id", adminMiddleware, updateOfflineShopsHandler)

	//

	selection := app.Group("/selection")
	selection.Post("", authMiddleware, createSelectionHandler)
	//Route: DELETE /selection/:id
	//DELETE /selection/selection_cneq8k9u9g5j3m6ft0v0
	selection.Delete("/:id", deleteSelectionHandler)

	selection.Get("", getManySelectionsHandler)

	selection.Get("/:selection_id", getOneSelectionHandler)

	selection.Patch("", authMiddleware, updateSelectionHandler)

	//
	giftToSelection := app.Group("/giftToSelection", authMiddleware)
	giftToSelection.Post("", createGiftToSelectionHandler)

	giftToSelection.Delete("/:gift_id/:selection_id", deleteGiftToSelectionHandler)

	giftToSelection.Get("/:id", findGiftToSelectionHandler)

	giftToSelection.Patch("/:id", updateGiftToSelectionHandler)

	//
	SelectionCategory := app.Group("/SelectionCategory")
	SelectionCategory.Post("", authMiddleware, createSelectionCategoryHandler)

	SelectionCategory.Patch("/:id", adminMiddleware, updatedSelectionCategoryHandler)

	SelectionCategory.Get("", findManySelectionCategoryHandler)
	SelectionCategory.Get("/:id", findOneSelectionCategoryHandler)

	SelectionCategory.Delete("/:id", authMiddleware, deleteSelectionCategoryHandler)

	//
	LikeToSelection := app.Group("/LikeToSelection")
	LikeToSelection.Post("", authMiddleware, createLikeToSelectionHandler)

	LikeToSelection.Get("/:selection_id", getLikesCountToSelectionHandler)

	LikeToSelection.Delete("/:selection_id", authMiddleware, deleteLikeToSelectionHandler)

	//
	CommentToSelection := app.Group("/CommentToSelection")
	CommentToSelection.Post("", authMiddleware, createCommentToSelectionHandler)

	CommentToSelection.Patch("/:id", authMiddleware, updateCommentToSelectionHandler)

	CommentToSelection.Get("/:selection_id", getCommentsToSelectionHandler)

	CommentToSelection.Delete("/:id", authMiddleware, deleteCommentToSelectionHandler)

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
	wishlists.Get("/:name", findWishlistByNameHandler)
	wishlists.Post("", CreateWishlistHandler)
	wishlists.Put("/:id", UpdateWishlistHandler)
	wishlists.Delete("/:id/:gift_id/:user_id", DeleteWishlistHandler)

	wishes := app.Group("/wishes", authMiddleware)
	wishes.Get("/:wishlist_id", FindAllWishesInWishlistHandler)
	wishes.Post("/:gift_id/:wishlist_id", AddWishHandler)
	wishes.Delete("/:wishlist_id/:gift_id", DeleteWishHandler)

	// user := app.Group("/users")
	// user.Post("", CreateUserHandler)

	app.Post("/register", registerHandler)
	app.Post("/login", loginHandler)

	app.Post("/registerSeller", registerSellerHandler)
	app.Post("/loginSeller", loginSellerHandler)
}

func Start() {
	app.Listen(":" + os.Getenv("PORT"))
}
