package db

import "gorm.io/gorm"

type Gift struct {
	//LIFEHACK: use string id like 'gift_ajdsjanjklsnls'
	ID string 		   `json:"id"`
	// UserID      string `json:"user_id"`
	Name        string `json:"name"`
	Price       int    `json:"price"` //TODO: use decimal.Decimal instead of int
	Photo       string `json:"photo"`
	Description string `json:"description"`
	IsFavorite  bool   `json:"is_favorite"`
	Link        string `json:"link"`
	Comments    string `json:"comments"`
}

type User struct {
	ID           int    		`json:"id"`
	Name         string 		`json:"name"`
	Price        int    		`json:"price"`
	Description  string 		`json:"description"`
	IsFavorite   string 		`json:"is_favorite"`
	Link         string 		`json:"link"`
	Birthday     string 		`json:"birthday"`
	Coins        int    		`json:"coins"`
	gorm.Model
	Role 		 Role   		`gorm:"foreigKey:role_id"`
	UserWishlist []UserWishlist `gorm:"foreignKey:id"`
}

type UserWishlist struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	UserID int    `json:"user_id"`
	GiftID int    `json:"gift_id"`

}

type Role struct {
	ID     int 	  `json:"id"`
	RoleID int    `json:"role_id"`
	Name   string `json:"name"`
	gorm.Model
	
}