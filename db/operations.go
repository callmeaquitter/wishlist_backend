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

func UpdateSelection(selection Selection) bool {
	result := Database.Model(&selection).Updates(Selection{Name: selection.Name, Description: selection.Description})
	if result.Error != nil {
		fmt.Println("Error in UpdateSelection", result.Error)
		return false
	}
	return true
}

func FindManySelection(selection Selection) (bool, Selection) {
	result := Database.Find(&selection)
	if result.Error != nil {
		fmt.Println("Error in FindManySelection", result.Error)
		return false, selection
	}
	return true, selection
}

func FindOneSelection(selection Selection) bool {
	result := Database.Take(&selection)
	if result.Error != nil {
		fmt.Println("Error in FindOneSelection", result.Error)
		return false
	}
	return true
}

func DeleteSelection(id string) bool {
	var selection Selection
	result := Database.Delete(&selection, id)
	if result.Error != nil {
		fmt.Println("Error in DeleteSelection", result.Error)
		return false
	}
	return true
}

func CreateGiftToSelection(giftToSelection GiftToSelection) bool {
	result := Database.Create(&giftToSelection)
	if result.Error != nil {
		fmt.Println("Error in CreateGiftToSelection", result.Error)
		return false
	}
	return true
}

func UpdateGiftToSelection(giftToSelection GiftToSelection) bool {
	result := Database.Save(&giftToSelection)
	if result.Error != nil {
		fmt.Println("Error in updateGiftToSelection", result.Error)
		return false
	}
	return true
}

func FindGiftToSelection(giftToSelection GiftToSelection) bool {
	result := Database.Find(&giftToSelection)
	if result.Error != nil {
		fmt.Println("Error in findGiftToSelection", result.Error)
		return false
	}
	return true
}

func DeleteGiftToSelection(SelectionID, GiftID string) bool {
	result := Database.Delete(GiftToSelection{SelectionID: SelectionID, GiftID: GiftID})
	if result.Error != nil {
		fmt.Println("Error in deleteGiftToSelection", result.Error)
		return false
	}
	return true
}

func CreateSelectionCategory(selectionCategory SelectionCategory) bool {
	result := Database.Create(&selectionCategory)
	if result.Error != nil {
		fmt.Println("Error in CreateSelectionCategory", result.Error)
		return false
	}
	return true
}

func UpdatedSelectionCategory(selectionCategory SelectionCategory) bool {
	result := Database.Save(&selectionCategory)
	if result.Error != nil {
		fmt.Println("Error in updateSelection", result.Error)
		return false
	}
	return true
}

func FindSelectionCategory(selectionCategory SelectionCategory) bool {
	result := Database.Find(&selectionCategory)
	if result.Error != nil {
		fmt.Println("Error in findSelection", result.Error)
		return false
	}
	return true
}

func DeleteSelectionCategory(id string) bool {
	result := Database.Delete(SelectionCategory{ID: id})
	if result.Error != nil {
		fmt.Println("Error in deleteSelection", result.Error)
		return false
	}
	return true
}

func CreateLikeToSelection(likeToSelection LikeToSelection) bool {
	result := Database.Create(&likeToSelection)
	if result.Error != nil {
		fmt.Println("Error in CreateLikeToSelection", result.Error)
		return false
	}
	return true
}

func GetLikesCountToSelection(selectionID string) int {
	var count int64
	result := Database.Model(&LikeToSelection{}).Where("selection_id = ?", selectionID).Count(&count)
	if result.Error != nil {
		fmt.Println("Error in GetLikesCountToSelection", result.Error)
		return -1
	}
	return int(count)
}

func DeleteLikeToSelection(UserID, SelectionID string) bool {
	result := Database.Delete(LikeToSelection{UserID: UserID, SelectionID: SelectionID})
	if result.Error != nil {
		fmt.Println("Error in DeleteLikeToSelection", result.Error)
		return false
	}
	return true
}

func CreateCommentToSelection(commentToSelection CommentToSelection) bool {
	result := Database.Create(&commentToSelection)
	if result.Error != nil {
		fmt.Println("Error in CreateCommentToSelection", result.Error)
		return false
	}
	return true
}

func GetCommentsToSelection(id string) ([]CommentToSelection, bool) {
	var comments []CommentToSelection
	result := Database.Where("id = ?", id).Find(&comments)
	if result.Error != nil {
		fmt.Println("Error in GetCommentsToSelection", result.Error)
		return nil, false
	}
	return comments, true
}

func UpdateCommentToSelection(commentToSelection CommentToSelection) bool {
	result := Database.Save(&commentToSelection)
	if result.Error != nil {
		fmt.Println("Error in UpdateCommentToSelection", result.Error)
		return false
	}
	return true
}

func DeleteCommentToSelection(id string) bool {
	result := Database.Delete(CommentToSelection{ID: id})
	if result.Error != nil {
		fmt.Println("Error in DeleteCommentToSelection", result.Error)
		return false
	}
	return true
}
