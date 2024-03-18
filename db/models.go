package db

import "time"

type Gift struct {
	//LIFEHACK: use string id like 'gift_ajdsjanjklsnls'
	ID string `json:"id"`
	// UserID      string `json:"user_id"`
	Name        string `json:"name"`
	Price       int    `json:"price"` //TODO: use decimal.Decimal instead of int
	Photo       string `json:"photo"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Category    string `json:"category"`
}

type BookedGiftInWishlist struct {
	UserID string `json:"user_id" gorm:"primaryKey"`
	GiftID string `json:"gift_id" gorm:"primaryKey"`
}

type GiftCategory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GiftReview struct {
	ID      string  	`json:"id"`
	GiftID  string  	`json:"gift_id"`
	Mark    float32 	`json:"mark"`
	Comment string  	`json:"comment"`
	Date    time.Time	`json:"date"`
}
