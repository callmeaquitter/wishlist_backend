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
// @Accept json
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

func deleteGiftHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	ok := db.DeleteGift(id)
	if !ok {
		return c.SendString("Error in deleteGift operation")
	}
	return c.SendString("Gift deleted successfully")
}

// GetAllGifts is a function to get all books data from database
// @Summary То что делает функция делает
// @Description Описание функции
// @Tags тут пишите свою область ответственности
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{data=[]db.ваша_структура}
// @Failure 503 {object} ResponseHTTP{}
// @Router /путь/ендпоинт [get]

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

// createSelectionHandler обрабатывает HTTP POST запросы на /selection/create.
// @Summary Создает новый Selection
// @Description Принимает JSON тело запроса с полями Selection и создает новый Selection
// @Tags Selection
// @Accept json
// @Produce json
// @Param Selection body db.Selection true "Create Selection"
// @Success 200 {object} ResponseHTTP{data=db.Selection}
// @Failure 400 {object} ResponseHTTP{}
// @Router /selection/create [post]
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

	selection.UserID = "111" //c.Locals("user")

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
// @Success 200 {object} ResponseHTTP{data=db.Selection}
// @Failure 400 {object} ResponseHTTP{}
// @Router /selection/{id} [put]
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
	var selection db.Selection
	ok, result := db.FindManySelection(selection)
	if !ok {
		return c.SendString("Error in FindManySelection")
	}
	return c.JSON(result)
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
func getOneSelectionHandler(c *fiber.Ctx) error {
	var selection db.Selection
	ok := db.FindOneSelection(selection)
	if !ok {
		return c.SendString("Error in FindOneSelection")
	}
	return c.SendString("Selection find successfully")
}

// deleteSelectionHandler обрабатывает HTTP DELETE запросы на /selection/{id}.
// @Summary Удаляет существующий Selection
// @Description Принимает id Selection в качестве параметра пути и удаляет соответствующий Selection
// @Tags Selection
// @Accept json
// @Produce json
// @Param Selection body db.Selection true "Create Selection"
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

// findGiftToSelectionHandler обрабатывает HTTP GET запросы на /giftToSelection/{id}.
// @Summary Находит существующий GiftToSelection
// @Description Принимает id GiftToSelection в качестве параметра пути и находит соответствующий GiftToSelection
// @Tags GiftToSelection
// @Accept json
// @Produce json
// @Param id path string true "GiftToSelection ID"
// @Success 200 {object} ResponseHTTP{}
// @Failure 400 {object} ResponseHTTP{}
// @Router /giftToSelection/{id} [get]
func findGiftToSelectionHandler(c *fiber.Ctx) error {
	var giftToSelection db.GiftToSelection
	if err := c.BodyParser(&giftToSelection); err != nil {
		return c.SendString(err.Error())
	}

	ok := db.FindGiftToSelection(giftToSelection)
	if !ok {
		return c.SendString("Error in findGiftToSelection operation")
	}
	return c.JSON(giftToSelection)
}

// deleteGiftToSelectionHandler обрабатывает HTTP DELETE запросы на /giftToSelection/{id}.
// @Summary Удаляет существующий GiftToSelection
// @Description Принимает id GiftToSelection в качестве параметра пути и удаляет соответствующий GiftToSelection
// @Tags GiftToSelection
// @Accept json
// @Produce json
// @Param id path string true "GiftToSelection ID"
// @Success 200 {object} ResponseHTTP{}
// @Failure 400 {object} ResponseHTTP{}
// @Router /giftToSelection/{id} [delete]
func deleteGiftToSelectionHandler(c *fiber.Ctx) error {
	GiftID := c.Params("id")
	SelectionID := ""
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
// @Success 200 {object} ResponseHTTP{}
// @Failure 400 {object} ResponseHTTP{}
// @Router /selectionCategory/{id} [put]
func updatedSelectionCategoryHandler(c *fiber.Ctx) error {
	var selectionCategory db.SelectionCategory
	if err := c.BodyParser(&selectionCategory); err != nil {
		return c.SendString(err.Error())
	}

	ok := db.UpdatedSelectionCategory(selectionCategory)
	if !ok {
		return c.SendString("Error in updateSelectionCategory operation")
	}
	return c.SendString("SelectionCategory updated successfully")
}

// findSelectionCategoryHandler обрабатывает HTTP GET запросы на /selectionCategory/{id}.
// @Summary Находит существующий SelectionCategory
// @Description Принимает id SelectionCategory в качестве параметра пути и находит соответствующий SelectionCategory
// @Tags SelectionCategory
// @Accept json
// @Produce json
// @Param id path string true "SelectionCategory ID"
// @Success 200 {object} ResponseHTTP{}
// @Failure 400 {object} ResponseHTTP{}
// @Router /selectionCategory/{id} [get]
func findSelectionCategoryHandler(c *fiber.Ctx) error {
	var selectionCategory db.SelectionCategory
	if err := c.BodyParser(&selectionCategory); err != nil {
		return c.SendString(err.Error())
	}

	ok := db.FindSelectionCategory(selectionCategory)
	if !ok {
		return c.SendString("Error in findSelectionCategory operation")
	}
	return c.JSON(selectionCategory)
}

// deleteSelectionCategoryHandler обрабатывает HTTP DELETE запросы на /selectionCategory/{id}.
// @Summary Удаляет существующий SelectionCategory
// @Description Принимает id SelectionCategory в качестве параметра пути и удаляет соответствующий SelectionCategory
// @Tags SelectionCategory
// @Accept json
// @Produce json
// @Param id path string true "SelectionCategory ID"
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
// @Success 200 {object} ResponseHTTP{}
// @Failure 400 {object} ResponseHTTP{}
// @Router /likeToSelection [post]
func createLikeToSelectionHandler(c *fiber.Ctx) error {
	var likeToSelection db.LikeToSelection
	if err := c.BodyParser(&likeToSelection); err != nil {
		return c.SendString(err.Error())
	}

	likeToSelection.SelectionID = "likeToSelection_" + xid.New().String()

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
// @Success 200 {object} ResponseHTTP{}
// @Failure 400 {object} ResponseHTTP{}
// @Router /likeToSelection/{id} [delete]
func deleteLikeToSelectionHandler(c *fiber.Ctx) error {
	UserID := c.Params("user_id")
	SelectionID, ok := c.Locals("selection_id").(string)

	if !ok {
		return c.SendString("Error: selection_id is not a string or is missing")
	}

	ok = db.DeleteLikeToSelection(UserID, SelectionID)
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
// @Success 200 {object} ResponseHTTP{}
// @Failure 400 {object} ResponseHTTP{}
// @Router /commentToSelection [post]
func createCommentToSelectionHandler(c *fiber.Ctx) error {
	var commentToSelection db.CommentToSelection
	if err := c.BodyParser(&commentToSelection); err != nil {
		return c.SendString(err.Error())
	}

	commentToSelection.ID = "commentToSelection_" + xid.New().String()

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
	id := c.Params("id")

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
