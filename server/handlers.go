package server

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"time"

	"wishlist/db"
	_ "wishlist/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/xid"
	"github.com/shopspring/decimal"
)

type ResponseHTTP struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// TODO: Change raw data input into formData
// createGift godoc
// @Summary Creates a new gift.
// @Tags Gifts
// @Accept json
// @Produce json
// @Param Gift body db.Gift true "Create Gift"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{data=db.Gift}
// @Failure 400 {object} ResponseHTTP{}
// @Router /gifts [post]
func createGiftHandler(c *fiber.Ctx) error {
	var gift db.Gift
	if err := c.BodyParser(&gift); err != nil {
		return c.SendString(err.Error())
	}

	gift.ID = "gift_" + xid.New().String()

	err := validate.Struct(gift)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).
			SendString(err.Error())
	}

	//gift.UserID = getUserID()

	ok := db.CreateGift(gift)
	if !ok {
		return c.SendString("Error in createGift operation")
	}

	return c.JSON(gift)
}

// deleteGift godoc
// @Summary Delete a gift by ID.
// @Description Deletes a gift from the database using the provided ID.
// @Tags Gifts
// @Accept  json
// @Produce json
// @Param Authorization header string true "Bearer токен"
// @Param id path string true "Gift ID to delete"
// @Success 200 {string} string "Gift deleted successfully"
// @Failure 400 {string} string "Error in deleteGift operation"
// @Router /gifts/{id} [delete]
func deleteGiftHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	ok := db.DeleteGift(id)
	if !ok {
		return c.SendString("Error in deleteGift operation")
	}
	return c.SendString("Gift deleted successfully")
}

// GetAllGifts is a function to get all gifts data from database
// @Summary Get all gifts
// @Description Get all gifts
// @Tags Gifts
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{data=[]db.Gift}
// @Failure 503 {object} ResponseHTTP{}
// @Router /gifts [get]
func getManyGiftsHandler(c *fiber.Ctx) error {
	gifts, ok := db.FindManyGift(db.Gift{})
	// fmt.Println("Gifts:", gifts)
	if !ok {
		return c.SendString("Error in findManyGifts operation")
	}
	if len(gifts) == 0 {
		return c.SendString("No gifts found")
	}
	return c.JSON(gifts)
}

// GetOneGifts is a function to get all books data from database
// @Summary Get one gift
// @Description Get one gift
// @Tags Gifts
// @Accept json
// @Produce json
// @Param id path string true "Gift ID"
// @Success 200 {object} ResponseHTTP{data=[]db.Gift}
// @Failure 503 {object} ResponseHTTP{}
// @Router /gifts/{id} [get]
func getOneGiftHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	gift, ok := db.FindOneGift(id)
	if !ok {
		return c.SendString("Error in findOneGift operation")
	}
	return c.JSON(gift)
}

// Update Gift godoc
// @Summary update gift by ID
// @Description get the status of server.
// @Tags 	Gifts
// @Accept  json
// @Produce json
// @Param id path string true "Gift id"
// @Param Gift body db.Gift true "Update Gift"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{data=db.Gift}
// @Failure 400 {object} ResponseHTTP{}
// @Failure 500 {object} ResponseHTTP{}
// @Router /gifts/{id} [patch]
func updateGiftHandler(c *fiber.Ctx) error {
	giftID := c.Params("id")

	var updatedGift db.Gift
	if err := c.BodyParser(&updatedGift); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error parsing request body")
	}

	// err := validate.Struct(updatedGift)
	// if err != nil {
	// 	return c.Status(fiber.StatusUnprocessableEntity).
	// 		SendString(err.Error())
	// }

	ok := db.UpdateGift(giftID, updatedGift)
	if !ok {
		return c.SendString("Error in updateGift operation")
	}

	return c.SendString("Gift updated successfully")
}

// createBookedGiftInWishlist godoc
// @Summary Creates a booked gift in the wishlist.
// @Description Creates a booked gift in the wishlist based on the provided data.
// @Tags BookedGifts
// @Accept json
// @Produce json
// @Param BookedGiftInWishlist body db.BookedGiftInWishlist true "Booked Gift in Wishlist object to be created"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{data=db.BookedGiftInWishlist}
// @Failure 400 {object} ResponseHTTP{}
// @Router /booked_gifts [post]
func createBookedGiftInWishlist(c *fiber.Ctx) error {
	var bookedGiftInWishlist db.BookedGiftInWishlist
	if err := c.BodyParser(&bookedGiftInWishlist); err != nil {
		return c.SendString(err.Error())
	}

	err := validate.Struct(bookedGiftInWishlist)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).
			SendString(err.Error())
	}

	ok := db.CreateBookedGift(bookedGiftInWishlist)
	if !ok {
		return c.SendString("Error in CreateBookedGift operation")
	}
	return c.SendString("BookedGift is created")
}

// deleteBookedGiftInWishlist godoc
// @Summary Deletes a booked gift from the wishlist.
// @Description Deletes a booked gift from the wishlist based on the provided gift ID and user ID.
// @Tags BookedGifts
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer токен"
// @Param gift_id path string true "ID of the booked gift to be deleted"
// @Success 200 {object} ResponseHTTP{data=db.BookedGiftInWishlist}
// @Failure 400 {object} ResponseHTTP{}
// @Router /booked_gifts/{gift_id} [delete]
func deleteBookedGiftInWishlist(c *fiber.Ctx) error {
	giftID := c.Params("gift_id")

	ok := db.DeleteBookedGift(giftID)
	if !ok {
		return c.SendString("Error in deleteBookedGift operation")
	}
	return c.SendString("bookedGift deleted successfully")
}

// findUserBookedGifts godoc
// @Summary Finds booked gifts for a specific user.
// @Description Finds all booked gifts in the wishlist for a specific user based on the provided user ID.
// @Tags BookedGifts
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer токен"
// @Param user_id path string true "ID of the user"
// @Success 200 {object} ResponseHTTP{data=[]db.BookedGiftInWishlist}
// @Failure 503 {object} ResponseHTTP{}
// @Router /booked_gifts/{user_id} [get]
func findUserBookedGifts(c *fiber.Ctx) error {
	userID := c.Params("user_id")
	gifts, ok := db.FindManyUsersGift(userID)
	if !ok {
		return c.SendString("Error in findUserBookedGifts operation")
	}
	return c.JSON(gifts)
}

// createGiftCategory godoc
// @Summary Creates a new gift category.
// @Description Creates a new gift category based on the provided data.
// @Tags GiftCategory
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer токен"
// @Param GiftCategory body db.GiftCategory true "Gift Category object to be created"
// @Success 200 {object} ResponseHTTP{data=[]db.GiftCategory}
// @Failure 400 {string} string "CategoryName is required"
// @Failure 400 {string} string "Failed to create gift category"
// @Router /gift_category [post]
func createGiftCategory(c *fiber.Ctx) error {
	var giftCategory db.GiftCategory
	if err := c.BodyParser(&giftCategory); err != nil {
		return c.SendString(err.Error())
	}

	err := validate.Struct(giftCategory)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).
			SendString(err.Error())
	}

	giftCategory.ID = "category_" + xid.New().String()

	ok := db.CreateGiftCategory(giftCategory)
	if !ok {
		return c.SendString("Error in createGiftCategory operation")
	}
	return c.JSON(giftCategory)
}

// deleteGiftCategory godoc
// @Summary Deletes a gift category.
// @Description Deletes a gift category based on the provided category ID.
// @Tags GiftCategory
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer токен"
// @Param id path string true "ID of the gift category to be deleted"
// @Success 200 {object} ResponseHTTP{data=db.GiftCategory}
// @Failure 400 {object} ResponseHTTP{}
// @Router /gift_category/{id} [delete]
func deleteGiftCategory(c *fiber.Ctx) error {
	id := c.Params("id")

	ok := db.DeleteGiftCategory(id)
	if !ok {
		return c.SendString("Error in deleteGiftCategory operation")
	}
	return c.SendString("GiftCategory deleted successfully")
}

// GetAllGiftsCategory is a function to get all gifts data from database
// @Summary Get all gift categories
// @Description Get all gift categories
// @Tags GiftCategory
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{data=[]db.Gift}
// @Failure 503 {object} ResponseHTTP{}
// @Router /gift_category [get]
func getManyGiftsCategoryHandler(c *fiber.Ctx) error {
	giftCategories, ok := db.FindManyGiftCategory(db.GiftCategory{})
	// fmt.Println("Gifts:", gifts)
	if !ok {
		return c.SendString("Error in findManyGiftsCategory operation")
	}
	if len(giftCategories) == 0 {
		return c.SendString("No giftCateegories found")
	}
	return c.JSON(giftCategories)
}

