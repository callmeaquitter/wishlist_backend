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

// createQuestHandler обрабатывает HTTP POST запросы на /quest/create.
// @Summary Создает новый Quest
// @Description Принимает JSON тело запроса с полями Quest и создает новый Quest
// @Tags Quest
// @Accept json
// @Produce json
// @Param Quest body db.Quest true "Create Quest"
// @Success 200 {object} ResponseHTTP{data=db.Quest}
// @Failure 400 {object} ResponseHTTP{}
// @Router /quest/create [post]
func createQuestHandler(c *fiber.Ctx) error {
	var quest db.Quest
	if err := c.BodyParser(&quest); err != nil {
		return c.SendString(err.Error())
	}

	quest.ID = "quest_" + xid.New().String()

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
	var quest db.Quest
	ok := db.UpdateQuest(quest)
	if !ok {
		return c.SendString("Error in updateQuest operation")
	}
	return c.SendString("Quest updated Succesfully")
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
	var quest db.Quest
	ok := db.FindOneQuest(quest)
	if !ok {
		return c.SendString("Error in findOneQuest operation")
	}
	return c.SendString("Quest Found Succesfully")
}

// getManyQuestHandler обрабатывает HTTP GET запросы на /quest/getmany.
// @Summary Получает список квестов Quest
// @Description Возвращает список всех квестов Quest
// @Tags Quest
// @Produce json
// @Success 200 {object} ResponseHTTP{data=db.Quest}
// @Router /quest/getmany [get]
func getManyQuestHandler(c *fiber.Ctx) error {
	var quest db.Quest
	ok := db.FindManyQuest(quest)
	if !ok {
		return c.SendString("Error in findManyQuest operation")
	}
	return c.SendString("Quest Found Succesfully")
}

// deleteQuestHandler обрабатывает HTTP DELETE запросы на /quest/delete/{id}.
// @Summary Удаляет существующий Quest по ID
// @Description Принимает ID квеста в URL и удаляет соответствующий квест
// @Tags Quest
// @Param id path int true "Quest ID"
// @Success 200 {string} string "Quest deleted successfully"
// @Failure 404 {string} string "Quest not found"
// @Router /quest/delete/{id} [delete]
func deleteQuestHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	ok := db.DeleteQuest(id)
	if !ok {
		return c.SendString("Error in deleteQuest operation")
	}
	return c.SendString("Quest deleted successfully")
}

//Subquest

// createSubquestHandler обрабатывает HTTP POST запросы на /subquest/create.
// @Summary Создает новый Subquest
// @Description Принимает JSON тело запроса с полями Subquest и создает новый Subquest
// @Tags Subquest
// @Accept json
// @Produce json
// @Param Subquest body db.Subquest true "Create Subquest"
// @Success 200 {object} ResponseHTTP{data=db.Subquest}
// @Failure 400 {object} ResponseHTTP{}
// @Router /subquest/create [post]
func createSubquestHandler(c *fiber.Ctx) error {
	var subquest db.Subquest
	if err := c.BodyParser(&subquest); err != nil {
		return c.SendString(err.Error())
	}

	subquest.ID = "subquest_" + xid.New().String()

	ok := db.CreateSubquest(subquest)
	if !ok {
		return c.SendString("Error in createSubquest operation")
	}

	return c.JSON(subquest)
}

// getManySubquestHandler обрабатывает HTTP GET запросы на /subquest/getmany.
// @Summary Получает список Subquest
// @Description Возвращает список всех подзаданий (Subquest)
// @Tags Subquest
// @Produce json
// @Success 200 {object} ResponseHTTP{data=db.Subquest}
// @Router /subquest/getmany [get]
func getManySubquestHandler(c *fiber.Ctx) error {
	var subquest db.Subquest
	ok := db.FindManySubquest(subquest)
	if !ok {
		return c.SendString("Error in findManySubquest operation")
	}
	return c.SendString("Subquest Found Succesfully")
}

// getOneSubquestHandler обрабатывает HTTP GET запросы на /subquest/getone/{id}.
// @Summary Получает одно Subquest по ID
// @Description Возвращает информацию о конкретном подзадании (Subquest) по его ID
// @Tags Subquest
// @Param id path int true "Subquest ID"
// @Produce json
// @Success 200 {object} ResponseHTTP{data=db.Subquest}
// @Failure 404 {string} string "Subquest not found"
// @Router /subquest/getone/{id} [get]
func getOneSubquestHandler(c *fiber.Ctx) error {
	var subquest db.Subquest
	ok := db.FindOneSubquest(subquest)
	if !ok {
		return c.SendString("Error in findOneSubquest operation")
	}
	return c.SendString("Subquest Found Succesfully")
}

