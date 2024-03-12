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
	if result != nil {
		fmt.Println("Error in FindManyWishlists", result.Error)
		return false
	}
	return true
}

func FindWishlistByName(wishlist UserWishlist) bool {
	result := Database.Take(&wishlist)
	if result != nil {
		fmt.Println("Error in FindWishlistByName", result.Error)
		return false
	}
	return true
}

func CreateWish(wishes Wishes, wishlistID string) bool {
	user
	result := Database.Create(&wishes)
	if result != nil {
		fmt.Println("Error in CreateWish", result.Error)
	}
	return true
}

func GetManyWishes(wishes Wishes) bool {
	result := Database.Find(&wishes)
	if result != nil {
		fmt.Println("Error in GetManyWishes")
		return false
	}
	return true
}

func GetOneWish(wish Wishes) bool {
	result := Database.Take(&wish)
	if result != nil {
		fmt.Println("Error in GetOneWish")
		return false
	}
	return true
}

func DeleteWish(giftID, wishlistID string) bool {
	result := Database.Delete(Wishes{GiftID: giftID, WishlistID: wishlistID})
	if result != nil {
		fmt.Println("Error in DeleteWish")
		return false
	}
	return true
}
