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
	result := Database.Model(&gift).Update("name", "hello")
	if result.Error != nil {
		fmt.Println("Error in updateGift", result.Error)
		return false
	}
	return true
}


func CreateSelection(selection Selection) bool {
	result := Database.Create(&selection)
	if result.Error != nil {
		fmt.Println("Error in CreateSelection", result.Error)
		return false
	}
	return true
}

// func ReadSelection(id int) (*Selection, bool) {
// 	var selection Selection
// 	result := Database.First(&selection, id)
// 	if result.Error != nil {
// 		fmt.Println("Error in ReadSelection", result.Error)
// 		return nil, false
// 	}
// 	return &selection, true
// }

func UpdateSelection(selection Selection) bool {
	result := Database.Model(&selection).Updates(Selection{Name: selection.Name, Description: selection.Description, IsGenerated: selection.IsGenerated})
	if result.Error != nil {
		fmt.Println("Error in UpdateSelection", result.Error)
		return false
	}
	return true
}

func DeleteSelection(id int) bool {
	var selection Selection
	result := Database.Delete(&selection, id)
	if result.Error != nil {
		fmt.Println("Error in DeleteSelection", result.Error)
		return false
	}
	return true
}

