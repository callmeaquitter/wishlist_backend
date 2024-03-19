package db

import (
	"time"

	"github.com/shopspring/decimal"
)

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

type Selection struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      string    `json:"user_id"`
	IsGenerated bool   `json:"is_generated"`
}

type GiftToSelection struct {
	SelectionID string `json:"selection_id" gorm:"primaryKey"`
	GiftID      string `json:"gift_id" gorm:"primaryKey"`
}

type SelectionCategory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type LikeToSelection struct {
	UserID      string `json:"user_id" gorm:"primaryKey"`
	SelectionID string `json:"selection_id" gorm:"primaryKey"`
}

type CommentToSelection struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	SelectionID string    `json:"selection_id"`
	Text        string    `json:"text"`
	CreatedAt   time.Time `json:"created_at"`
}

type Seller struct {
	SellerID string `json:"id" gorm:"primaryKey" swaggerignore:"true"`
	Name     string `json:"name" validate:"required,min=5,max=50"`
	Email    string `json:"email" validate:"required,min=5,email"`
	Photo    string `json:"photo"`
	
}

type SellerToService struct {
	SellerID  string `json:"seller_id" gorm:"primaryKey" validate:"required,seller_"`
	ServiceID string `json:"service_id" gorm:"primaryKey" validate:"required,service_"`
}

type Service struct {
	ServiceID string          `json:"id" gorm:"primaryKey" swaggerignore:"true"`
	Name      string          `json:"name" validate:"required,min=5,max=50"`
	Price     decimal.Decimal `json:"price" validate:"required"`
	Location  string          `json:"location" validate:"required,min=5"`
	Photos    string          `json:"photos"`
}

type ServiceReview struct {
	ID         string          `json:"id" swaggerignore:"true"`
	ServiceID  string          `json:"service_id" validate:"required,service_"`
	Mark       decimal.Decimal `json:"mark" validate:"required"`
	Comment    string          `json:"comment" validate:"required,max=5000"`
	UserID     string          `json:"user_id" validate:"required,user_"`
	CreateDate time.Time       `json:"create_date" swaggerignore:"true"`
	UpdateDate time.Time       `json:"update_date" swaggerignore:"true"`
}


type Quest struct {
	ID         string `json:"id"`
	SubquestID string `json:"subquest_id"`
	UserID     string `json:"user_id"`
	IsDone     bool   `json:"is_done"`
}

type Subquest struct {
	ID     string `json:"id"`
	TaskID string `json:"task_id"`
	Reward int    `json:"reward"`
	IsDone int    `json:"is_done"`
}

type Tasks struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type OfflineShops struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}