// createGiftReviwHandler godoc
// @Summary Create a new review for gift.
// @Description Create a new review for gift.
// @Tags GiftReview
// @Accept  json
// @Produce json
// @Param Gift body db.GiftReview true "Create GiftReview"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{data=db.GiftReview}
// @Failure 400 {object} ResponseHTTP{}
// @Router /gift_review [post]
func createGiftReviwHandler(c *fiber.Ctx) error {
	var giftReview db.GiftReview
	if err := c.BodyParser(&giftReview); err != nil {
		return c.SendString(err.Error())
	}
	err := validate.Struct(giftReview)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).
			SendString(err.Error())
	}

	giftReview.ID = "review_" + xid.New().String()

	ok := db.CreateGiftReview(giftReview)
	if !ok {
		return c.SendString("Error in CreateGiftReview operation")
	}
	return c.SendString("GiftReview is created")
}

// deleteGiftReview godoc
// @Summary Delete a giftReview by ID.
// @Description Deletes a giftReview from the database using the provided ID.
// @Tags GiftReview
// @Accept  json
// @Produce json
// @Param id path string true "GiftReview ID to delete"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {string} string "GiftReview deleted successfully"
// @Failure 400 {string} string "Error in deleteGiftReview operation"
// @Router /gift_review/{id} [delete]
func deleteGiftReviewHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	ok := db.DeleteGiftReview(id)
	if !ok {
		return c.SendString("Error in deleteGiftReview operation")
	}
	return c.SendString("GiftReview deleted successfully")
}

// getGiftReviewByID godoc
// @Summary Get gift review by id
// @Description Get gift review by id
// @Tags GiftReview
// @Accept json
// @Produce json
// @Param id path string true "GiftReview ID"
// @Success 200 {object} ResponseHTTP{data=[]db.GiftReview}
// @Failure 503 {object} ResponseHTTP{}
// @Router /gift_review/review/{id} [get]
func getGiftReviewByIDHandler(c *fiber.Ctx) error {
	reviewID := c.Params("id")
	giftReview, ok := db.GetGiftReviewByID(reviewID)
	if !ok {
		return c.SendString("Error in getGiftReviewByID operation")
	}
	return c.JSON(giftReview)
}

// getGiftReviewsByGiftID godoc
// @Summary Get all gift reviews by giftId
// @Description  Get all gift reviews by giftId
// @Tags GiftReview
// @Accept json
// @Produce json
// @Param gift_id path string true "Gift ID"
// @Success 200 {object} ResponseHTTP{data=[]db.GiftReview}
// @Failure 503 {object} ResponseHTTP{}
// @Router /gift_review/gift/{gift_id} [get]
func getGiftReviewsByGiftIDHandler(c *fiber.Ctx) error {
	giftID := c.Params("gift_id")
	giftReviews, ok := db.GetGiftReviewsByGiftID(giftID)
	if !ok {
		return c.SendString("Error in getGiftReviewsByGiftID operation")
	}
	return c.JSON(giftReviews)
}

// calculateAverageMarkByGiftIDHandler godoc
// @Summary Calculate average mark for a gift by its ID
// @Description Calculate average mark for a gift by its ID
// @Tags GiftReview
// @Accept json
// @Produce json
// @Param gift_id path string true "Gift ID"
// @Success 200 {object} float32 "Average mark"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /gift_review/mark/{gift_id} [get]
func calculateAverageMarkByGiftIDHandler(c *fiber.Ctx) error {
	giftID := c.Params("gift_id")
	averageMark, ok := db.CalculateAverageMarkByGiftID(giftID)
	if !ok {
		return c.SendString("calculateAverageMarkByGiftID operation")
	}
	return c.JSON(averageMark)
}

// createService godoc
// @Summary Creates a new service.
// @Tags Services
// @Accept json
// @Produce json
// @Param Service body db.Service true "Create Service"
// @Param seller_Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{data=db.Service}
// @Failure 400 {object} ResponseHTTP{}
// @Router /services [post]
func createServiceHandler(c *fiber.Ctx) error {
	var service db.Service
	if err := c.BodyParser(&service); err != nil {
		return c.SendString(err.Error())
	}

	err := validate.Struct(service)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).
			SendString(err.Error())
	}
	// Костыль. Кастомную валидацию для Decimal не прописать :(
	if !service.Price.IsPositive() {
		return c.Status(fiber.StatusUnprocessableEntity).
			SendString("Only positive deicmals are allowed!")
	}

	service.ServiceID = "service_" + xid.New().String()

	ok := db.CreateService(service)
	if !ok {
		return c.SendString("Error in createService operation")
	}

	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success register a service.",
		Data:    &service,
	})
}

// getManyServices godoc
// @Summary Fetches all services.
// @Tags Services
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{data=db.Service}
// @Failure 400 {object} ResponseHTTP{}
// @Router /services [get]
func getManyServicesHandler(c *fiber.Ctx) error {
	result, ok := db.FindManyService()
	if !ok {
		return c.SendString("Error in findManyServices operation")
	}
	return c.JSON(result)
}

// getSingleService godoc
// @Summary Fetches all services of a specified seller.
// @Tags Services
// @Accept json
// @Produce json
// @Param id path string true "Seller ID"
// @Success 200 {object} ResponseHTTP{data=db.Service}
// @Failure 400 {object} ResponseHTTP{}
// @Router /services/seller/{seller_id} [get]
func getSingleServiceHandler(c *fiber.Ctx) error {
	sellerId := c.Params("seller_id")
	result, ok := db.FindSingleService(sellerId)
	if !ok {
		return c.SendString("Error in findOneService operation")
	}
	return c.JSON(result)
}

// getOneService godoc
// @Summary Fetches a specific service.
// @Tags Services
// @Accept json
// @Produce json
// @Param service_id path string true "Service ID"
// @Success 200 {object} ResponseHTTP{data=db.Service}
// @Failure 400 {object} ResponseHTTP{}
// @Router /services/{service_id} [get]
func getOneServiceHandler(c *fiber.Ctx) error {
	serviceId := c.Params("service_id")
	result, ok := db.FindOneService(serviceId)
	if !ok {
		return c.SendString("Error in findOneService operation")
	}
	return c.JSON(result)
}

// updateService godoc
// @Summary Updates an existing service.
// @Tags Services
// @Accept json
// @Produce json
// @Param id path string true "Service ID"
// @Param Service body db.Service true "Update Service"
// @Param seller_Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{data=db.Service}
// @Failure 400 {object} ResponseHTTP{}
// @Router /services/{id} [patch]
func updateServiceHandler(c *fiber.Ctx) error {
	serviceId := c.Params("id")

	var updatedService db.Service
	if err := c.BodyParser(&updatedService); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error parsing request body")
	}

	if !updatedService.Price.IsPositive() {
		return c.Status(fiber.StatusUnprocessableEntity).
			SendString("Only positive deicmals are allowed!")
	}

	ok := db.UpdateService(serviceId, updatedService)
	if !ok {
		return c.SendString("Error in updateService operation")
	}
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Successfully updated the Service.",
		Data:    &updatedService,
	})
}

// deleteService godoc
// @Summary Deletes a specified service.
// @Tags Services
// @Accept json
// @Produce json
// @Param service_id path string true "Delete Service"
// @Param seller_Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{data=db.Service}
// @Failure 400 {object} ResponseHTTP{}
// @Router /services/{service_id} [delete]
func deleteServiceHandler(c *fiber.Ctx) error {
	id := c.Params("service_id")

	ok := db.DeleteService(id)
	if !ok {
		return c.SendString("Error in deleteService operation")
	}
	return c.SendString("Service deleted successfully")
}

// createSellerToService godoc
// @Summary Creates a new connection of Seller-Service.
// @Tags SellerToService
// @Accept json
// @Produce json
// @Param SellerToService body db.SellerToService true "Create Sellers-Services"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{data=db.SellerToService}
// @Failure 400 {object} ResponseHTTP{}
// @Router /sellerToService [post]
func createSellerToServiceHandler(c *fiber.Ctx) error {
	var sellerToService db.SellerToService
	if err := c.BodyParser(&sellerToService); err != nil {
		return c.SendString(err.Error())
	}

	err := validate.Struct(sellerToService)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}

	ok := db.CreateSellerToService(sellerToService)
	if !ok {
		return c.SendString("Error in createsellerToService operation")
	}

	return c.JSON(sellerToService)
}

// getManySellerToService godoc
// @Summary Fetches all Seller-Service connections.
// @Tags SellerToService
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{data=db.SellerToService}
// @Failure 400 {object} ResponseHTTP{}
// @Router /sellerToService [get]
func getManySellerToServiceHandler(c *fiber.Ctx) error {
	result, ok := db.FindManySellerToService()
	if !ok {
		return c.SendString("Error in findManySellerToService operation")
	}
	return c.JSON(result)
}

