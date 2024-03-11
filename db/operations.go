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

func FindManySeller() ([]Seller, bool) {
	var seller []Seller
	result := Database.Find(&seller)
	if result.Error != nil {
		fmt.Println("Error in findManySeller", result.Error)
		return seller, false
	}
	return seller, true
}

func FindOneSeller(sellerId string) (Seller, bool) {
	var seller Seller
	result := Database.Take(&seller, "id = ?", sellerId)
	if result.Error != nil {
		fmt.Println("Error in findOneSeller", result.Error)
		return seller, false
	}
	return seller, true
}

func UpdateSeller(seller Seller) bool {
	validation := Database.Find(&seller, "id = ?", seller.ID)
	if validation.Error != nil {
		fmt.Println("Error in updateSeller", validation.Error)
		return false
	}

	if seller.Name != "string" {
		result := Database.Model(&seller).Update("name", seller.Name)
		if result.Error != nil {
			fmt.Println("Error in updateSeller", result.Error)
			return false
		}
	}
	if seller.Email != "string" {
		result := Database.Model(&seller).Update("email", seller.Email)
		if result.Error != nil {
			fmt.Println("Error in updateSeller", result.Error)
			return false
		}
	}
	if seller.Photo != "string" {
		result := Database.Model(&seller).Update("photo", seller.Photo)
		if result.Error != nil {
			fmt.Println("Error in updateSeller", result.Error)
			return false
		}
	}
	
	return true
}

func DeleteSeller(id string) bool {
	result := Database.Delete(&Seller{}, "id = ?", id)
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
	validation := Database.Find(&service, "id = ?", service.ID)
	if validation.Error != nil {
		fmt.Println("Error in updateservice", validation.Error)
		return false
	}

	if service.Name != "string" {
		result := Database.Model(&service).Update("name", service.Name)
		if result.Error != nil {
			fmt.Println("Error in updateservice", result.Error)
			return false
		}
	}
	if service.Price != 0 {
		result := Database.Model(&service).Update("price", service.Price)
		if result.Error != nil {
			fmt.Println("Error in updateservice", result.Error)
			return false
		}
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

func DeleteSellerToService(serviceId string) bool {	// Связь удаляется по услуге, т.к. оная удаляется чаще
	result := Database.Delete(SellerToService{}, "service_id = ?", serviceId)
	if result.Error != nil {
		fmt.Println("Error in deleteSellerToService", result.Error)
		return false
	}
	return true
}
