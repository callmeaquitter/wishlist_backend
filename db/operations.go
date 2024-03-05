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
	result := Database.Model(&quest)
	if result.Error != nil {
		fmt.Println("Error in createQuest", result.Error)
		return false
	}
	return true
}

func DeleteQuest(id string) bool {
	result := Database.Delete(Quest{ID: id})
	if result.Error != nil {
		fmt.Println("Error in deleteQuest", result.Error)
		return false
	}
	return true
}

func UpdateQuest(quest Quest) bool {
	result := Database.Create(&quest).Update("name", "hello") //TODO
	if result.Error != nil {
		fmt.Println("Error in updateQuest", result.Error)
		return false
	}
	return true
}

//Subquest

func CreateSubquest(subquest Subquest) bool {
	result := Database.Model(&subquest)
	if result.Error != nil {
		fmt.Println("Error in createSubquest", result.Error)
		return false
	}
	return true
}

func DeleteSubquest(id string) bool {
	result := Database.Delete(Subquest{ID: id})
	if result.Error != nil {
		fmt.Println("Error in deleteSubquest", result.Error)
		return false
	}
	return true
}

//Tasks

func CreateTasks(tasks Tasks) bool {
	result := Database.Model(&tasks)
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
	result := Database.Create(&tasks).Update("name", "hello") //TODO
	if result.Error != nil {
		fmt.Println("Error in updateTasks", result.Error)
		return false
	}
	return true
}

func FindManyTasks(tasks Tasks) bool {
	result := Database.Find(&tasks)
	if result.Error != nil {
		fmt.Println("Error in findManyTasks", result.Error)
		return false
	}
	return true
}

func FindOneTasks(tasks Tasks) bool {
	result := Database.Take(&tasks)
	if result.Error != nil {
		fmt.Println("Error in findOneTasks", result.Error)
		return false
	}
	return true
}


таски созд наход удал редак
саб не редакт 