// getOneSellerToService godoc
// @Summary Fetches all Services that belong to the speicfied Seller.
// @Tags SellerToService
// @Accept json
// @Produce json
// @Param id path string true "Seller ID"
// @Success 200 {object} ResponseHTTP{data=db.SellerToService}
// @Failure 400 {object} ResponseHTTP{}
// @Router /sellerToService/{id} [get]
func getOneSellerToServiceHandler(c *fiber.Ctx) error {
	sellerId := c.Params("id")
	result, ok := db.FindOneSellerToService(sellerId)
	if !ok {
		return c.SendString("Error in findOneSellersService operation")
	}
	return c.JSON(result)
}

// deleteSellersService godoc
// @Summary Deletes a specified Seller-Service connection.
// @Tags SellerToService
// @Accept json
// @Produce json
// @Param id path string true "Service ID"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{data=db.SellerToService}
// @Failure 400 {object} ResponseHTTP{}
// @Router /sellerToService/{id} [delete]
func deleteSellerToServiceHandler(c *fiber.Ctx) error {
	serviceId := c.Params("id")

	ok := db.DeleteSellerToService(serviceId)
	if !ok {
		return c.SendString("Error in deleteSellerToService operation")
	}
	return c.SendString("SellerService deleted successfully")
}

// createServiceReview godoc
// @Summary Creates a new serviceReview.
// @Tags ServiceReviews
// @Accept json
// @Produce json
// @Param ServiceReview body db.ServiceReview true "Create ServiceReview"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{data=db.ServiceReview}
// @Failure 400 {object} ResponseHTTP{}
// @Router /serviceReviews [post]
func createServiceReviewHandler(c *fiber.Ctx) error {
	var serviceReview db.ServiceReview
	if err := c.BodyParser(&serviceReview); err != nil {
		return c.SendString(err.Error())
	}

	err := validate.Struct(serviceReview)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).
			SendString(err.Error())
	}
	if serviceReview.Mark.IsNegative() ||
		serviceReview.Mark.GreaterThan(decimal.NewFromInt(5)) {
		return c.Status(fiber.StatusUnprocessableEntity).
			SendString("Only positive marks less or equal to 5 are allowed!")
	}

	serviceReview.ID = "serviceReview_" + xid.New().String()
	serviceReview.CreateDate = time.Now()
	serviceReview.UpdateDate = time.Now()

	ok := db.CreateServiceReview(serviceReview)
	if !ok {
		return c.SendString("Error in createServiceReview operation")
	}

	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success register a serviceReview.",
		Data:    &serviceReview,
	})
}

// getManyServiceReviews godoc
// @Summary Fetches all serviceReviews.
// @Tags ServiceReviews
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{data=db.ServiceReview}
// @Failure 400 {object} ResponseHTTP{}
// @Router /serviceReviews [get]
func getManyServiceReviewsHandler(c *fiber.Ctx) error {
	result, ok := db.FindManyServiceReview()
	if !ok {
		return c.SendString("Error in findManyServiceReviews operation")
	}
	return c.JSON(result)
}

// getOneServiceReview godoc
// @Summary Fetches a specific serviceReview.
// @Tags ServiceReviews
// @Accept json
// @Produce json
// @Param id path string true "ServiceReview ID"
// @Success 200 {object} ResponseHTTP{data=db.ServiceReview}
// @Failure 400 {object} ResponseHTTP{}
// @Router /serviceReviews/{id} [get]
func getOneServiceReviewHandler(c *fiber.Ctx) error {
	serviceReviewId := c.Params("id")
	result, ok := db.FindOneServiceReview(serviceReviewId)
	if !ok {
		return c.SendString("Error in findOneServiceReview operation")
	}
	return c.JSON(result)
}

// getSingleServiceReview godoc
// @Summary Fetches all Service Reviews for a specified Service.
// @Tags ServiceReviews
// @Accept json
// @Produce json
// @Param service_id path string true "Service ID"
// @Success 200 {object} ResponseHTTP{data=db.ServiceReview}
// @Failure 400 {object} ResponseHTTP{}
// @Router /serviceReviews/service/{service_id} [get]
func getSingleServiceReviewHandler(c *fiber.Ctx) error {
	serviceId := c.Params("service_id")
	result, ok := db.FindSingleServiceReview(serviceId)
	if !ok {
		return c.SendString("Error in findSingleServiceReview operation")
	}
	return c.JSON(result)
}

// updateServiceReview godoc
// @Summary Updates an existing serviceReview.
// @Tags ServiceReviews
// @Accept json
// @Produce json
// @Param id path string true "ServiceReview ID"
// @Param ServiceReview body db.ServiceReview true "Update ServiceReview"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{data=db.ServiceReview}
// @Failure 400 {object} ResponseHTTP{}
// @Router /serviceReviews/{id} [patch]
func updateServiceReviewHandler(c *fiber.Ctx) error {
	serviceReviewId := c.Params("id")

	var updatedServiceReview db.ServiceReview
	if err := c.BodyParser(&updatedServiceReview); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error parsing request body")
	}

	if updatedServiceReview.Mark.IsNegative() ||
		updatedServiceReview.Mark.GreaterThan(decimal.NewFromInt(5)) {
		return c.Status(fiber.StatusUnprocessableEntity).
			SendString("Only positive marks less or equal to 5 are allowed!")
	}
	updatedServiceReview.UpdateDate = time.Now()

	ok := db.UpdateServiceReview(serviceReviewId, updatedServiceReview)
	if !ok {
		return c.SendString("Error in updateServiceReview operation")
	}
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Successfully updated the ServiceReview",
		Data:    &updatedServiceReview,
	})
}

// Step by step guide to authenticate a user:
// 1. Create a middleware (attach to group of routes)
// - Take session/jwt token from Authorization header
// - Get user from session/jwt token (error if not found)
// - Add user to context (c.Locals)
// 2. Use middleware in routes
// - Take user from context (c.Locals)
// - Use user in handler
// 3. Create a register & login handlers
// - Register: add user to db, return session/jwt token
// - Login: check user in db, return session/jwt token
func superSecretHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(string)
	return c.SendString("This is a super secret route. Hi " + user + "!")
}

// Register Handler godoc
// @Summary Creates a new gift.
// @Description get the status of server.
// @Tags auth
// @Accept  json
// @Produce json
// @Param User body db.User true "Reg user"
// @Success 200 {object} ResponseHTTP{data=db.User}
// @Failure 400 {object} ResponseHTTP{}
// @Router /register [post]
func registerHandler(c *fiber.Ctx) error {
	var user db.User
	if err := c.BodyParser(&user); err != nil {
		return c.SendString(err.Error())
	}

	err := validate.Struct(user)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).
			SendString(err.Error())
	}

	existingUser := new(db.User)

	if err := db.Database.Model(user).Where("Login = ?", user.Login).First(existingUser).Error; err == nil {
		// Если пользователь с таким Login уже существует
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Пользователь с таким Login уже существует"})
	}

	user.ID = ""
	user.ID = "user_" + xid.New().String()

	ok := db.CreateUser(user)
	if !ok {
		return c.SendString("Error in CreateUser operation")
	}

	session := db.Session{
		ID:     "session_" + xid.New().String(),
		UserID: user.ID,
	}
	ok = db.CreateSession(session)
	if !ok {
		return c.SendString("Cannot create session")
	}

	return c.JSON(session)

}

// Login Handler godoc
// @Summary Creates a new gift.
// @Description get the status of server.
// @Tags auth
// @Accept  json
// @Produce json
// @Param User body db.User true "Reg user"
// @Success 200 {object} ResponseHTTP{data=db.User}
// @Failure 400 {object} ResponseHTTP{}
// @Router /login [post]
func loginHandler(c *fiber.Ctx) error {
	var authCredentials AuthCredentials
	if err := c.BodyParser(&authCredentials); err != nil {
		return c.SendString(err.Error())
	}

	err := validate.Struct(authCredentials)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).
			SendString(err.Error())
	}

	user, ok := db.FindUser(authCredentials.Login, authCredentials.Password)
	if !ok {
		return c.SendString("Invalid creditials")
	}
	session := db.Session{
		ID:     "session_" + xid.New().String(),
		UserID: user.ID,
	}
	ok = db.CreateSession(session)
	if !ok {
		return c.SendString("Cannot create session")
	}

	return c.JSON(session)

	// session, ok := getUser(authCredentials.Login, authCredentials.Password)
	// if !ok {
	// 	return c.SendString("Invalid credentials")
	// }

	// return c.JSON(AuthResponse{Session: session})
}

