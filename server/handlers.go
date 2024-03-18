package server

import (
	"fmt"
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
// @Param Gift body db.Gift true "Create Gift"
// @Success 200 {object} ResponseHTTP{data=db.Gift}
// @Failure 400 {object} ResponseHTTP{}
// @Router /gifts [post]
// @Router /gifts [post]
func createGiftHandler(c *fiber.Ctx) error {
	var gift db.Gift
	if err := c.BodyParser(&gift); err != nil {
		return c.SendString(err.Error())
	}

	if gift.Name == "" {
		return c.SendString("Name is required")
	}
	if gift.Price == 0 {
		return c.SendString("Price is required")
	}
	if gift.Link == "" {
		return c.SendString("Link is required")
	}
	if gift.Photo == "" {
		return c.SendString("Photo is required")
	}

	gift.ID = "gift_" + xid.New().String()
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
	var gift db.Gift
	ok := db.FindManyGift(gift)
	if !ok {
		return c.SendString("Error in findManyGifts operation")
	}
	return c.SendString("Gifts Found Succesfully")
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
	var gift db.Gift
	ok := db.FindOneGift(gift)
	if !ok {
		return c.SendString("Error in findOneGift operation")
	}
	return c.SendString("Gift Found Succesfully")
}

// Update Gift godoc
// @Summary update gift by ID
// @Description get the status of server.
// @Tags 	Gifts
// @Accept  json
// @Produce json
// @Param Gift body db.Gift true "Update Gift"
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
// @Success 200 {object} ResponseHTTP{data=db.BookedGiftInWishlist}
// @Failure 400 {object} ResponseHTTP{}
// @Router /booked_gifts/create [post]
func createBookedGiftInWishlist(c *fiber.Ctx) error {
	var bookedGiftInWishlist db.BookedGiftInWishlist
	if err := c.BodyParser(&bookedGiftInWishlist); err != nil {
		return c.SendString(err.Error())
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
// @Param GiftCategory body db.GiftCategory true "Gift Category object to be created"
// @Success 200 {object} ResponseHTTP{data=[]db.GiftCategory}
// @Failure 400 {string} string "CategoryName is required"
// @Failure 400 {string} string "Failed to create gift category"
// @Router /gift_category/create [post]
func createGiftCategory(c *fiber.Ctx) error {
	var giftCategory db.GiftCategory
	if err := c.BodyParser(&giftCategory); err != nil {
		return c.SendString(err.Error())
	}
	if giftCategory.Name == "" {
		return c.SendString("CategoryName is required")
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

// createGiftReviwHandler godoc
// @Summary Create a new review for gift.
// @Description Create a new review for gift.
// @Tags GiftReview
// @Accept  json
// @Produce json
// @Param Gift body db.GiftReview true "Create GiftReview"
// @Success 200 {object} ResponseHTTP{data=db.GiftReview}
// @Failure 400 {object} ResponseHTTP{}
// @Router /gift_review [post]
func createGiftReviwHandler(c *fiber.Ctx) error {
	var giftReview db.GiftReview
	if err := c.BodyParser(&giftReview); err != nil {
		return c.SendString(err.Error())
	}
	giftReview.ID = "review_" + xid.New().String()

	if giftReview.Mark == 0.0 {
		return c.SendString("Mark is required")
	}

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
// @Success 200 {object} ResponseHTTP{data=[]db.GiftReview}
// @Failure 503 {object} ResponseHTTP{}
// @Router /gift_review/{id} [get]
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
// @Success 200 {object} ResponseHTTP{data=[]db.GiftReview}
// @Failure 503 {object} ResponseHTTP{}
// @Router /gift_review/{gift_id} [get]
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
// @Router /gift_review/{gift_id} [get]
func calculateAverageMarkByGiftIDHandler(c *fiber.Ctx) error {
	giftID := c.Params("gift_id")
	averageMark, ok := db.CalculateAverageMarkByGiftID(giftID)
	if !ok {
		return c.SendString("calculateAverageMarkByGiftID operation")
	}
	return c.JSON(averageMark)
}

// createSeller godoc
// @Summary Creates a new seller.
// @Tags Sellers
// @Accept json
// @Produce json
// @Param Seller body db.Seller true "Create Seller"
// @Success 200 {object} ResponseHTTP{data=db.Seller}
// @Failure 400 {object} ResponseHTTP{}
// @Router /sellers [post]
func createSellerHandler(c *fiber.Ctx) error {
	var seller db.Seller
	if err := c.BodyParser(&seller); err != nil {
		return c.SendString(err.Error())
	}

	err := validate.Struct(seller)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}

	seller.SellerID = "seller_" + xid.New().String()

	ok := db.CreateSeller(seller)
	if !ok {
		return c.SendString("Error in createSeller operation")
	}

	return c.JSON(seller)
}

// deleteSeller godoc
// @Summary Deletes a specified seller.
// @Tags Sellers
// @Accept json
// @Produce json
// @Param id path string true "Delete Seller"
// @Success 200 {object} ResponseHTTP{data=db.Seller}
// @Failure 400 {object} ResponseHTTP{}
// @Router /sellers/{id} [delete]
func deleteSellerHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	ok := db.DeleteSeller(id)
	if !ok {
		return c.SendString("Error in deleteSeller operation")
	}
	return c.SendString("Seller deleted successfully")
}

// getManySellers godoc
// @Summary Fetches all sellers.
// @Tags Sellers
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{data=db.Seller}
// @Failure 400 {object} ResponseHTTP{}
// @Router /sellers [get]
func getManySellersHandler(c *fiber.Ctx) error {
	result, ok := db.FindManySeller()
	if !ok {
		return c.SendString("Error in findManySellers operation")
	}
	return c.JSON(result)
}

// getOneSeller godoc
// @Summary Fetches a specific seller.
// @Tags Sellers
// @Accept json
// @Produce json
// @Param id path string true "Seller ID"
// @Success 200 {object} ResponseHTTP{data=db.Seller}
// @Failure 400 {object} ResponseHTTP{}
// @Router /sellers/{id} [get]
func getOneSellerHandler(c *fiber.Ctx) error {
	sellerId := c.Params("id")
	result, ok := db.FindOneSeller(sellerId)
	if !ok {
		return c.SendString("Error in findOneSeller operation")
	}
	return c.JSON(result)
}

// updateSeller godoc
// @Summary Updates an existing seller.
// @Tags Sellers
// @Accept json
// @Produce json
// @Param Seller body db.Seller true "Update Seller"
// @Success 200 {object} ResponseHTTP{data=db.Seller}
// @Failure 400 {object} ResponseHTTP{}
// @Router /sellers/{id} [patch]
func updateSellerHandler(c *fiber.Ctx) error {
	var seller db.Seller
	if err := c.BodyParser(&seller); err != nil {
		return c.SendString(err.Error())
	}

	err := validate.Struct(seller)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}

	ok := db.UpdateSeller(seller)
	if !ok {
		return c.SendString("Error in updateSeller operation")
	}
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Successfully updated the Seller.",
		Data:    &seller,
	})
}

// createService godoc
// @Summary Creates a new service.
// @Tags Services
// @Accept json
// @Produce json
// @Param Service body db.Service true "Create Service"
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

// getOneService godoc
// @Summary Fetches a specific service.
// @Tags Services
// @Accept json
// @Produce json
// @Param id path string true "Service ID"
// @Success 200 {object} ResponseHTTP{data=db.Service}
// @Failure 400 {object} ResponseHTTP{}
// @Router /services/{id} [get]
func getOneServiceHandler(c *fiber.Ctx) error {
	serviceId := c.Params("id")
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
// @Param Service body db.Service true "Update Service"
// @Success 200 {object} ResponseHTTP{data=db.Service}
// @Failure 400 {object} ResponseHTTP{}
// @Router /services/{id} [patch]
func updateServiceHandler(c *fiber.Ctx) error {
	var service db.Service
	if err := c.BodyParser(&service); err != nil {
		return c.SendString(err.Error())
	}

	err := validate.Struct(service)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}
	if !service.Price.IsPositive() {
		return c.Status(fiber.StatusUnprocessableEntity).
			SendString("Only positive deicmals are allowed!")
	}

	ok := db.UpdateService(service)
	if !ok {
		return c.SendString("Error in updateService operation")
	}
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Successfully updated the Service.",
		Data:    &service,
	})
}

