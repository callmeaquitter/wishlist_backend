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



func UpdateGift(id string, updatedGift Gift) bool {
    
    existingGift := Gift{}
    result := Database.First(&existingGift, "id = ?", id)
    if result.Error != nil {
        fmt.Println("Error in finding gift:", result.Error)
        return false
    }

	if updatedGift.Name != "" {
        existingGift.Name = updatedGift.Name
    }
    if updatedGift.Price != 0 {
        existingGift.Price = updatedGift.Price
    }
    if updatedGift.Photo != "" {
        existingGift.Photo = updatedGift.Photo
    }
	if updatedGift.Description != "" {
        existingGift.Description = updatedGift.Description
    }
	if updatedGift.Link != "" {
        existingGift.Link = updatedGift.Link
    }

	if updatedGift.Category != "" {
        existingGift.Category = updatedGift.Category
    }
	result = Database.Save(&existingGift)
    if result.Error != nil {
        fmt.Println("Error updating the gift:", result.Error)
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

func DeleteBookedGift(giftID string) bool {
	var bookedGiftInWishlist BookedGiftInWishlist
	result := Database.Where(&BookedGiftInWishlist{GiftID: giftID}).Delete(&bookedGiftInWishlist)
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
		fmt.Println("Error in findManyUsersGift", result.Error)
		return bookedGiftInWishlist, false
	}
	return bookedGiftInWishlist, true
}

// gift categore
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

// gift rewiev
func CreateGiftReview(GiftReview GiftReview) bool {
	result := Database.Create(&GiftReview)
	if result.Error != nil {
		fmt.Println("Error in createGiftReview", result.Error)
		return false
	}
	return true
}

func DeleteGiftReview(id string) bool {
	result := Database.Delete(GiftReview{ID: id})
	if result.Error != nil {
		fmt.Println("Error in deleteGiftReview", result.Error)
		return false
	}
	return true
}
//получение review по его id
func GetGiftReviewByID(id string) (GiftReview, bool) {
	var giftReview GiftReview
	result := Database.First(&giftReview, "id = ?", id)
	if result.Error != nil {
		fmt.Println("Error in deleteBookedGift", result.Error)
		return GiftReview{}, false
	}
	return giftReview, true
}

func GetGiftReviewsByGiftID(giftID string) ([]GiftReview, bool) {
	var giftReviews []GiftReview
	result := Database.Where("gift_id = ?", giftID).Find(&giftReviews)
	if result.Error != nil {
		fmt.Println("Error in getGiftReviewsGiftID", result.Error)
        return nil, false
    }
    return giftReviews, true
}

func CalculateAverageMarkByGiftID(giftID string) (float32, bool){
	giftReviews, found := GetGiftReviewsByGiftID(giftID)
	if !found || len(giftReviews) == 0 {
        return 0.0, false
    }
	totalMarks := float32(0)
    for _, review := range giftReviews {
        totalMarks += review.Mark
    }
	averageMark := totalMarks / float32(len(giftReviews))
    return averageMark, true
}