// RegisterSeller Handler godoc
// @Summary Creates a new seller.
// @Tags auth
// @Accept  json
// @Produce json
// @Param Seller body db.Seller true "Register seller"
// @Success 200 {object} ResponseHTTP{data=db.Seller}
// @Failure 400 {object} ResponseHTTP{}
// @Router /registerSeller [post]
func registerSellerHandler(c *fiber.Ctx) error {
	var seller db.Seller
	if err := c.BodyParser(&seller); err != nil {
		return c.SendString(err.Error())
	}

	err := validate.Struct(seller)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).
			SendString(err.Error())
	}

	seller.SellerID = ""
	seller.SellerID = "seller_" + xid.New().String()
	ok := db.CreateSeller(seller)
	if !ok {
		return c.SendString("Error in CreateSeller operation")
	}

	return c.SendString("Registered successfully!")
}

// LoginSeller Handler godoc
// @Summary Logs a Seller in.
// @Tags auth
// @Accept  json
// @Produce json
// @Param Seller body db.Seller true "Reg seller"
// @Success 200 {object} ResponseHTTP{data=db.Seller}
// @Failure 400 {object} ResponseHTTP{}
// @Router /loginSeller [post]
func loginSellerHandler(c *fiber.Ctx) error {
	var authCredentials SellerAuthCredentials
	if err := c.BodyParser(&authCredentials); err != nil {
		return c.SendString(err.Error())
	}

	err := validate.Struct(authCredentials)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).
			SendString(err.Error())
	}

	seller, ok := db.FindSeller(authCredentials.Login, authCredentials.Password)
	if !ok {
		return c.SendString("Invalid creditials")
	}
	sellerSession := db.SellerSession{
		ID:       "session_" + xid.New().String(),
		SellerID: seller.SellerID,
	}
	ok = db.CreateSellerSession(sellerSession)
	if !ok {
		return c.SendString("Cannot create session")
	}

	return c.JSON(sellerSession)
}

// AddWishHandler godoc
// @Summary Adds a gift in your wishlist.
// @Description There you can add wish in your wishlist.
// @Tags Wishes
// @Accept  json
// @Produce json
// @Param Wishes body db.Wishes true "Add wishes"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{data=db.Wishes}
// @Failure 400 {object} ResponseHTTP{}
// @Router /wishes/{wishlist_id} [post]
func AddWishHandler(c *fiber.Ctx) error {
	var wish db.Wishes
	gift_id := c.Params("gift_id")
	wishlist_id := c.Params("wishlist_id")

	if err := c.BodyParser(&wish); err != nil {
		return c.SendString(err.Error())
	}

	// err := validate.Struct(wish)
	// if err != nil {
	// 	return c.Status(fiber.StatusUnprocessableEntity).
	// 		SendString(err.Error())
	// }

	ok := db.AddWish(wishlist_id, gift_id)
	if !ok {
		return c.SendString("Error in Create operation")
	}
	return c.SendString("Create Wish succesfully")
}

// DeleteWishHandler godoc
// @Summary Deletes a wish from wishlist.
// @Description
// @Tags Wishes
// @Accept  json
// @Produce json
// @Param Wishes body db.Wishes true "Add wishes"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{data=db.Wishes}
// @Failure 400 {object} ResponseHTTP{}
// @Router /wishes/{id}/{wishlist_id} [delete]
func DeleteWishHandler(c *fiber.Ctx) error {
	wishlistID := c.Params("wishlist_id")
	giftID := c.Params("gift_id")
	fmt.Println("giftID", giftID)
	ok := db.DeleteWish(wishlistID, giftID)
	if !ok {
		return c.SendString("Error in Delete wish operation")
	}
	return c.SendString("Create Wish succesfully")
}

// createWishlist godoc
// @Summary Creates a brand new wishlist.
// @Description There is handler that creates a new wishlist.
// @Tags Wishlist
// @Accept  json
// @Produce json
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{data=db.UserWishlist}
// @Failure 400 {object} ResponseHTTP{}
// @Router /wishlists [post]
func CreateWishlistHandler(c *fiber.Ctx) error {
	var wishlist db.UserWishlist
	if err := c.BodyParser(&wishlist); err != nil {
		return c.SendString(err.Error())
	}
	wishlist.ID = "wishlist_" + xid.New().String()
	wishlist.UserID = c.Locals("user").(string)

	err := validate.Struct(wishlist)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).
			SendString(err.Error())
	}

	ok := db.CreateWishlist(wishlist)
	if !ok {
		return c.SendString("Error in Create wishlist operation")
	}

	return c.SendString("Create Wishlist succesfully")

}

// FindManyWishlists godoc
// @Summary Finds you all of your wishlists.
// @Tags Wishlist
// @Accept  json
// @Produce json
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{data=[]db.UserWishlist}
// @Failure 400 {object} ResponseHTTP{}
// @Router /wishlists [get]
func FindManyWishlistsHandler(c *fiber.Ctx) error {
	userID := c.Locals("user").(string)
	wishlists, ok := db.FindManyWishlists(userID)
	if !ok {
		return c.SendString("Error in FindManyWishlists operation")
	}
	return c.JSON(wishlists)
}

// findWishlistByNameHandler godoc
// @Summary Finds wishlist by name
// @Tags Wishlist
// @Accept  json
// @Produce json
// @Param name path string true "Delete ServiceReview"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{data=db.UserWishlist}
// @Failure 400 {object} ResponseHTTP{}
// @Router /wishlists/{name} [get]
func findWishlistByNameHandler(c *fiber.Ctx) error {
	nameID := c.Params("name")
	wishlist, ok := db.FindWishlistByName(nameID)
	if !ok {
		return c.SendString("error in findManyWishlists operation")
	}
	return c.JSON(wishlist)
}

// FindAllWishesInWishlistHandler godoc
// @Summary finds all wishes in your wishlist
// @Description get the status of server.
// @Tags Wishes
// @Accept  json
// @Produce json
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{data=db.Wishes}
// @Failure 400 {object} ResponseHTTP{}
// @Router /wishes/{wishlist_id} [get]
func FindAllWishesInWishlistHandler(c *fiber.Ctx) error {
	wishlistID := c.Params("wishlist_id")
	wishes, ok := db.GetManyWishesInWishlist(wishlistID)
	if !ok {
		return c.SendString("Invalid credentials")
	}

	return c.JSON(wishes)
}

// deleteServiceReview godoc
// @Summary Deletes a specified serviceReview.
// @Tags ServiceReviews
// @Accept json
// @Produce json
// @Param id path string true "Delete ServiceReview"
// @Success 200 {object} ResponseHTTP{data=db.ServiceReview}
// @Param Authorization header string true "Bearer токен"
// @Failure 400 {object} ResponseHTTP{}
// @Router /serviceReviews/{id} [delete]
func deleteServiceReviewHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	ok := db.DeleteServiceReview(id)
	if !ok {
		return c.SendString("Error in deleteServiceReview operation")
	}
	return c.SendString("ServiceReview deleted successfully")
}

// updateWishlist godoc
// @Summary updates your wishlist.
// @Tags Wishlist
// @Accept  json
// @Produce json
// @Param UserWishlist body db.UserWishlist true "Create Wishlist"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{data=db.UserWishlist}
// @Failure 400 {object} ResponseHTTP{}
// @Router /wishlists [put]
func UpdateWishlistHandler(c *fiber.Ctx) error {
	wishlistID := c.Params("id")
	fmt.Println("wishlist id", wishlistID)
	var wishlistName struct {
		Name string `json:"name"`
	}
	if err := c.BodyParser(&wishlistName); err != nil {
		return err
	}

	ok := db.UpdateWishlist(wishlistID, wishlistName.Name)
	if !ok {
		return c.SendString("Error in UpdateWishlist Operation")
	}
	return c.SendString("Update wishlist succesfully")
}

// deleteWishlist godoc
// @Summary Deletes your wishlist.
// @Description delete a wishlist and wishes related with it
// @Tags Wishlist
// @Accept  json
// @Produce json
// @Param Wishlist body db.UserWishlist true "Delete Wishlist"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{data=db.UserWishlist}
// @Failure 400 {object} ResponseHTTP{}
// @Router /wishlists/{wishlist_id}/{gift_id} [delete]
func DeleteWishlistHandler(c *fiber.Ctx) error {
	wishlistID := c.Params("id")
	giftID := c.Params("gift_id")
	userID := c.Locals("user").(string)
	ok := db.DeleteWishlist(wishlistID, giftID, userID)
	if !ok {
		return c.SendString("Error in DeleteGift Operation")
	}
	return c.SendString("Delete wishlist succesfully")
}

//Quest

