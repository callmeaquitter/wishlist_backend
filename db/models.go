package db

import "time"

type Gift struct {
	//LIFEHACK: use string id like 'gift_ajdsjanjklsnls'
	ID          string `json:"id"`
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
	ID      string    `json:"id"`
	GiftID  string    `json:"gift_id"`
	Mark    float32   `json:"mark"`
	Comment string    `json:"comment"`
	Date    time.Time `json:"date"`
}

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Birthday string `json:"birthday"`
	Coins    int    `json:"coins"`
	RoleName string `json:"role_name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Session struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
}

type UserWishlist struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	UserID string `json:"user_id"`
}

type Wishes struct {
	GiftID     string `json:"gift_id" gorm:"primaryKey"`
	WishlistID string `json:"wishlist_id" gorm:"primaryKey"`
}

type Role struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserRole struct {
	UserID string `gorm:"primaryKey"`
	RoleID string `gorm:"primaryKey"`
}
