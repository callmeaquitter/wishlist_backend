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
	// result := Database.Model(&service).Updates(map[string]interface{}{})	// TODO: Откуда взять данные на обнову?
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

func CreateSellersServices(sellersServices Sellers_services) bool {
	result := Database.Create(&sellersServices)
	if result.Error != nil {
		fmt.Println("Error in CreateSellersServices")
		return false
	}
	return true
}

func FindManySellersServices(sellerId string) ([]Sellers_services, bool) {
	var sellersServices []Sellers_services
	result := Database.Find(&sellersServices, "seller_id = ?", sellerId)
	if result.Error != nil {
		fmt.Println("Error in findManySellersServices", result.Error)
		return sellersServices, false
	}
	return sellersServices, true
}

func FindOneSellersServices(sellersServices Sellers_services) (Sellers_services, bool) {
	result := Database.Take(&sellersServices)
	if result.Error != nil {
		fmt.Println("Error in findOneSellersServices", result.Error)
		return sellersServices, false
	}
	return sellersServices, true
}

func DeleteSellersServices(serviceId string) bool {	// Связь удаляется по услуге, т.к. оная удаляется чаще
	result := Database.Delete(Sellers_services{ServiceID: serviceId})
	if result.Error != nil {
		fmt.Println("Error in deleteSellersServices", result.Error)
		return false
	}
	return true
}