// createQuestHandler обрабатывает HTTP POST запросы на /quest
// @Summary Создает новый Quest
// @Description Принимает JSON тело запроса с полями Quest и создает новый Quest
// @Tags Quest
// @Accept json
// @Produce json
// @Param Quest body db.Quest true "Create Quest"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{data=db.Quest}
// @Failure 400 {object} ResponseHTTP{}
// @Router /quest [post]
func createQuestHandler(c *fiber.Ctx) error {
	var quest db.Quest
	if err := c.BodyParser(&quest); err != nil {
		return c.SendString(err.Error())
	}

	err := validate.Struct(quest)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).
			SendString(err.Error())
	}

	quest.ID = "quest_" + xid.New().String()

	quest.UserID = c.Locals("user").(string)

	ok := db.CreateQuest(quest)
	if !ok {
		return c.SendString("Error in createQuest operation")
	}

	return c.JSON(quest)
}

// updateQuestHandler обрабатывает HTTP PUT запросы на /quest/update.
// @Summary Обновляет существующий Quest
// @Description Принимает JSON тело запроса с обновленными полями Quest и обновляет существующий Quest
// @Tags Quest
// @Accept json
// @Produce json
// @Param Quest body db.Quest true "Update Quest"
// @Success 200 {object} ResponseHTTP{data=db.Quest}
// @Failure 400 {object} ResponseHTTP{}
// @Router /quest/update [put]
func updateQuestHandler(c *fiber.Ctx) error {
	questId := c.Params("id")
	var quest db.Quest
	if err := c.BodyParser(&quest); err != nil {
		return c.SendString(err.Error())
	}
	ok := db.UpdateQuest(questId, quest)
	if !ok {
		return c.SendString("Error in updateQuest operation")
	}
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Quest updated Succesfully",
		Data:    &quest,
	})
}

// getOneQuestHandler обрабатывает HTTP GET запросы на /quest/getone/{id}.
// @Summary Получает один квест Quest по ID
// @Description Возвращает информацию о конкретном квесте Quest по его ID
// @Tags Quest
// @Param id path int true "Quest ID"
// @Produce json
// @Success 200 {object} ResponseHTTP{data=db.Quest}
// @Failure 404 {string} string "Quest not found"
// @Router /quest/getone/{id} [get]
func getOneQuestHandler(c *fiber.Ctx) error {
	questId := c.Params("id")
	reuslt, ok := db.FindOneQuest(questId)
	if !ok {
		return c.SendString("Error in findOneQuest operation")
	}
	return c.JSON(reuslt)
}

// getManyQuestHandler обрабатывает HTTP GET запросы на /quest/getmany.
// @Summary Получает список квестов Quest
// @Description Возвращает список всех квестов Quest
// @Tags Quest
// @Produce json
// @Success 200 {object} ResponseHTTP{data=db.Quest}
// @Router /quest [get]
func getManyQuestHandler(c *fiber.Ctx) error {
	result, ok := db.FindManyQuest()
	if !ok {
		return c.SendString("Error in findManyQuest operation")
	}
	return c.JSON(result)
}

// deleteQuestHandler обрабатывает HTTP DELETE запросы на /quest/{id}
// @Summary Удаляет существующий Quest по ID
// @Description Принимает ID квеста в URL и удаляет соответствующий квест
// @Tags Quest
// @Param id path int true "Quest ID"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {string} string "Quest deleted successfully"
// @Failure 404 {string} string "Quest not found"
// @Router /quest/{id} [delete]
func deleteQuestHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	ok := db.DeleteQuest(id)
	if !ok {
		return c.SendString("Error in deleteQuest operation")
	}
	return c.SendString("Quest deleted successfully")
}

//Subquest

// createSubquestHandler обрабатывает HTTP POST запросы на /subquest
// @Summary Создает новый Subquest
// @Description Принимает JSON тело запроса с полями Subquest и создает новый Subquest
// @Tags Subquest
// @Accept json
// @Produce json
// @Param Subquest body db.Subquest true "Create Subquest"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{data=db.Subquest}
// @Failure 400 {object} ResponseHTTP{}
// @Router /subquest [post]
func createSubquestHandler(c *fiber.Ctx) error {
	var subquest db.Subquest
	if err := c.BodyParser(&subquest); err != nil {
		return c.SendString(err.Error())
	}

	err := validate.Struct(subquest)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).
			SendString(err.Error())
	}

	subquest.ID = "subquest_" + xid.New().String()

	ok := db.CreateSubquest(subquest)
	if !ok {
		return c.SendString("Error in createSubquest operation")
	}

	return c.JSON(subquest)
}

// getManySubquestHandler обрабатывает HTTP GET запросы на /subquest
// @Summary Получает список Subquest
// @Description Возвращает список всех подзаданий (Subquest)
// @Tags Subquest
// @Produce json
// @Success 200 {object} ResponseHTTP{data=db.Subquest}
// @Router /subquest [get]
func getManySubquestHandler(c *fiber.Ctx) error {
	result, ok := db.FindManySubquest()
	if !ok {
		return c.SendString("Error in findManySubquest operation")
	}
	return c.JSON(result)
}

// getOneSubquestHandler обрабатывает HTTP GET запросы на /subquest/{id}
// @Summary Получает одно Subquest по ID
// @Description Возвращает информацию о конкретном подзадании (Subquest) по его ID
// @Tags Subquest
// @Param id path int true "Subquest ID"
// @Produce json
// @Success 200 {object} ResponseHTTP{data=db.Subquest}
// @Failure 404 {string} string "Subquest not found"
// @Router /subquest/{id} [get]
func getOneSubquestHandler(c *fiber.Ctx) error {
	subquestId := c.Params("id")
	result, ok := db.FindOneSubquest(subquestId)
	if !ok {
		return c.SendString("Error in findOneSubquest operation")
	}
	return c.JSON(result)
}

// deleteSubquestHandler обрабатывает HTTP DELETE запросы на /subquest/{id}
// @Summary Удаляет существующий Subquest по ID
// @Description Принимает ID подзадания в URL и удаляет соответствующее подзадание
// @Tags Subquest
// @Param id path int true "Subquest ID"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {string} string "Subquest deleted successfully"
// @Failure 404 {string} string "Subquest not found"
// @Router /subquest/{id} [delete]
func deleteSubquestHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	ok := db.DeleteSubquest(id)
	if !ok {
		return c.SendString("Error in deleteSubquest operation")
	}
	return c.SendString("Subquest deleted successfully")
}

// updateSubquestHandler обрабатывает HTTP PUT запросы на /subquest/{id}
// @Summary Обновляет существующий Subquest
// @Description Принимает JSON тело запроса с обновленными полями Subquest и обновляет существующий Subquest
// @Tags Subquest
// @Accept json
// @Produce json
// @Param Subquest body db.Subquest true "Update Subquest"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{data=db.Subquest}
// @Failure 400 {object} ResponseHTTP{}
// @Router /quest/{id} [put]
func updateSubquestHandler(c *fiber.Ctx) error {
	subquestId := c.Params("id")
	var subquest db.Subquest
	if err := c.BodyParser(&subquest); err != nil {
		return c.SendString(err.Error())
	}
	ok := db.UpdateSubquest(subquestId, subquest)
	if !ok {
		return c.SendString("Error in updateSubquest operation")
	}
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Subquest updated Succesfully",
		Data:    &subquest,
	})
}

//Tasks

// createTasksHandler обрабатывает HTTP POST запросы на /tasks
// @Summary Создает новое задание Tasks
// @Description Принимает JSON тело запроса с полями Tasks и создает новое задание
// @Tags Tasks
// @Accept json
// @Produce json
// @Param Tasks body db.Tasks true "Create Tasks"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{data=db.Tasks}
// @Failure 400 {object} ResponseHTTP{}
// @Router /tasks [post]
func createTasksHandler(c *fiber.Ctx) error {
	var tasks db.Tasks
	if err := c.BodyParser(&tasks); err != nil {
		return c.SendString(err.Error())
	}

	err := validate.Struct(tasks)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).
			SendString(err.Error())
	}

	tasks.ID = "tasks_" + xid.New().String()

	ok := db.CreateTasks(tasks)
	if !ok {
		return c.SendString("Error in createTasks operation")
	}

	return c.JSON(tasks)
}

// updateTasksHandler обрабатывает HTTP PUT запросы на /tasks/{id}
// @Summary Обновляет существующее задание Tasks
// @Description Принимает JSON тело запроса с обновленными полями Tasks и обновляет существующее задание
// @Tags Tasks
// @Accept json
// @Produce json
// @Param Tasks body db.Tasks true "Update Tasks"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{data=db.Tasks}
// @Failure 400 {object} ResponseHTTP{}
// @Router /tasks/{id} [put]
func updateTasksHandler(c *fiber.Ctx) error {
	tasksId := c.Params("id")
	var tasks db.Tasks
	if err := c.BodyParser(&tasks); err != nil {
		return c.SendString(err.Error())
	}
	ok := db.UpdateTasks(tasksId, tasks)
	if !ok {
		return c.SendString("Error in updateTasks operation")
	}
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Tasks updated Succesfully",
		Data:    &tasks,
	})

}