// deleteService godoc
// @Summary Deletes a specified service.
// @Tags Services
// @Accept json
// @Produce json
// @Param id path string true "Delete Service"
// @Success 200 {object} ResponseHTTP{data=db.Service}
// @Failure 400 {object} ResponseHTTP{}
// @Router /services/{id} [delete]
func deleteServiceHandler(c *fiber.Ctx) error {
	id := c.Params("id")

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
// @Param SellerToService body db.SellerToService true "Create Selllers-Services"
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
// @Param ServiceReview body db.ServiceReview true "Update ServiceReview"
// @Success 200 {object} ResponseHTTP{data=db.ServiceReview}
// @Failure 400 {object} ResponseHTTP{}
// @Router /serviceReviews/{id} [patch]
func updateServiceReviewHandler(c *fiber.Ctx) error {
	var serviceReview db.ServiceReview
	if err := c.BodyParser(&serviceReview); err != nil {
		return c.SendString(err.Error())
	}

	err := validate.Struct(serviceReview)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}
	if serviceReview.Mark.IsNegative() ||
		serviceReview.Mark.GreaterThan(decimal.NewFromInt(5)) {
		return c.Status(fiber.StatusUnprocessableEntity).
			SendString("Only positive marks less or equal to 5 are allowed!")
	}

	serviceReview.UpdateDate = time.Now()

	ok := db.UpdateServiceReview(serviceReview)
	if !ok {
		return c.SendString("Error in updateServiceReview operation")
	}
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Successfully updated the ServiceReview",
		Data:    &serviceReview,
	})
}

