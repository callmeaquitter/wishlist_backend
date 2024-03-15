package server

import (
	"wishlist/db"
	_ "wishlist/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/xid"
)

type ResponseHTTP struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// createGift godoc
// @Summary Creates a new gift.
// @Description get the status of server.
// @Tags Gifts
// @Accept  json
// @Produce json
// @Param Gift body db.Gift true "Create Gift"
// @Success 200 {object} ResponseHTTP{data=db.Gift}
// @Failure 400 {object} ResponseHTTP{}
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

// GetAllGifts is a function to get all books data from database
// @Summary Get all books
// @Description Get all books
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
// @Summary Get one books
// @Description Get one books
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
func getGiftReviewByIDHandler(c *fiber.Ctx) error{
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
func getGiftReviewsByGiftIDHandler(c *fiber.Ctx) error{
	giftID := c.Params("gift_id")
	giftReviews, ok := db.GetGiftReviewsByGiftID(giftID)
	if !ok {
		return c.SendString("Error in getGiftReviewsByGiftID operation")
	}
	return c.JSON(giftReviews)
}

func superSecretHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(string)
	return c.SendString("This is a super secret route. Hi " + user + "!")
}

func registerHandler(c *fiber.Ctx) error {
	return c.SendString("Register")
}

func loginHandler(c *fiber.Ctx) error {
	var authCredentials AuthCredentials
	if err := c.BodyParser(&authCredentials); err != nil {
		return c.SendString(err.Error())
	}

	session, ok := getUser(authCredentials.Login, authCredentials.Password)
	if !ok {
		return c.SendString("Invalid credentials")
	}

	return c.JSON(AuthResponse{Session: session})
}