// getOneTasksHandler обрабатывает HTTP GET запросы на /tasks/{id}
// @Summary Получает одно задание Tasks по ID
// @Description Возвращает информацию о конкретном задании Tasks по его ID
// @Tags Tasks
// @Param id path int true "Tasks ID"
// @Produce json
// @Success 200 {object} ResponseHTTP{data=db.Tasks}
// @Failure 404 {string} string "Tasks not found"
// @Router /tasks/{id} [get]
func getOneTasksHandler(c *fiber.Ctx) error {
	taskId := c.Params("id")
	result, ok := db.FindOneTasks(taskId)
	if !ok {
		return c.SendString("Error in findOneTasks operation")
	}
	return c.JSON(result)
}

// getManyTasksHandler обрабатывает HTTP GET запросы на /tasks
// @Summary Получает список заданий Tasks
// @Description Возвращает список все заданий Tasks
// @Tags Tasks
// @Produce json
// @Success 200 {object} ResponseHTTP{data=db.Tasks}
// @Router /tasks [get]
func getManyTasksHandler(c *fiber.Ctx) error {
	result, ok := db.FindManyTasks()
	if !ok {
		return c.SendString("Error in findManyTasks operation")
	}
	return c.JSON(result)
}

// deleteTasksHandler обрабатывает HTTP DELETE запросы на /tasks/{id}
// @Summary Удаляет существующее задание Tasks по ID
// @Description Принимает ID задания в URL и удаляет соответствующее задание
// @Tags Tasks
// @Param id path int true "Tasks ID"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {string} string "Tasks deleted successfully"
// @Failure 404 {string} string "Tasks not found"
// @Router /tasks/{id} [delete]
func deleteTasksHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	ok := db.DeleteTasks(id)
	if !ok {
		return c.SendString("Error in deleteTasks operation")
	}
	return c.SendString("Tasks deleted successfully")
}

//OfflineShops

// createOfflineShopHandler обрабатывает HTTP POST запросы на /offlineshop
// @Summary Создает новый Offline Shop
// @Description Принимает JSON тело запроса с полями Offline Shop и создает новый Offline Shop
// @Tags OfflineShops
// @Accept json
// @Produce json
// @Param OfflineShop body db.OfflineShops true "Create Offline Shop"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{data=db.OfflineShops}
// @Failure 400 {object} ResponseHTTP{}
// @Router /offlineshop [post]
func createOfflineShopsHandler(c *fiber.Ctx) error {
	var offlineshops db.OfflineShops
	if err := c.BodyParser(&offlineshops); err != nil {
		return c.SendString(err.Error())
	}

	err := validate.Struct(offlineshops)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).
			SendString(err.Error())
	}

	offlineshops.ID = "offlineshops_" + xid.New().String()

	ok := db.CreateOfflineShops(offlineshops)
	if !ok {
		return c.SendString("Error in createOfflineShops operation")
	}

	return c.JSON(offlineshops)
}

// updateOfflineShopHandler обрабатывает HTTP PUT запросы на /offlineshop/{id}
// @Summary Обновляет существующий Offline Shop по ID
// @Description Принимает JSON тело запроса с обновленными полями Offline Shop и обновляет существующий Offline Shop по его ID
// @Tags OfflineShops
// @Accept json
// @Produce json
// @Param id path string true "Offline Shop ID"
// @Param OfflineShop body db.OfflineShops true "Update Offline Shop"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{data=db.OfflineShops}
// @Failure 400 {object} ResponseHTTP{}
// @Router /offlineshop/{id} [put]
func updateOfflineShopsHandler(c *fiber.Ctx) error {
	offlineshopsId := c.Params("id")
	var offlineshops db.OfflineShops
	if err := c.BodyParser(&offlineshops); err != nil {
		return c.SendString(err.Error())
	}
	ok := db.UpdateOfflineShops(offlineshopsId, offlineshops)
	if !ok {
		return c.SendString("Error in updateOfflineShops operation")
	}
	return c.SendString("OfflineShops updated Succesfully")
}

// getOneOfflineShopsHandler обрабатывает HTTP GET запросы на /offlineshops/{id}
// @Summary Получает один офлайн магазин OfflineShops по ID
// @Description Возвращает информацию о конкретном офлайн магазине OfflineShops по его ID
// @Tags OfflineShops
// @Param id path int true "OfflineShops ID"
// @Produce json
// @Success 200 {object} ResponseHTTP{data=db.OfflineShops}
// @Failure 404 {string} string "OfflineShops not found"
// @Router /offlineshops/{id} [get]
func getOneOfflineShopsHandler(c *fiber.Ctx) error {
	offlineshopsId := c.Params("id")
	result, ok := db.FindOneOfflineShops(offlineshopsId)
	if !ok {
		return c.SendString("Error in findOneOfflineShops operation")
	}
	return c.JSON(result)
}

// getManyOfflineShopsHandler обрабатывает HTTP GET запросы на /offlineshops
// @Summary Получает список офлайн магазинов OfflineShops
// @Description Возвращает список всех офлайн магазинов OfflineShops
// @Tags OfflineShops
// @Produce json
// @Success 200 {object} ResponseHTTP{data=db.OfflineShops}
// @Router /offlineshops [get]
func getManyOfflineShopsHandler(c *fiber.Ctx) error {
	result, ok := db.FindManyOfflineShops()
	if !ok {
		return c.SendString("Error in findManyOfflineShops operation")
	}
	return c.JSON(result)
}

// deleteOfflineShopHandler обрабатывает HTTP DELETE запросы на /offlineshop/{id}
// @Summary Удаляет существующий Offline Shop по ID
// @Description Принимает ID офлайн магазина в URL и удаляет соответствующий офлайн магазин
// @Tags OfflineShops
// @Param id path string true "Offline Shop ID"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {string} string "Offline Shop deleted successfully"
// @Failure 404 {string} string "Offline Shop not found"
// @Router /offlineshop/{id} [delete]
func deleteOfflineShopsHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	ok := db.DeleteOfflineShops(id)
	if !ok {
		return c.SendString("Error in deleteOfflineShops operation")
	}
	return c.SendString("OfflineShops deleted successfully")
}

// createSelectionHandler обрабатывает HTTP POST запросы на /selection/create.
// @Summary Создает новый Selection
// @Description Принимает JSON тело запроса с полями Selection и создает новый Selection
// @Tags Selection
// @Accept json
// @Produce json
// @Param Selection body db.Selection true "Create Selection"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{data=db.Selection}
// @Failure 400 {object} ResponseHTTP{}
// @Router /selection [post]
func createSelectionHandler(c *fiber.Ctx) error {
	var selection db.Selection
	if err := c.BodyParser(&selection); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if selection.Name == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Name is required")
	}
	if selection.Description == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Description is required")
	}

	selection.ID = "selection_" + xid.New().String()
	if selection.ID == "" {
		return c.Status(fiber.StatusInternalServerError).SendString("Error generating ID")
	}

	selection.UserID = c.Locals("user").(string)

	ok := db.CreateSelection(selection)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).SendString("Error in createSelection operation")
	}

	return c.JSON(selection)
}

// updateSelectionHandler обрабатывает HTTP PUT запросы на /selection/{id}.
// @Summary Обновляет существующий Selection
// @Description Принимает id Selection в качестве параметра пути и JSON тело запроса с новыми полями Selection
// @Tags Selection
// @Accept json
// @Produce json
// @Param Selection body db.Selection true "Create Selection"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{data=db.Selection}
// @Failure 400 {object} ResponseHTTP{}
// @Router /selection/{id} [patch]
func updateSelectionHandler(c *fiber.Ctx) error {
	var selection db.Selection
	if err := c.BodyParser(&selection); err != nil {
		return c.SendString(err.Error())
	}

	ok := db.UpdateSelection(selection)
	if !ok {
		return c.SendString("Error in updateSelection operation")
	}
	return c.SendString("Selection updated successfully")
}

// getManySelectionsHandler обрабатывает HTTP GET запросы на /selections.
// @Summary Получает все Selections
// @Description Возвращает все Selections из базы данных
// @Tags Selection
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{data=[]db.Selection}
// @Failure 400 {object} ResponseHTTP{}
// @Router /selection [get]
func getManySelectionsHandler(c *fiber.Ctx) error {
	ok, selections := db.FindManySelection()
	if !ok {
		return c.SendString("Error in FindManySelection")
	}
	return c.JSON(selections)
}

