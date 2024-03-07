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

func CreateBookedGift(BookedGiftInWishlist BookedGiftInWishlist) bool {
	result := Database.Create(&BookedGiftInWishlist)
	if result.Error != nil {
		fmt.Println("Error in createBookedGiftInWishlist", result.Error)
		return false
	}
	return true
}

func DeleteBookedGift(UserID, GiftID string) bool {
	result := Database.Delete(BookedGiftInWishlist{UserID: UserID, GiftID: GiftID})
	if result.Error != nil {
		fmt.Println("Error in deleteBookedGift", result.Error)
		return false
	}
	return true

}

func FindManyUsersGift(UserID string) ([]BookedGiftInWishlist, bool) {
	var bookedGiftInWishlist []BookedGiftInWishlist 
	result := Database.Find(&bookedGiftInWishlist, "user_id = ?", UserID)
	if result.Error != nil {
		fmt.Println("Error in deleteBookedGift", result.Error)
		return bookedGiftInWishlist, false
	}
	return bookedGiftInWishlist, true
}

func CreateGiftCategory(GiftCategory GiftCategory) bool {
	result := Database.Create(&GiftCategory)
	if result.Error != nil {
		fmt.Println("Error in CreateGiftCategory", result.Error)
		return false
	}
	return true
}

func DeleteGiftCategory(id string) bool {
	result := Database.Delete(GiftCategory{ID: id})
	if result.Error != nil {
		fmt.Println("Error in deleteGiftCategory", result.Error)
		return false
	}
	return true
}