// deleteServiceReview godoc
// @Summary Deletes a specified serviceReview.
// @Tags ServiceReviews
// @Accept json
// @Produce json
// @Param id path string true "Delete ServiceReview"
// @Success 200 {object} ResponseHTTP{data=db.ServiceReview}
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

// createSeller godoc
// @Summary Creates a new seller.
// @Tags Sellers
// @Accept json
// @Produce json
// @Param Seller body db.Seller true "Create Seller"
// @Success 200 {object} ResponseHTTP{data=db.Seller}
// @Failure 400 {object} ResponseHTTP{}
// @Router /sellers [post]
func createSellerHandler(c *fiber.Ctx) error {
	var seller db.Seller
	if err := c.BodyParser(&seller); err != nil {
		return c.SendString(err.Error())
	}

	err := validate.Struct(seller)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}

	seller.SellerID = "seller_" + xid.New().String()

	ok := db.CreateSeller(seller)
	if !ok {
		return c.SendString("Error in createSeller operation")
	}

	return c.JSON(seller)
}

// deleteSeller godoc
// @Summary Deletes a specified seller.
// @Tags Sellers
// @Accept json
// @Produce json
// @Param id path string true "Delete Seller"
// @Success 200 {object} ResponseHTTP{data=db.Seller}
// @Failure 400 {object} ResponseHTTP{}
// @Router /sellers/{id} [delete]
func deleteSellerHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	ok := db.DeleteSeller(id)
	if !ok {
		return c.SendString("Error in deleteSeller operation")
	}
	return c.SendString("Seller deleted successfully")
}

