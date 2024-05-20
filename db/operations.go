package db

import (
	"fmt"

	_ "wishlist/docs"
)

func CreateGift(gift Gift) bool {
	result := Database.Create(&gift)
	if result.Error != nil {
		fmt.Println("Error in createGift", result.Error)
		return false
	}
	return true
}

func DeleteGift(id string) bool {
	result := Database.Delete(Gift{ID: id})
	if result.Error != nil {
		fmt.Println("Error in deleteGift", result.Error)
		return false
	}
	return true
}

func FindManyGift(gift Gift) bool {
	result := Database.Find(&gift)
	if result.Error != nil {
		fmt.Println("Error in findManyGift", result.Error)
		return false
	}
	return true
}

func FindOneGift(gift Gift) bool {
	result := Database.Take(&gift)
	if result.Error != nil {
		fmt.Println("Error in findOneGift", result.Error)
		return false
	}
	return true
}

func UpdateGift(gift Gift) bool {
	result := Database.Create(&gift).Update("name", "hello") //TODO
	if result.Error != nil {
		fmt.Println("Error in updateGift", result.Error)
		return false
	}
	return true
}

//Quest

func CreateQuest(quest Quest) bool {
	result := Database.Create(&quest)
	if result.Error != nil {
		fmt.Println("Error in createQuest", result.Error)
		return false
	}
	return true
}

func FindManyQuest() ([]Quest, bool) {
	var quests []Quest
	result := Database.Find(&quests)
	if result.Error != nil {
		fmt.Println("Error in findManyQuest", result.Error)
		return nil, false
	}
	return quests, true
}

func DeleteQuest(id string) bool {
	result := Database.Delete(Quest{ID: id})
	if result.Error != nil {
		fmt.Println("Error in deleteQuest", result.Error)
		return false
	}
	return true
}

//Subquest

func CreateSubquest(subquest Subquest) bool {
	result := Database.Create(&subquest)
	if result.Error != nil {
		fmt.Println("Error in createSubquest", result.Error)
		return false
	}
	return true
}

func FindManySubquest() ([]Subquest, bool) {
	var subquests []Subquest
	result := Database.Find(&subquests)
	if result.Error != nil {
		fmt.Println("Error in findManySubquest", result.Error)
		return nil, false
	}
	return subquests, true
}

func FindOneSubquest(subquestId string) (Subquest, bool) {
	var subquest Subquest
	result := Database.Where(&Subquest{ID: subquestId}).Take(&subquest)
	if result.Error != nil {
		fmt.Println("Error in findOneSubquest", result.Error)
		return subquest, false
	}
	return subquest, true
}

func DeleteSubquest(id string) bool {
	result := Database.Delete(Subquest{ID: id})
	if result.Error != nil {
		fmt.Println("Error in deleteSubquest", result.Error)
		return false
	}
	return true
}

func UpdateSubquest(subquest Subquest) bool {
	result := Database.Model(&subquest).Updates(Subquest{TaskID: subquest.TaskID, Reward: subquest.Reward, IsDone: subquest.IsDone})
	if result.Error != nil {
		fmt.Println("Error in updateSubquest", result.Error)
		return false
	}
	return true
}

//Tasks

func CreateTasks(tasks Tasks) bool {
	result := Database.Create(&tasks)
	if result.Error != nil {
		fmt.Println("Error in createTasks", result.Error)
		return false
	}
	return true
}

func DeleteTasks(id string) bool {
	result := Database.Delete(Tasks{ID: id})
	if result.Error != nil {
		fmt.Println("Error in deleteTasks", result.Error)
		return false
	}
	return true
}

func UpdateTasks(tasks Tasks) bool {
	result := Database.Model(&tasks).Updates(Tasks{Name: tasks.Name, Description: tasks.Description})
	if result.Error != nil {
		fmt.Println("Error in updateTasks", result.Error)
		return false
	}
	return true
}

func FindManyTasks() ([]Tasks, bool) {
	var tasks []Tasks
	result := Database.Find(&tasks)
	if result.Error != nil {
		fmt.Println("Error in findManyTasks", result.Error)
		return nil, false
	}
	return tasks, true
}

func FindOneTasks(taskId string) (Tasks, bool) {
	var task Tasks
	result := Database.Where(&Tasks{ID: taskId}).Take(&task)
	if result.Error != nil {
		fmt.Println("Error in findOneTasks", result.Error)
		return task, false
	}
	return task, true
}

//OfflineShops

func CreateOfflineShops(offlineshops OfflineShops) bool {
	result := Database.Create(&offlineshops)
	if result.Error != nil {
		fmt.Println("Error in createOfflineShops", result.Error)
		return false
	}
	return true
}

func UpdateOfflineShops(offlineshops OfflineShops) bool {
	result := Database.Model(&offlineshops).Updates(OfflineShops{Name: offlineshops.Name, Location: offlineshops.Location})
	if result.Error != nil {
		fmt.Println("Error in updateOfflineShops", result.Error)
		return false
	}
	return true
}

func FindManyOfflineShops(offlineshops OfflineShops) bool {
	result := Database.Find(&offlineshops)
	if result.Error != nil {
		fmt.Println("Error in findManyOfflineShops", result.Error)
		return false
	}
	return true
}

func FindOneOfflineShops(offlineshops OfflineShops) bool {
	result := Database.Take(&offlineshops)
	if result.Error != nil {
		fmt.Println("Error in findOneOfflineShops", result.Error)
		return false
	}
	return true
}

func DeleteOfflineShops(id string) bool {
	result := Database.Delete(OfflineShops{ID: id})
	if result.Error != nil {
		fmt.Println("Error in deleteOfflineShops", result.Error)
		return false
	}
	return true
}