// deleteSubquestHandler обрабатывает HTTP DELETE запросы на /subquest/delete/{id}.
// @Summary Удаляет существующий Subquest по ID
// @Description Принимает ID подзадания в URL и удаляет соответствующее подзадание
// @Tags Subquest
// @Param id path int true "Subquest ID"
// @Success 200 {string} string "Subquest deleted successfully"
// @Failure 404 {string} string "Subquest not found"
// @Router /subquest/delete/{id} [delete]
func deleteSubquestHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	ok := db.DeleteSubquest(id)
	if !ok {
		return c.SendString("Error in deleteSubquest operation")
	}
	return c.SendString("Subquest deleted successfully")
}

//Tasks

// createTasksHandler обрабатывает HTTP POST запросы на /tasks/create.
// @Summary Создает новое задание Tasks
// @Description Принимает JSON тело запроса с полями Tasks и создает новое задание
// @Tags Tasks
// @Accept json
// @Produce json
// @Param Tasks body db.Tasks true "Create Tasks"
// @Success 200 {object} ResponseHTTP{data=db.Tasks}
// @Failure 400 {object} ResponseHTTP{}
// @Router /tasks/create [post]
func createTasksHandler(c *fiber.Ctx) error {
	var tasks db.Tasks
	if err := c.BodyParser(&tasks); err != nil {
		return c.SendString(err.Error())
	}

	if tasks.Name == "" {
		return c.SendString("Name is required")
	}

	tasks.ID = "tasks_" + xid.New().String()

	ok := db.CreateTasks(tasks)
	if !ok {
		return c.SendString("Error in createTasks operation")
	}

	return c.JSON(tasks)
}

// updateTasksHandler обрабатывает HTTP PUT запросы на /tasks/update.
// @Summary Обновляет существующее задание Tasks
// @Description Принимает JSON тело запроса с обновленными полями Tasks и обновляет существующее задание
// @Tags Tasks
// @Accept json
// @Produce json
// @Param Tasks body db.Tasks true "Update Tasks"
// @Success 200 {object} ResponseHTTP{data=db.Tasks}
// @Failure 400 {object} ResponseHTTP{}
// @Router /tasks/update [put]
func updateTasksHandler(c *fiber.Ctx) error {
	var tasks db.Tasks
	ok := db.UpdateTasks(tasks)
	if !ok {
		return c.SendString("Error in updateTasks operation")
	}
	return c.SendString("Tasks updated Succesfully")
}

// getOneTasksHandler обрабатывает HTTP GET запросы на /tasks/getone/{id}.
// @Summary Получает одно задание Tasks по ID
// @Description Возвращает информацию о конкретном задании Tasks по его ID
// @Tags Tasks
// @Param id path int true "Tasks ID"
// @Produce json
// @Success 200 {object} ResponseHTTP{data=db.Tasks}
// @Failure 404 {string} string "Tasks not found"
// @Router /tasks/getone/{id} [get]
func getOneTasksHandler(c *fiber.Ctx) error {
	var tasks db.Tasks
	ok := db.FindOneTasks(tasks)
	if !ok {
		return c.SendString("Error in findOneTasks operation")
	}
	return c.SendString("Tasks Found Succesfully")
}

// getManyTasksHandler обрабатывает HTTP GET запросы на /tasks/getmany.
// @Summary Получает список заданий Tasks
// @Description Возвращает список всех заданий Tasks
// @Tags Tasks
// @Produce json
// @Success 200 {object} ResponseHTTP{data=db.Tasks}
// @Router /tasks/getmany [get]
func getManyTasksHandler(c *fiber.Ctx) error {
	var tasks db.Tasks
	ok := db.FindManyTasks(tasks)
	if !ok {
		return c.SendString("Error in findManyTasks operation")
	}
	return c.SendString("Tasks Found Succesfully")
}

// deleteTasksHandler обрабатывает HTTP DELETE запросы на /tasks/delete/{id}.
// @Summary Удаляет существующее задание Tasks по ID
// @Description Принимает ID задания в URL и удаляет соответствующее задание
// @Tags Tasks
// @Param id path int true "Tasks ID"
// @Success 200 {string} string "Tasks deleted successfully"
// @Failure 404 {string} string "Tasks not found"
// @Router /tasks/delete/{id} [delete]
func deleteTasksHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	ok := db.DeleteTasks(id)
	if !ok {
		return c.SendString("Error in deleteTasks operation")
	}
	return c.SendString("Tasks deleted successfully")
}

