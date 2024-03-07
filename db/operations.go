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

func CreateSeller(seller Seller) bool {
	result := Database.Create(&seller)
	if result.Error != nil {
		fmt.Println("Error in createSeller", result.Error)
		return false
	}
	return true
}

func FindManySeller(seller Seller) bool {
	result := Database.Find(&seller)
	if result.Error != nil {
		fmt.Println("Error in findManySeller", result.Error)
		return false
	}
	return true
}

func FindOneSeller(seller Seller) bool {
	result := Database.Take(&seller)
	if result.Error != nil {
		fmt.Println("Error in findOneSeller", result.Error)
		return false
	}
	return true
}

func UpdateSeller(seller Seller) bool {
	result := Database.Model(&seller).Update("name", "hello")
	if result.Error != nil {
		fmt.Println("Error in updateSeller", result.Error)
		return false
	}
	return true
}

func DeleteSeller(id string) bool {
	result := Database.Delete(Seller{ID: id})
	if result.Error != nil {
		fmt.Println("Error in deleteSeller", result.Error)
		return false
	}
	return true
}

func CreateService(service Service) bool {
	result := Database.Create(&service)
	if result.Error != nil {
		fmt.Println("Error in createService")
		return false
	}
	return true
}

func FindManyService(service Service) bool {
	result := Database.Find(&service)
	if result.Error != nil {
		fmt.Println("Error in findManyService", result.Error)
		return false
	}
	return true
}

func FindOneService(service Service) bool {
	result := Database.Take(&service)
	if result.Error != nil {
		fmt.Println("Error in findOneService", result.Error)
		return false
	}
	return true
}

func UpdateService(service Service) bool {
	result := Database.Model(&service).Update("name", "hello")
	// result := Database.Model(&service).Update(map[string]interface{}{"Name": "name", "Price": "price", "Location": "location", "Photos": "photos"})		 // TODO: Откуда взять инфу на апдейт? Как вообще апдейтить?
	if result.Error != nil {
		fmt.Println("Error in updateService", result.Error)
		return false
	}
	return true
}

func DeleteService(id string) bool {
	result := Database.Delete(Service{ID: id})
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

func FindManySellerToService(sellerId string) ([]SellerToService, bool) {
	var sellerToService []SellerToService
	result := Database.Find(&sellerToService, "seller_id = ?", sellerId)
	if result.Error != nil {
		fmt.Println("Error in findManySellerToService", result.Error)
		return sellerToService, false
	}
	return sellerToService, true
}

func FindOneSellerToService(sellerToService SellerToService) (SellerToService, bool) {
	result := Database.Take(&sellerToService)
	if result.Error != nil {
		fmt.Println("Error in findOneSellerToService", result.Error)
		return sellerToService, false
	}
	return sellerToService, true
}

func DeleteSellerToService(serviceId string) bool {	// Связь удаляется по услуге, т.к. оная удаляется чаще
	result := Database.Delete(SellerToService{ServiceID: serviceId})
	if result.Error != nil {
		fmt.Println("Error in deleteSellerToService", result.Error)
		return false
	}
	return true
}
