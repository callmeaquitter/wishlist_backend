package db

import (
	"fmt"
	_ "wishlist/docs"
)

type IError struct {
	Field string
	Tag   string
	Value string
}

// var Validator = validator.New()

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

func CreateWishlist(wishlist UserWishlist) bool {
	result := Database.Create(&wishlist)
	if result.Error != nil {
		fmt.Println("Error in CreateWishlist", result.Error)
		return false
	}
	return true
}

func FindManyWishlists(wishlists UserWishlist) bool {
	result := Database.Find(&wishlists)
	if result.Error != nil {
		fmt.Println("Error in FindManyWishlists", result.Error)
		return false
	}
	return true
}

func FindWishlistByName(name string) bool {
	var wishlist UserWishlist
	result := Database.Where(&wishlist, name).Take(&wishlist)
	if result.Error != nil {
		fmt.Println("Error in FindWishlistByName", result.Error)
		return false
	}
	return true
}

func UpdateWishlist(wishlistID, wishlistName string) bool {
	var wishlist UserWishlist
	if err := Database.Where(&wishlist, wishlistID); err != nil {
		fmt.Println("Error in finding wishlist for update", err)
		return false
	}

	result := Database.Update("name", wishlistName)
	if result.Error != nil {
		fmt.Println("Error in UpdateWishlist", result.Error)
		return false
	}

	return true
}

func AddWish(wishlistID, wishID string) bool {
	var wish Wishes
	result := Database.Select(wishlistID, wishID).Create(&wish)
	if result.Error != nil {
		fmt.Println("Error in CreateWish", result.Error)
		return false
	}

	return true
}

func GetManyWishesInWishlist(wishlistID string) bool {
	var wishes Wishes
	result := Database.Where(&wishes, wishlistID).Find(&wishes)
	if result.Error != nil {
		fmt.Println("Error in GetManyWishes", result.Error)
		return false
	}
	return true
}

// func GetOneWish(wish Wishes) bool {
// 	result := Database.Take(&wish)
// 	if result != nil {
// 		fmt.Println("Error in GetOneWish")
// 		return false
// 	}
// 	return true
// }

func DeleteWish(wishlistID, GiftID string) bool {
	var wish Wishes
	result := Database.Where(&wish, wishlistID, GiftID).Delete(&wish)
	if result.Error != nil {
		fmt.Println("Error in DeleteWish", result.Error)
		return false
	}
	return true
}

func DeleteWishlist(wishlistID, GiftID, userID string) bool {
	var wishlist UserWishlist
	var wishes Wishes

	if err := Database.Where(&wishlist, wishlistID, userID).First(&wishlist); err != nil {
		fmt.Println("Error in finding wishlist for deleting", err)
		return false
	}

	if err := Database.Where(&wishes, wishlistID, GiftID).Delete(&wishes); err != nil {
		fmt.Println("Error in deleting wishes", err)
		return false
	}

	result := Database.Where(&wishlist, wishlistID, userID).Delete(&wishlist)
	if result.Error != nil {
		fmt.Println("Error in DeleteWishlist", result.Error)
	}
	
	return true

}
