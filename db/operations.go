package db

import (
	"fmt"
	// "github.com/shopspring/decimal"
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

func CreateWishlist(wishlist UserWishlist) bool {
	result := Database.Create(&wishlist)
	if result.Error != nil {
		fmt.Println("Error in CreateWishlist", result.Error)
		return false
	}
	return true
}

func FindManyWishlists(userID string) (UserWishlist, bool) {
	var wishlist UserWishlist
	result := Database.Where(&UserWishlist{UserID: userID}).Find(&wishlist)
	if result.Error != nil {
		fmt.Println("Error in FindManyWishlists", result.Error)
		return wishlist, false
	}
	return wishlist, true
}

func FindWishlistByName(name string) (UserWishlist, bool) {
	var wishlist UserWishlist
	result := Database.Where(&UserWishlist{Name: name}).Take(&wishlist)
	if result.Error != nil {
		fmt.Println("Error in FindWishlistByName", result.Error)
		return wishlist, false
	}
	return wishlist, true
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

func GetManyWishesInWishlist(wishlistID string) ([]Wishes, bool) {
	var wishes []Wishes
	result := Database.Where(&wishes, wishlistID).Find(&wishes)
	if result.Error != nil {
		fmt.Println("Error in GetManyWishes", result.Error)
		return wishes, false
	}
	return wishes, true
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

func DeleteGiftReview(id string) bool {
	result := Database.Delete(GiftReview{ID: id})
	if result.Error != nil {
		fmt.Println("Error in deleteGiftReview", result.Error)
		return false
	}
	return true
}

func CreateSeller(seller Seller) bool {
	result := Database.Create(&seller)
	if result.Error != nil {
		fmt.Println("Error in createSeller", result.Error)
		return false
	}
	return true
}

func FindSeller(login, password string) (Seller, bool) {
	var seller Seller
	result := Database.Where(&Seller{Login: login, Password: password}).First(&seller)
	if result.Error != nil {
		fmt.Println("Error in findOneSeller", result.Error)
		return seller, false
	}
	return seller, true
}

func CreateSellerSession(sellerSession SellerSession) bool {
	result := Database.Create(&sellerSession)
	if result.Error != nil {
		fmt.Println("Error in CreateSellerSession", result.Error)
		return false
	}
	return true
}

func FindSellerSession(sellerSessionID string) (SellerSession, bool) {
	var sellerSession SellerSession
	result := Database.Where(&SellerSession{ID: sellerSessionID}).First(&sellerSession)
	if result.Error != nil {
		fmt.Println("Error in FindSellerSession", result.Error)
		return sellerSession, false
	}
	return sellerSession, true
}

func CreateService(service Service) bool {
	result := Database.Create(&service)
	if result.Error != nil {
		fmt.Println("Error in createService")
		return false
	}
	return true
}

func FindManyService() ([]Service, bool) {
	var service []Service
	result := Database.Find(&service)
	if result.Error != nil {
		fmt.Println("Error in findManyService", result.Error)
		return service, false
	}
	return service, true
}

func FindOneService(serviceId string) (Service, bool) {
	var service Service
	result := Database.Take(&service, "id = ?", serviceId)
	if result.Error != nil {
		fmt.Println("Error in findOneService", result.Error)
		return service, false
	}
	return service, true
}

func UpdateService(service Service) bool {
	if service.Name != "string" {
		result := Database.Model(&service).Update("name", service.Name)
		if result.Error != nil {
			fmt.Println("Error in updateservice", result.Error)
			return false
		}
	}
	result := Database.Model(&service).Update("price", service.Price)
	if result.Error != nil {
		fmt.Println("Error in updateservice", result.Error)
		return false
	}
	if service.Location != "string" {
		result := Database.Model(&service).Update("location", service.Location)
		if result.Error != nil {
			fmt.Println("Error in updateservice", result.Error)
			return false
		}
	}
	if service.Photos != "string" {
		result := Database.Model(&service).Update("photos", service.Photos)
		if result.Error != nil {
			fmt.Println("Error in updateservice", result.Error)
			return false
		}
	}

	return true
}

func DeleteService(id string) bool {
	result := Database.Delete(Service{ServiceID: id})
	if result.Error != nil {
		fmt.Println("Error in deleteService", result.Error)
		return false
	}
	return true
}

func CreateSellerToService(sellerToService SellerToService) bool {
	result := Database.Create(&sellerToService)
	if result.Error != nil {
		fmt.Println("Error in CreateSellerToService")
		return false
	}
	return true
}

func FindManySellerToService() ([]SellerToService, bool) {
	var sellerToService []SellerToService
	result := Database.Find(&sellerToService)
	if result.Error != nil {
		fmt.Println("Error in findOneSellerToService", result.Error)
		return sellerToService, false
	}
	return sellerToService, true
}

func FindOneSellerToService(sellerId string) ([]SellerToService, bool) {
	var sellerToService []SellerToService
	result := Database.Find(&sellerToService, "seller_id = ?", sellerId)
	if result.Error != nil {
		fmt.Println("Error in findManySellerToService", result.Error)
		return sellerToService, false
	}
	return sellerToService, true
}

func DeleteSellerToService(serviceId string) bool { // Связь удаляется по услуге, т.к. оная удаляется чаще
	result := Database.Delete(SellerToService{}, "service_id = ?", serviceId)
	if result.Error != nil {
		fmt.Println("Error in deleteSellerToService", result.Error)
		return false
	}
	return true
}

func CreateServiceReview(serviceReview ServiceReview) bool {
	result := Database.Create(&serviceReview)
	if result.Error != nil {
		fmt.Println("Error in createServiceReview", result.Error)
		return false
	}
	return true
}

func FindManyServiceReview() ([]ServiceReview, bool) {
	var serviceReview []ServiceReview
	result := Database.Find(&serviceReview)
	if result.Error != nil {
		fmt.Println("Error in findManyServiceReview", result.Error)
		return serviceReview, false
	}
	return serviceReview, true
}

func FindOneServiceReview(serviceReviewId string) (ServiceReview, bool) {
	var serviceReview ServiceReview
	result := Database.Take(&serviceReview, "id = ?", serviceReviewId)
	if result.Error != nil {
		fmt.Println("Error in findOneServiceReview", result.Error)
		return serviceReview, false
	}
	return serviceReview, true
}

func FindSingleServiceReview(serviceId string) ([]ServiceReview, bool) {
	var serviceReview []ServiceReview
	result := Database.Find(&serviceReview, "service_id = ?", serviceId)
	if result.Error != nil {
		fmt.Println("Error in findSingleServiceReview", result.Error)
		return serviceReview, false
	}
	return serviceReview, true
}

// Не работает лол
func UpdateServiceReview(serviceReview ServiceReview) bool {
	result := Database.Model(&serviceReview).
		Update("service_id", serviceReview.ServiceID)
	if result.Error != nil {
		fmt.Println("Error in updateServiceReview", result.Error)
		return false
	}
	result = Database.Model(&serviceReview).
		Update("mark", serviceReview.Mark)
	if result.Error != nil {
		fmt.Println("Error in updateServiceReview", result.Error)
		return false
	}
	if serviceReview.Comment != "string" {
		result := Database.Model(&serviceReview).
			Update("comment", serviceReview.Comment)
		if result.Error != nil {
			fmt.Println("Error in updateServiceReview", result.Error)
			return false
		}
	}
	result = Database.Model(&serviceReview).
		Update("user_id", serviceReview.UserID)
	if result.Error != nil {
		fmt.Println("Error in updateServiceReview", result.Error)
		return false
	}
	result = Database.Model(&serviceReview).
		Update("update_date", serviceReview.UpdateDate)
	if result.Error != nil {
		fmt.Println("Error in updateServiceReview", result.Error)
		return false
	}

	return true
}

func DeleteServiceReview(id string) bool {
	result := Database.Delete(&ServiceReview{}, "id = ?", id)
	if result.Error != nil {
		fmt.Println("Error in deleteServiceReview", result.Error)
		return false
	}
	return true
}

// получение review по его id
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

func CalculateAverageMarkByGiftID(giftID string) (float32, bool) {
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

func FindManySelection() (bool, []Selection) {
	var selections []Selection
	result := Database.Find(&selections)
	if result.Error != nil {
		fmt.Println("Error in FindManySelection", result.Error)
		return false, selections
	}
	return true, selections
}

func FindOneSelection(selectionID, userID string) (Selection, bool) {
	var selection Selection
	result := Database.Where(&Selection{ID: selectionID, UserID: userID}).Take(&selection)
	if result.Error != nil {
		fmt.Println("Error in FindOneSelection", result.Error)
		return selection, false
	}
	return selection, true
}

func DeleteSelection(selectionID string) bool {
	result := Database.Where("id = ?", selectionID).Delete(&Selection{})
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

func FindManySelectionCategory() ([]SelectionCategory, bool) {
	var selectionCategory []SelectionCategory
	result := Database.Find(&selectionCategory)
	if result.Error != nil {
		fmt.Println("Error in findSelection", result.Error)
		return selectionCategory, false
	}
	return selectionCategory, true
}

func FindOneSelectionCategory(selectionCategoryID string) (SelectionCategory, bool) {
	var selectionCategory SelectionCategory
	result := Database.Where("id = ?", selectionCategoryID).Take(&selectionCategory)
	if result.Error != nil {
		fmt.Println("Error in findSelection", result.Error)
		return selectionCategory, false
	}
	return selectionCategory, true
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
	result := Database.Where("selection_id = ?", id).Find(&comments)
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

//Quest

func CreateQuest(quest Quest) bool {
	result := Database.Model(&quest)
	if result.Error != nil {
		fmt.Println("Error in createQuest", result.Error)
		return false
	}
	return true
}

func FindManyQuest(quest Quest) bool {
	result := Database.Find(&quest)
	if result.Error != nil {
		fmt.Println("Error in findManyQuest", result.Error)
		return false
	}
	return true
}

func FindOneQuest(quest Quest) bool {
	result := Database.Take(&quest)
	if result.Error != nil {
		fmt.Println("Error in findOneQuest", result.Error)
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
	result := Database.Model(&quest).Updates(map[string]interface{}{"subquest_id": quest.SubquestID, "user_id": quest.UserID, "is_done": quest.IsDone})
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

func FindManySubquest(subquest Subquest) bool {
	result := Database.Find(&subquest)
	if result.Error != nil {
		fmt.Println("Error in findManySubquest", result.Error)
		return false
	}
	return true
}

func FindOneSubquest(subquest Subquest) bool {
	result := Database.Take(&subquest)
	if result.Error != nil {
		fmt.Println("Error in findOneSubquest", result.Error)
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
	result := Database.Model(&tasks).Updates(map[string]interface{}{"name": tasks.Name, "description": tasks.Description})
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

//OfflineShops

func CreateOfflineShops(offlineshops OfflineShops) bool {
	result := Database.Model(&offlineshops)
	if result.Error != nil {
		fmt.Println("Error in createOfflineShops", result.Error)
		return false
	}
	return true
}

func UpdateOfflineShops(offlineshops OfflineShops) bool {
	result := Database.Model(&offlineshops).Updates(map[string]interface{}{"name": offlineshops.Name, "location": offlineshops.Location})
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