//OfflineShops

// createOfflineShopHandler обрабатывает HTTP POST запросы на /offlineshop/create.
// @Summary Создает новый Offline Shop
// @Description Принимает JSON тело запроса с полями Offline Shop и создает новый Offline Shop
// @Tags OfflineShops
// @Accept json
// @Produce json
// @Param OfflineShop body db.OfflineShops true "Create Offline Shop"
// @Success 200 {object} ResponseHTTP{data=db.OfflineShops}
// @Failure 400 {object} ResponseHTTP{}
// @Router /offlineshop/create [post]
func createOfflineShopsHandler(c *fiber.Ctx) error {
	var offlineshops db.OfflineShops
	if err := c.BodyParser(&offlineshops); err != nil {
		return c.SendString(err.Error())
	}

	if offlineshops.Name == "" {
		return c.SendString("Name is required")
	}
	if offlineshops.Location == "" {
		return c.SendString("Location is required")
	}

	offlineshops.ID = "offlineshops_" + xid.New().String()

	ok := db.CreateOfflineShops(offlineshops)
	if !ok {
		return c.SendString("Error in createOfflineShops operation")
	}

	return c.JSON(offlineshops)
}

// updateOfflineShopHandler обрабатывает HTTP PUT запросы на /offlineshop/update/{id}.
// @Summary Обновляет существующий Offline Shop по ID
// @Description Принимает JSON тело запроса с обновленными полями Offline Shop и обновляет существующий Offline Shop по его ID
// @Tags OfflineShops
// @Accept json
// @Produce json
// @Param id path string true "Offline Shop ID"
// @Param OfflineShop body db.OfflineShops true "Update Offline Shop"
// @Success 200 {object} ResponseHTTP{data=db.OfflineShops}
// @Failure 400 {object} ResponseHTTP{}
// @Router /offlineshop/update/{id} [put]
func updateOfflineShopsHandler(c *fiber.Ctx) error {
	var offlineshops db.OfflineShops
	ok := db.UpdateOfflineShops(offlineshops)
	if !ok {
		return c.SendString("Error in updateOfflineShops operation")
	}
	return c.SendString("OfflineShops updated Succesfully")
}

// getOneOfflineShopsHandler обрабатывает HTTP GET запросы на /offlineshops/getone/{id}.
// @Summary Получает один офлайн магазин OfflineShops по ID
// @Description Возвращает информацию о конкретном офлайн магазине OfflineShops по его ID
// @Tags OfflineShops
// @Param id path int true "OfflineShops ID"
// @Produce json
// @Success 200 {object} ResponseHTTP{data=db.OfflineShops}
// @Failure 404 {string} string "OfflineShops not found"
// @Router /offlineshops/getone/{id} [get]
func getOneOfflineShopsHandler(c *fiber.Ctx) error {
	var offlineshops db.OfflineShops
	ok := db.FindOneOfflineShops(offlineshops)
	if !ok {
		return c.SendString("Error in findOneOfflineShops operation")
	}
	return c.SendString("OfflineShops Found Succesfully")
}

// getManyOfflineShopsHandler обрабатывает HTTP GET запросы на /offlineshops/getmany.
// @Summary Получает список офлайн магазинов OfflineShops
// @Description Возвращает список всех офлайн магазинов OfflineShops
// @Tags OfflineShops
// @Produce json
// @Success 200 {object} ResponseHTTP{data=db.OfflineShops}
// @Router /offlineshops/getmany [get]
func getManyOfflineShopsHandler(c *fiber.Ctx) error {
	var offlineshops db.OfflineShops
	ok := db.FindManyOfflineShops(offlineshops)
	if !ok {
		return c.SendString("Error in findManyOfflineShops operation")
	}
	return c.SendString("OfflineShops Found Succesfully")
}

// deleteOfflineShopHandler обрабатывает HTTP DELETE запросы на /offlineshop/delete/{id}.
// @Summary Удаляет существующий Offline Shop по ID
// @Description Принимает ID офлайн магазина в URL и удаляет соответствующий офлайн магазин
// @Tags OfflineShops
// @Param id path string true "Offline Shop ID"
// @Success 200 {string} string "Offline Shop deleted successfully"
// @Failure 404 {string} string "Offline Shop not found"
// @Router /offlineshop/delete/{id} [delete]
func deleteOfflineShopsHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	ok := db.DeleteOfflineShops(id)
	if !ok {
		return c.SendString("Error in deleteOfflineShops operation")
	}
	return c.SendString("OfflineShops deleted successfully")
}
