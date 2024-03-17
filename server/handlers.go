package server

import (
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

