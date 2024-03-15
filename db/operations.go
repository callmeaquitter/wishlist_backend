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

func FindManyWishlists(wishlists []UserWishlist) bool {
	result := Database.Find(&wishlists)
	if result.Error != nil {
		fmt.Println("Error in FindManyWishlists", result.Error)
		return false
	}
	return true
}

func FindWishlistByName(name string) bool {
	var wishlist UserWishlist
	result := Database.Where(&UserWishlist{Name: name}).Take(&wishlist)
	if result.Error != nil {
		fmt.Println("Error in FindWishlistByName", result.Error)
		return false
	}
	return true
}

func UpdateWishlist(wishlistID, wishlistName string) bool {
	var wishlist UserWishlist

	err := Database.Where(&UserWishlist{ID: wishlistID}).First(&wishlist)
	if err.Error != nil {
		fmt.Println("Ошибка при поиске списка желаемого для обновления", err.Error)
		return false
	}

	result := Database.Model(&wishlist).Update("name", wishlistName)
	if result.Error != nil {
		fmt.Println("Ошибка в UpdateWishlist", result.Error)
		return false
	}

	return true
}

func AddWish(wishlistID, giftID string) bool {
	wish := Wishes{GiftID: giftID, WishlistID: wishlistID}
	result := Database.Create(wish)
	if result.Error != nil {
		fmt.Println("Error in CreateWish", result.Error)
		return false
	}

	return true
}

func GetManyWishesInWishlist(wishlistID string) bool {
	var wishes []Wishes
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

func DeleteWish(wishlistID, giftID string) bool {
	var wish Wishes
	result := Database.Where(&Wishes{GiftID: giftID, WishlistID: wishlistID}).Delete(&wish)
	if result.Error != nil {
		fmt.Println("Error in DeleteWish", result.Error)
		return false
	}
	return true
}

func DeleteWishlist(wishlistID, giftID, userID string) bool {
	var wishlist UserWishlist
	var wishes Wishes

	if err := Database.Where(&UserWishlist{ID: wishlistID, UserID: userID}).First(&wishlist); err != nil {
		fmt.Println("Error in finding wishlist for deleting", err)
		return false
	}

	if err := Database.Where(&Wishes{WishlistID: wishlistID, GiftID: giftID}).Find(&wishes); err != nil {
		fmt.Println("Error in deleting wishes", err)
		return false
	}

	if err := Database.Delete(&wishes); err.Error != nil {
		fmt.Println("Error in deleting wish", err.Error)
		return false
	}

	result := Database.Delete(&wishlist)
	if result.Error != nil {
		fmt.Println("Error in DeleteWishlist", result.Error)
		return false
	}

	return true

	// // Находим wishlist
	// if err := Database.Where(&UserWishlist{ID: wishlistID, UserID: userID}).First(&wishlist); err.Error != nil {
	// 	fmt.Println("Error in finding wishlist for deleting", err.Error)
	// 	return false
	// }

	// // Находим соответствующую запись wishes
	// if err := Database.Where(&Wishes{WishlistID: wishlistID, GiftID: giftID}).First(&wishes); err.Error != nil {
	// 	fmt.Println("Error in finding wish for deletion", err.Error)
	// 	return false
	// }

	// // Удаляем запись wishes
	// if err := Database.Delete(&wishes); err.Error != nil {
	// 	fmt.Println("Error in deleting wish", err.Error)
	// 	return false
	// }

	// // Удаляем wishlist
	// if err := Database.Delete(&wishlist); err.Error != nil {
	// 	fmt.Println("Error in deleting wishlist", err.Error)
	// 	return false
	// }

	// // Возвращаем успешное завершение операции
	// return true

}

func CreateUser(user User) bool {
	result := Database.Create(&user)
	if result.Error != nil {
		fmt.Println("Error in CreateUser", result.Error)
		return false
	}
	return true
}

func FindUser(login, password string) (User, bool) {
	var user User
	result := Database.Where(&User{Login: login, Password: password}).First(&user)
	if result.Error != nil {
		fmt.Println("Error in FindUser", result.Error)
		return user, false
	}
	return user, true
}

func CreateSession(session Session) bool {
	result := Database.Create(&session)
	if result.Error != nil {
		fmt.Println("Error in CreateSession", result.Error)
		return false
	}
	return true
}

func FindSession(sessionID string) (Session, bool) {
	var session Session
	result := Database.Where(&Session{ID: sessionID}).First(&session)
	if result.Error != nil {
		fmt.Println("Error in FindSession", result.Error)
		return session, false
	}
	return session, true
}
