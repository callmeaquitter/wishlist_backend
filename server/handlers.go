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
// @Tags Gift
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
		return c.SendString("Error in findManyGift operation")
	}
	return c.SendString("Gift Found Succesfully")
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
	// TODO: fix this!
	// err := validate.Struct(quest)
	// if err != nil {
	// 	return c.Status(fiber.StatusUnprocessableEntity).
	// 		SendString(err.Error())
	// }

	quest.ID = "quest_" + xid.New().String()

	quest.UserID = c.Locals("user").(string)

	ok := db.CreateQuest(quest)
	if !ok {
		return c.SendString("Error in createQuest operation")
	}

	return c.JSON(quest)
}

// getManyQuestHandler обрабатывает HTTP GET запросы на /quest
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

	// err := validate.Struct(subquest)
	// if err != nil {
	// 	return c.Status(fiber.StatusUnprocessableEntity).
	// 		SendString(err.Error())
	// }

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
	var subquest db.Subquest
	ok := db.UpdateSubquest(subquest)
	//TODO: А где здесь парсер?
	if !ok {
		return c.SendString("Error in updateSubquest operation")
	}
	return c.SendString("Subquest updated Succesfully")
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

	// err := validate.Struct(tasks)
	// if err != nil {
	// 	return c.Status(fiber.StatusUnprocessableEntity).
	// 		SendString(err.Error())
	// }

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
	var tasks db.Tasks
	ok := db.UpdateTasks(tasks)
	if !ok {
		return c.SendString("Error in updateTasks operation")
	}
	return c.SendString("Tasks updated Succesfully")
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

	// err := validate.Struct(offlineshops)
	// if err != nil {
	// 	return c.Status(fiber.StatusUnprocessableEntity).
	// 		SendString(err.Error())
	// }

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
	var offlineshops db.OfflineShops
	ok := db.UpdateOfflineShops(offlineshops)
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
	var offlineshops db.OfflineShops
	ok := db.FindOneOfflineShops(offlineshops)
	if !ok {
		return c.SendString("Error in findOneOfflineShops operation")
	}
	return c.SendString("OfflineShops Found Succesfully")
}

// getManyOfflineShopsHandler обрабатывает HTTP GET запросы на /offlineshops
// @Summary Получает список офлайн магазинов OfflineShops
// @Description Возвращает список всех офлайн магазинов OfflineShops
// @Tags OfflineShops
// @Produce json
// @Success 200 {object} ResponseHTTP{data=db.OfflineShops}
// @Router /offlineshops [get]
func getManyOfflineShopsHandler(c *fiber.Ctx) error {
	var offlineshops db.OfflineShops
	ok := db.FindManyOfflineShops(offlineshops)
	if !ok {
		return c.SendString("Error in findManyOfflineShops operation")
	}
	return c.SendString("OfflineShops Found Succesfully")
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