// getManySellers godoc
// @Summary Fetches all sellers.
// @Tags Sellers
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{data=db.Seller}
// @Failure 400 {object} ResponseHTTP{}
// @Router /sellers [get]
func getManySellersHandler(c *fiber.Ctx) error {
	result, ok := db.FindManySeller()
	if !ok {
		return c.SendString("Error in findManySellers operation")
	}
	return c.JSON(result)
}

// getOneSeller godoc
// @Summary Fetches a specific seller.
// @Tags Sellers
// @Accept json
// @Produce json
// @Param id path string true "Seller ID"
// @Success 200 {object} ResponseHTTP{data=db.Seller}
// @Failure 400 {object} ResponseHTTP{}
// @Router /sellers/{id} [get]
func getOneSellerHandler(c *fiber.Ctx) error {
	sellerId := c.Params("id")
	result, ok := db.FindOneSeller(sellerId)
	if !ok {
		return c.SendString("Error in findOneSeller operation")
	}
	return c.JSON(result)
}

// updateSeller godoc
// @Summary Updates an existing seller.
// @Tags Sellers
// @Accept json
// @Produce json
// @Param Seller body db.Seller true "Update Seller"
// @Success 200 {object} ResponseHTTP{data=db.Seller}
// @Failure 400 {object} ResponseHTTP{}
// @Router /sellers/{id} [patch]
func updateSellerHandler(c *fiber.Ctx) error {
	var seller db.Seller
	if err := c.BodyParser(&seller); err != nil {
		return c.SendString(err.Error())
	}

	err := validate.Struct(seller)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}

	ok := db.UpdateSeller(seller)
	if !ok {
		return c.SendString("Error in updateSeller operation")
	}
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Successfully updated the Seller.",
		Data:    &seller,
	})
}

// createService godoc
// @Summary Creates a new service.
// @Tags Services
// @Accept json
// @Produce json
// @Param Service body db.Service true "Create Service"
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

// getOneService godoc
// @Summary Fetches a specific service.
// @Tags Services
// @Accept json
// @Produce json
// @Param id path string true "Service ID"
// @Success 200 {object} ResponseHTTP{data=db.Service}
// @Failure 400 {object} ResponseHTTP{}
// @Router /services/{id} [get]
func getOneServiceHandler(c *fiber.Ctx) error {
	serviceId := c.Params("id")
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
// @Param Service body db.Service true "Update Service"
// @Success 200 {object} ResponseHTTP{data=db.Service}
// @Failure 400 {object} ResponseHTTP{}
// @Router /services/{id} [patch]
func updateServiceHandler(c *fiber.Ctx) error {
	var service db.Service
	if err := c.BodyParser(&service); err != nil {
		return c.SendString(err.Error())
	}

	err := validate.Struct(service)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}
	if !service.Price.IsPositive() {
		return c.Status(fiber.StatusUnprocessableEntity).
			SendString("Only positive deicmals are allowed!")
	}

	ok := db.UpdateService(service)
	if !ok {
		return c.SendString("Error in updateService operation")
	}
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Successfully updated the Service.",
		Data:    &service,
	})
}

// deleteService godoc
// @Summary Deletes a specified service.
// @Tags Services
// @Accept json
// @Produce json
// @Param id path string true "Delete Service"
// @Success 200 {object} ResponseHTTP{data=db.Service}
// @Failure 400 {object} ResponseHTTP{}
// @Router /services/{id} [delete]
func deleteServiceHandler(c *fiber.Ctx) error {
	id := c.Params("id")

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
// @Param SellerToService body db.SellerToService true "Create Selllers-Services"
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
// @Param ServiceReview body db.ServiceReview true "Update ServiceReview"
// @Success 200 {object} ResponseHTTP{data=db.ServiceReview}
// @Failure 400 {object} ResponseHTTP{}
// @Router /serviceReviews/{id} [patch]
func updateServiceReviewHandler(c *fiber.Ctx) error {
	var serviceReview db.ServiceReview
	if err := c.BodyParser(&serviceReview); err != nil {
		return c.SendString(err.Error())
	}

	err := validate.Struct(serviceReview)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}
	if serviceReview.Mark.IsNegative() ||
		serviceReview.Mark.GreaterThan(decimal.NewFromInt(5)) {
		return c.Status(fiber.StatusUnprocessableEntity).
			SendString("Only positive marks less or equal to 5 are allowed!")
	}

	serviceReview.UpdateDate = time.Now()

	ok := db.UpdateServiceReview(serviceReview)
	if !ok {
		return c.SendString("Error in updateServiceReview operation")
	}
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Successfully updated the ServiceReview",
		Data:    &serviceReview,
	})
}

