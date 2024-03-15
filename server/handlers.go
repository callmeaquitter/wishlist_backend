package server

import (
	"fmt"
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
// @Summary Creates a new gift.
// @Description get the status of server.
// @Tags Gifts
// @Accept  json
// @Produce json
// @Param Gift body db.Gift true "Create Gift"
// @Success 200 {object} ResponseHTTP{data=db.Gift}
// @Failure 400 {object} ResponseHTTP{}
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

func getOneGiftHandler(c *fiber.Ctx) error {
	var gift db.Gift
	ok := db.FindOneGift(gift)
	if !ok {
		return c.SendString("Error in findOneGift operation")
	}
	return c.SendString("Gift Found Succesfully")
}

// Update Gift godoc
// @Summary Creates a new gift.
// @Description get the status of server.
// @Tags 	Gifts
// @Accept  json
// @Produce json
// @Param Gift body db.Gift true "Create Gift"
// @Success 200 {object} ResponseHTTP{data=db.Gift}
// @Failure 400 {object} ResponseHTTP{}
// @Failure 500 {object} ResponseHTTP{}
// @Router /gifts/{id} [patch]

func updateGiftHandler(c *fiber.Ctx) error {
	var gift db.Gift
	ok := db.UpdateGift(gift)
	if !ok {
		return c.SendString("Error in updateGift operation")
	}
	return c.SendString("Gift updated Succesfully")
}

func superSecretHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(string)
	return c.SendString("This is a super secret route. Hi " + user + "!")
}

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
	user.ID = "user_" + xid.New().String()
	ok := db.CreateUser(user)
	if !ok {
		return c.SendString("Error in CreateUser operation")
	}

	return c.SendString("Register")

}

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

func CreateWishlistHandler(c *fiber.Ctx) error {
	var wishlist db.UserWishlist
	if err := c.BodyParser(&wishlist); err != nil {
		return c.SendString(err.Error())
	}

	wishlist.ID = "wishlist_" + xid.New().String()
	wishlist.UserID = c.Locals("user").(string)

	ok := db.CreateWishlist(wishlist)
	if !ok {
		return c.SendString("Error in Create wishlist operation")
	}

	return c.SendString("Create Wishlist succesfully")

}

func FindManyWishlistsHandler(c *fiber.Ctx) error {
	userID := c.Locals("user").(string)
	ok := db.FindManyWishlists(userID)
	if !ok {
		return c.SendString("Error in FindManyWishlists operation")
	}
	return c.SendString("Get all wishlists succesfully")
}

func FindAllWishesInWishlistHandler(c *fiber.Ctx) error {
	wishlistID := c.Params("wishlist_id")
	ok := db.GetManyWishesInWishlist(wishlistID)
	if !ok {
		return c.SendString("Error in FindAllWishesInWishlist operation")
	}

	return c.SendString("Get all wishlists succesfully")
}

func UpdateWishlist(c *fiber.Ctx) error {
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

// deleteGift godoc
// @Summary Creates a new gift.
// @Description get the status of server.
// @Tags Wishlist
// @Accept  json
// @Produce json
// @Param Wishlist body db.UserWishlist true "Delete Wishlist"
// @Success 200 {object} ResponseHTTP{data=db.UserWishlist}
// @Failure 400 {object} ResponseHTTP{}
// @Router /wishlists/wishlist_cnothhc69lbkfh15tmmg/gift_cnos0qk69lbkli6i79ug/user_cnot2oc69lbksn28kko0 [delete]
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
