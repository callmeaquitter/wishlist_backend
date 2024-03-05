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
// @Accept */*
// @Produce json
// @Success 200 {object} ResponseHTTP{data=db.Gift}
// @Failure 400 {object} ResponseHTTP{}
// @Router / [post]
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

func deleteGiftHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	ok := db.DeleteGift(id)
	if !ok {
		return c.SendString("Error in deleteGift operation")
	}
	return c.SendString("Gift deleted successfully")
}

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

func updateGiftHandler(c *fiber.Ctx) error {
	var gift db.Gift
	ok := db.UpdateGift(gift)
	if !ok {
		return c.SendString("Error in updateGift operation")
	}
	return c.SendString("Gift updated Succesfully")
}

// createSeller godoc
// @Summary Creates a new seller.
// @Description get the status of server.
// @Tags Sellers
// @Accept */*
// @Produce json
// @Success 200 {object} ResponseHTTP{data=db.Seller}
// @Failure 400 {object} ResponseHTTP{}
// @Router / [post]
func createSellerHandler(c *fiber.Ctx) error {
	var seller db.Seller
	if err := c.BodyParser(&seller); err != nil {
		return c.SendString(err.Error())
	}

	if seller.Name == "" {
		return c.SendString("Name is required")
	}
	if seller.Email == "" {
		return c.SendString("Email is required")
	}
	if seller.Photo == "" {
		return c.SendString("Photo is required")
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
// @Description get the status of server.
// @Tags Sellers
// @Accept */*
// @Produce json
// @Success 200 {object} ResponseHTTP{data=db.Seller}
// @Failure 400 {object} ResponseHTTP{}
// @Router / [delete]
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
// @Description get the status of server.
// @Tags Sellers
// @Accept */*
// @Produce json
// @Success 200 {object} ResponseHTTP{data=db.Seller}
// @Failure 400 {object} ResponseHTTP{}
// @Router / [get]
func getManySellersHandler(c *fiber.Ctx) error {
	var seller db.Seller
	ok := db.FindManySeller(seller)
	if !ok {
		return c.SendString("Error in findManySellers operation")
	}
	return c.SendString("Sellers Found Succesfully")
}

// getOneSeller godoc
// @Summary Fetches a specific seller.
// @Description get the status of server.
// @Tags Sellers
// @Accept */*
// @Produce json
// @Success 200 {object} ResponseHTTP{data=db.Seller}
// @Failure 400 {object} ResponseHTTP{}
// @Router / [get]
func getOneSellerHandler(c *fiber.Ctx) error {
	var seller db.Seller
	ok := db.FindOneSeller(seller)
	if !ok {
		return c.SendString("Error in findOneSeller operation")
	}
	return c.SendString("Seller Found Succesfully")
}

// updateSeller godoc
// @Summary Updates an existing seller.
// @Description get the status of server.
// @Tags Sellers
// @Accept */*
// @Produce json
// @Success 200 {object} ResponseHTTP{data=db.Seller}
// @Failure 400 {object} ResponseHTTP{}
// @Router / [put]
func updateSellerHandler(c *fiber.Ctx) error {
	var seller db.Seller
	ok := db.UpdateSeller(seller)
	if !ok {
		return c.SendString("Error in updateSeller operation")
	}
	return c.SendString("Seller updated Succesfully")
}

// createService godoc
// @Summary Creates a new service.
// @Description get the status of server.
// @Tags Services
// @Accept */*
// @Produce json
// @Success 200 {object} ResponseHTTP{data=db.Service}
// @Failure 400 {object} ResponseHTTP{}
// @Router / [post]
func createServiceHandler(c *fiber.Ctx) error {
	var service db.Service
	if err := c.BodyParser(&service); err != nil {
		return c.SendString(err.Error())
	}

	if service.Name == "" {
		return c.SendString("Name is required")
	}
	if service.Price == 0 {
		return c.SendString("Price is required")
	}
	if service.Location == "" {
		return c.SendString("Location is required")
	}
	if service.Photos == "" {
		return c.SendString("Photos are required")
	}

	service.ServiceID = "service_" + xid.New().String()

	ok := db.CreateService(service)
	if !ok {
		return c.SendString("Error in createService operation")
	}

	return c.JSON(service)
}

// deleteService godoc
// @Summary Deletes a specified service.
// @Description get the status of server.
// @Tags Services
// @Accept */*
// @Produce json
// @Success 200 {object} ResponseHTTP{data=db.Service}
// @Failure 400 {object} ResponseHTTP{}
// @Router / [delete]
func deleteServiceHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	ok := db.DeleteService(id)
	if !ok {
		return c.SendString("Error in deleteService operation")
	}
	return c.SendString("Service deleted successfully")
}

// getManyServices godoc
// @Summary Fetches all services.
// @Description get the status of server.
// @Tags Services
// @Accept */*
// @Produce json
// @Success 200 {object} ResponseHTTP{data=db.Service}
// @Failure 400 {object} ResponseHTTP{}
// @Router / [get]
func getManyServicesHandler(c *fiber.Ctx) error {
	var service db.Service
	ok := db.FindManyService(service)
	if !ok {
		return c.SendString("Error in findManyServices operation")
	}
	return c.SendString("Services Found Succesfully")
}

// getOneService godoc
// @Summary Fetches a specific service.
// @Description get the status of server.
// @Tags Services
// @Accept */*
// @Produce json
// @Success 200 {object} ResponseHTTP{data=db.Service}
// @Failure 400 {object} ResponseHTTP{}
// @Router / [get]
func getOneServiceHandler(c *fiber.Ctx) error {
	var service db.Service
	ok := db.FindOneService(service)
	if !ok {
		return c.SendString("Error in findOneService operation")
	}
	return c.SendString("Service Found Succesfully")
}

// updateService godoc
// @Summary Updates an existing service.
// @Description get the status of server.
// @Tags Services
// @Accept */*
// @Produce json
// @Success 200 {object} ResponseHTTP{data=db.Service}
// @Failure 400 {object} ResponseHTTP{}
// @Router / [put]
func updateServiceHandler(c *fiber.Ctx) error {
	var service db.Service
	ok := db.UpdateService(service)
	if !ok {
		return c.SendString("Error in updateService operation")
	}
	return c.SendString("Service updated Succesfully")
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