// getOneSelectionHandler обрабатывает HTTP GET запросы на /selection/{id}.
// @Summary Получает один Selection
// @Description Возвращает один Selection из базы данных по id
// @Tags Selection
// @Accept json
// @Produce json
// @Param id path int true "Selection ID"
// @Success 200 {object} ResponseHTTP{data=db.Selection}
// @Failure 400 {object} ResponseHTTP{}
// @Router /selection/{id} [get]

// func getOneSelectionHandler(c *fiber.Ctx) error {
// 	selectionID := c.Params("id")
// 	session := c.Locals("session")
// 	if session == nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "session not found"})
// 	}
// 	userID := session.(string)
// 	result, ok := db.FindOneSelection(selectionID, userID)
// 	if !ok {
// 		return c.JSON(result)
// 	}
// 	return c.JSON(result)
// }

func getOneSelectionHandler(c *fiber.Ctx) error {
	selectionID := c.Params("selection_id")
	userID := c.Locals("user").(string)
	result, ok := db.FindOneSelection(selectionID, userID)
	if !ok {
		return c.SendString("Failed to get Selection")
	}
	return c.JSON(result)
}

// deleteSelectionHandler обрабатывает HTTP DELETE запросы на /selection/{id}.
// @Summary Удаляет существующий Selection
// @Description Принимает id Selection в качестве параметра пути и удаляет соответствующий Selection
// @Tags Selection
// @Accept json
// @Produce json
// @Param Selection body db.Selection true "Create Selection"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{}
// @Failure 400 {object} ResponseHTTP{}
// @Router /selection/{id} [delete]
func deleteSelectionHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	ok := db.DeleteSelection(id)
	if !ok {
		return c.SendString("Error in deleteSelection operation")
	}
	return c.SendString("Selection deleted successfully")
}

// createGiftToSelectionHandler обрабатывает HTTP POST запросы на /giftToSelection.
// @Summary Создает новый GiftToSelection
// @Description Принимает GiftToSelection в теле запроса и создает соответствующий GiftToSelection
// @Tags GiftToSelection
// @Accept json
// @Produce json
// @Param GiftToSelection body db.GiftToSelection true "Create GiftToSelection"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{}
// @Failure 400 {object} ResponseHTTP{}
// @Router /giftToSelection [post]
func createGiftToSelectionHandler(c *fiber.Ctx) error {
	var giftToSelection db.GiftToSelection
	if err := c.BodyParser(&giftToSelection); err != nil {
		return c.SendString(err.Error())
	}

	ok := db.CreateGiftToSelection(giftToSelection)

	if !ok {
		return c.SendString("Error in createGiftToSelection operation")
	}

	return c.JSON(giftToSelection)
}

// updateGiftToSelectionHandler обрабатывает HTTP PUT запросы на /giftToSelection/{id}.
// @Summary Обновляет существующий GiftToSelection
// @Description Принимает id GiftToSelection в качестве параметра пути и обновляет соответствующий GiftToSelection
// @Tags GiftToSelection
// @Accept json
// @Produce json
// @Param GiftToSelection body db.GiftToSelection true "Update GiftToSelection"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{}
// @Failure 400 {object} ResponseHTTP{}
// @Router /giftToSelection/{id} [put]
func updateGiftToSelectionHandler(c *fiber.Ctx) error {
	var giftToSelection db.GiftToSelection
	if err := c.BodyParser(&giftToSelection); err != nil {
		return c.SendString(err.Error())
	}

	ok := db.UpdateGiftToSelection(giftToSelection)
	if !ok {
		return c.SendString("Error in updateGiftToSelection operation")
	}
	return c.SendString("GiftToSelection updated successfully")
}

// findGiftToSelectionHandler обрабатывает HTTP GET запросы на /giftToSelection.
// @Summary Находит существующий GiftToSelection
// @Description Принимает id GiftToSelection в качестве параметра пути и находит соответствующий GiftToSelection
// @Tags GiftToSelection
// @Accept json
// @Produce json
// @Param id path string true "GiftToSelection ID"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{}
// @Failure 400 {object} ResponseHTTP{}
// @Router /giftToSelection/{id} [get]
func findGiftToSelectionHandler(c *fiber.Ctx) error {

	selectionID := c.Params("id")
	giftInSelection, ok := db.FindGiftToSelection(selectionID)
	if !ok {
		return c.SendString("Error in findGiftToSelection operation")
	}
	return c.JSON(giftInSelection)
}

// deleteGiftToSelectionHandler обрабатывает HTTP DELETE запросы на /giftToSelection/{id}.
// @Summary Удаляет существующий GiftToSelection
// @Description Принимает id GiftToSelection в качестве параметра пути и удаляет соответствующий GiftToSelection
// @Tags GiftToSelection
// @Accept json
// @Produce json
// @Param gift_id path string true "GiftToSelection ID"
// @Param selection_id path string true "GiftToSelection ID"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{}
// @Failure 400 {object} ResponseHTTP{}
// @Router /giftToSelection/{id} [delete]
func deleteGiftToSelectionHandler(c *fiber.Ctx) error {
	GiftID := c.Params("gift_id")
	SelectionID := c.Params("selection_id")
	fmt.Println(GiftID, SelectionID)
	ok := db.DeleteGiftToSelection(SelectionID, GiftID)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).SendString("Error in deleteGiftToSelection operation")
	}
	return c.SendString("GiftToSelection deleted successfully")
}

// createSelectionCategoryHandler обрабатывает HTTP POST запросы на /selectionCategory.
// @Summary Создает новый SelectionCategory
// @Description Принимает SelectionCategory в теле запроса и создает соответствующий SelectionCategory
// @Tags SelectionCategory
// @Accept json
// @Produce json
// @Param SelectionCategory body db.SelectionCategory true "Create SelectionCategory"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{}
// @Failure 400 {object} ResponseHTTP{}
// @Router /selectionCategory [post]
func createSelectionCategoryHandler(c *fiber.Ctx) error {
	var selectionCategory db.SelectionCategory
	if err := c.BodyParser(&selectionCategory); err != nil {
		return c.SendString(err.Error())
	}

	selectionCategory.ID = "selectionCategory_" + xid.New().String()

	ok := db.CreateSelectionCategory(selectionCategory)
	if !ok {
		return c.SendString("Error in createSelectionCategory operation")
	}

	return c.JSON(selectionCategory)
}

// updateSelectionCategoryHandler обрабатывает HTTP PUT запросы на /selectionCategory/{id}.
// @Summary Обновляет существующий SelectionCategory
// @Description Принимает id SelectionCategory в качестве параметра пути и обновляет соответствующий SelectionCategory
// @Tags SelectionCategory
// @Accept json
// @Produce json
// @Param SelectionCategory body db.SelectionCategory true "Update SelectionCategory"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{}
// @Failure 400 {object} ResponseHTTP{}
// @Router /selectionCategory/{id} [put]
func updatedSelectionCategoryHandler(c *fiber.Ctx) error {
	var selectionCategory db.SelectionCategory
	if err := c.BodyParser(&selectionCategory); err != nil {
		return c.SendString(err.Error())
	}
	selectionCategory.ID = c.Params("id")
	ok := db.UpdatedSelectionCategory(selectionCategory)
	if !ok {
		return c.SendString("Error in updateSelectionCategory operation")
	}
	return c.SendString("SelectionCategory updated successfully")
}

// findManySelectionCategoryHandler обрабатывает HTTP GET запросы на /selectionCategory/{id}.
// @Summary Находит существующий SelectionCategory
// @Description Принимает id SelectionCategory в качестве параметра пути и находит соответствующий SelectionCategory
// @Tags SelectionCategory
// @Accept json
// @Produce json
// @Param id path string true "SelectionCategory ID"
// @Success 200 {object} ResponseHTTP{}
// @Failure 400 {object} ResponseHTTP{}
// @Router /selectionCategory/{id} [get]
func findManySelectionCategoryHandler(c *fiber.Ctx) error {
	result, ok := db.FindManySelectionCategory()
	if !ok {
		return c.SendString("Error in findSelectionCategory operation")
	}
	return c.JSON(result)
}

// findOneSelectionCategoryHandler обрабатывает HTTP GET запросы на /selectionCategory/{id}.
// @Summary Находит существующий SelectionCategory
// @Description Принимает id SelectionCategory в качестве параметра пути и находит соответствующий SelectionCategory
// @Tags SelectionCategory
// @Accept json
// @Produce json
// @Param id path string true "SelectionCategory ID"
// @Success 200 {object} ResponseHTTP{}
// @Failure 400 {object} ResponseHTTP{}
// @Router /selectionCategory/{id} [get]
func findOneSelectionCategoryHandler(c *fiber.Ctx) error {
	selectionCategoryID := c.Params("id")
	result, ok := db.FindOneSelectionCategory(selectionCategoryID)
	if !ok {
		return c.SendString("Error in findSelectionCategory operation")
	}
	return c.JSON(result)
}