// deleteServiceReview godoc
// @Summary Deletes a specified serviceReview.
// @Tags ServiceReviews
// @Accept json
// @Produce json
// @Param id path string true "Delete ServiceReview"
// @Success 200 {object} ResponseHTTP{data=db.ServiceReview}
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

	if user.Login == "" {
		return c.SendString("Login is required") //TODO: Check unique username
	}

	if user.Password == "" {
		return c.SendString("Password is required")
	}

	user.ID = ""
	user.ID = "user_" + xid.New().String()
	ok := db.CreateUser(user)
	if !ok {
		return c.SendString("Error in CreateUser operation")
	}

	return c.SendString("Register")

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

// AddWishHandler godoc
// @Summary Creates a new gift.
// @Description get the status of server.
// @Tags Wishes
// @Accept  json
// @Produce json
// @Param Wishes body db.Wishes true "Add wishes"
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
	ok := db.AddWish(wishlist_id, gift_id)
	if !ok {
		return c.SendString("Error in Create operation")
	}
	return c.SendString("Create Wish succesfully")
}

// DeleteWishHandler godoc
// @Summary Creates a new gift.
// @Description get the status of server.
// @Tags Wishes
// @Accept  json
// @Produce json
// @Param Wishes body db.Wishes true "Add wishes"
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
// @Summary Creates a new gift.
// @Description get the status of server.
// @Tags Wishlist
// @Accept  json
// @Produce json
// @Param UserWishlist body db.UserWishlist true "Create Wishlist"
// @Success 200 {object} ResponseHTTP{data=db.UserWishlist}
// @Failure 400 {object} ResponseHTTP{}
// @Router /wishlists [post]
func CreateWishlistHandler(c *fiber.Ctx) error {
	var wishlist db.UserWishlist
	if err := c.BodyParser(&wishlist); err != nil {
		return c.SendString(err.Error())
	}
	if wishlist.Name == "" {
		return c.SendString("Name is required")
	}

	wishlist.ID = "wishlist_" + xid.New().String()
	wishlist.UserID = c.Locals("user").(string)

	ok := db.CreateWishlist(wishlist)
	if !ok {
		return c.SendString("Error in Create wishlist operation")
	}

	return c.SendString("Create Wishlist succesfully")

}

// createWishlist godoc
// @Summary Creates a new gift.
// @Description get the status of server.
// @Tags Wishlist
// @Accept  json
// @Produce json
// @Param UserWishlist body db.UserWishlist true "Create Wishlist"
// @Success 200 {object} ResponseHTTP{data=db.UserWishlist}
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

// DeleteWishHandler godoc
// @Summary Creates a new gift.
// @Description get the status of server.
// @Tags Wishes
// @Accept  json
// @Produce json
// @Param Wishes body db.Wishes true "Add wishes"
// @Success 200 {object} ResponseHTTP{data=db.Wishes}
// @Failure 400 {object} ResponseHTTP{}
// @Router /wishes/{id}/{wishlist_id} [get]
func FindAllWishesInWishlistHandler(c *fiber.Ctx) error {
	wishlistID := c.Params("wishlist_id")
	wishes, ok := db.GetManyWishesInWishlist(wishlistID)
	if !ok {
		return c.SendString("Invalid credentials")
	}

	return c.JSON(AuthResponse{Session: session})
}