// deleteSelectionCategoryHandler обрабатывает HTTP DELETE запросы на /selectionCategory/{id}.
// @Summary Удаляет существующий SelectionCategory
// @Description Принимает id SelectionCategory в качестве параметра пути и удаляет соответствующий SelectionCategory
// @Tags SelectionCategory
// @Accept json
// @Produce json
// @Param id path string true "SelectionCategory ID"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{}
// @Failure 400 {object} ResponseHTTP{}
// @Router /selectionCategory/{id} [delete]
func deleteSelectionCategoryHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	ok := db.DeleteSelectionCategory(id)
	if !ok {
		return c.SendString("Error in deleteSelectionCategory operation")
	}
	return c.SendString("SelectionCategory deleted successfully")
}

// createLikeToSelectionHandler обрабатывает HTTP POST запросы на /likeToSelection.
// @Summary Создает новый LikeToSelection
// @Description Принимает LikeToSelection в теле запроса и создает соответствующий LikeToSelection
// @Tags LikeToSelection
// @Accept json
// @Produce json
// @Param LikeToSelection body db.LikeToSelection true "Create LikeToSelection"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{}
// @Failure 400 {object} ResponseHTTP{}
// @Router /likeToSelection [post]
func createLikeToSelectionHandler(c *fiber.Ctx) error {
	var likeToSelection db.LikeToSelection
	if err := c.BodyParser(&likeToSelection); err != nil {
		return c.SendString(err.Error())
	}
	// likeToSelection.UserID = c.Locals("user_id").(string)
	ok := db.CreateLikeToSelection(likeToSelection)
	if !ok {
		return c.SendString("Error in createLikeToSelection operation")
	}

	return c.JSON(likeToSelection)
}

// getLikesCountToSelectionHandler обрабатывает HTTP GET запросы на /likeToSelection/{id}/count.
// @Summary Получает количество лайков для Selection
// @Description Принимает id Selection в качестве параметра пути и возвращает количество лайков для соответствующего Selection
// @Tags LikeToSelection
// @Accept json
// @Produce json
// @Param id path string true "Selection ID"
// @Success 200 {object} ResponseHTTP{}
// @Failure 400 {object} ResponseHTTP{}
// @Router /likeToSelection/{id}/count [get]
func getLikesCountToSelectionHandler(c *fiber.Ctx) error {
	selectionID := c.Params("selection_id")

	count := db.GetLikesCountToSelection(selectionID)
	if count == -1 {
		return c.SendString("Error in getLikesCountToSelection operation")
	}
	return c.SendString(fmt.Sprintf("Likes count: %d", count))
}

// deleteLikeToSelectionHandler обрабатывает HTTP DELETE запросы на /likeToSelection/{id}.
// @Summary Удаляет существующий LikeToSelection
// @Description Принимает id LikeToSelection в качестве параметра пути и удаляет соответствующий LikeToSelection
// @Tags LikeToSelection
// @Accept json
// @Produce json
// @Param id path string true "LikeToSelection ID"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{}
// @Failure 400 {object} ResponseHTTP{}
// @Router /likeToSelection/{id} [delete]
func deleteLikeToSelectionHandler(c *fiber.Ctx) error {
	UserID := "" // c.Locals("user_id").(string)
	SelectionID := c.Params("selection_id")

	ok := db.DeleteLikeToSelection(UserID, SelectionID)
	if !ok {
		return c.SendString("Error in deleteLikeToSelection operation")
	}
	return c.SendString("LikeToSelection deleted successfully")
}

// createCommentToSelectionHandler обрабатывает HTTP POST запросы на /commentToSelection.
// @Summary Создает новый CommentToSelection
// @Description Принимает CommentToSelection в теле запроса и создает соответствующий CommentToSelection
// @Tags CommentToSelection
// @Accept json
// @Produce json
// @Param CommentToSelection body db.CommentToSelection true "Create CommentToSelection"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{}
// @Failure 400 {object} ResponseHTTP{}
// @Router /commentToSelection [post]
func createCommentToSelectionHandler(c *fiber.Ctx) error {
	var commentToSelection db.CommentToSelection
	if err := c.BodyParser(&commentToSelection); err != nil {
		return c.SendString(err.Error())
	}

	commentToSelection.ID = "commentToSelection_" + xid.New().String()
	commentToSelection.CreatedAt = time.Now()
	ok := db.CreateCommentToSelection(commentToSelection)
	if !ok {
		return c.SendString("Error in createCommentToSelection operation")
	}

	return c.JSON(commentToSelection)
}

// getCommentsToSelectionHandler обрабатывает HTTP GET запросы на /commentToSelection/{id}.
// @Summary Получает комментарии для Selection
// @Description Принимает id Selection в качестве параметра пути и возвращает комментарии для соответствующего Selection
// @Tags CommentToSelection
// @Accept json
// @Produce json
// @Param id path string true "Selection ID"
// @Success 200 {object} ResponseHTTP{}
// @Failure 400 {object} ResponseHTTP{}
// @Router /commentToSelection/{id} [get]
func getCommentsToSelectionHandler(c *fiber.Ctx) error {
	id := c.Params("selection_id")

	comments, ok := db.GetCommentsToSelection(id)
	if !ok {
		return c.SendString("Error in getCommentsToSelection operation")
	}
	return c.JSON(comments)
}

// updateCommentToSelectionHandler обрабатывает HTTP PUT запросы на /commentToSelection/{id}.
// @Summary Обновляет существующий CommentToSelection
// @Description Принимает id CommentToSelection в качестве параметра пути и обновляет соответствующий CommentToSelection
// @Tags CommentToSelection
// @Accept json
// @Produce json
// @Param CommentToSelection body db.CommentToSelection true "Update CommentToSelection"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{}
// @Failure 400 {object} ResponseHTTP{}
// @Router /commentToSelection/{id} [put]
func updateCommentToSelectionHandler(c *fiber.Ctx) error {
	var commentToSelection db.CommentToSelection
	if err := c.BodyParser(&commentToSelection); err != nil {
		return c.SendString(err.Error())
	}

	ok := db.UpdateCommentToSelection(commentToSelection)
	if !ok {
		return c.SendString("Error in updateCommentToSelection operation")
	}
	return c.SendString("CommentToSelection updated successfully")
}

// deleteCommentToSelectionHandler обрабатывает HTTP DELETE запросы на /commentToSelection/{id}.
// @Summary Удаляет существующий CommentToSelection
// @Description Принимает id CommentToSelection в качестве параметра пути и удаляет соответствующий CommentToSelection
// @Tags CommentToSelection
// @Accept json
// @Produce json
// @Param id path string true "CommentToSelection ID"
// @Param Authorization header string true "Bearer токен"
// @Success 200 {object} ResponseHTTP{}
// @Failure 400 {object} ResponseHTTP{}
// @Router /commentToSelection/{id} [delete]
func deleteCommentToSelectionHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	ok := db.DeleteCommentToSelection(id)
	if !ok {
		return c.SendString("Error in deleteCommentToSelection operation")
	}
	return c.SendString("CommentToSelection deleted successfully")
}

// Upload godoc
// @Summary Upload a beautiful picture
// @Tags Upload
// @Accept  json
// @Produce json
// @Param Photo body db.Photo true "Upload your beautiful picture"
// @Success 200 {object} ResponseHTTP{}
// @Failure 400 {object} ResponseHTTP{}
// @Failure 500 {object} ResponseHTTP{}
// @Router /upload [post]
func uploadHandler(c *fiber.Ctx) error {
	// Получаем base64 из тела запроса
	var photo db.Photo
	if err := c.BodyParser(&photo); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error parsing request body")
	}
	err := validate.Struct(photo)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).
			SendString(err.Error())
	}

	// Декодируем
	decodedPhoto, err := base64.StdEncoding.DecodeString(photo.Photo)
	if err != nil {
		fmt.Println("Failed to decode the photo!")
		return c.Status(fiber.StatusBadRequest).
			SendString(err.Error())
	}

	// Получаем категорию файла
	category := strings.ToLower(photo.Category)

	// Даём уникальное красивое имя
	newFileName := generateUniqueFileName(category)

	// Сохраняем
	destination := fmt.Sprintf("./public/%s/%s", category, newFileName)
	if err := os.WriteFile(destination, decodedPhoto, 0666); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			SendString(err.Error())
	}
	return c.SendString(fmt.Sprintf("File uploaded successfully: %s", newFileName))
}

func generateUniqueFileName(category string) string {
	return fmt.Sprintf("%s_", category) + 
		time.Now().Format("20060102_150405_") + 
		xid.New().String()
}
