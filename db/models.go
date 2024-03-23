package db

import (
	"time"

	"github.com/shopspring/decimal"
)

type Gift struct {
	//LIFEHACK: use string id like 'gift_ajdsjanjklsnls'
	ID          string `json:"id"`
	Name        string `json:"name" validate:"required"`
	Price       int    `json:"price" validate:"required"` //TODO: use decimal.Decimal instead of int
	Photo       string `json:"photo"`
	Description string `json:"description"`
	Link        string `json:"link" validate:"required"`
	Category    string `json:"category"`
}

type BookedGiftInWishlist struct {
	UserID string `json:"user_id" gorm:"primaryKey" validate:"required,user_"`
	GiftID string `json:"gift_id" gorm:"primaryKey" validate:"required,gift_"`
}

type GiftCategory struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type GiftReview struct {
	ID      string    `json:"id"`
	GiftID  string    `json:"gift_id" validate:"required,gift_"`
	Mark    float32   `json:"mark" validate:"required"`
	Comment string    `json:"comment" validate:"required"`
	Date    time.Time `json:"date"`
}

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name" validate:"required"`
	Lastname string `json:"lastname" validate:"required"`
	Birthday string `json:"birthday" validate:"required"`
	Coins    int    `json:"coins"`
	RoleName string `json:"role_name"`
	Login    string `json:"login" validate:"required,email" gorm:"unique"`
	Password string `json:"password" validate:"required,password"`
}

type Session struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
}

type UserWishlist struct {
	ID     string `json:"id"`
	Name   string `json:"name" validate:"required"`
	UserID string `json:"user_id" validate:"required,user_"`
}

type Wishes struct {
	GiftID     string `json:"gift_id" gorm:"primaryKey" validate:"required,gift_"`
	WishlistID string `json:"wishlist_id" gorm:"primaryKey" validate:"required,wishlist_"`
}

type Role struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type UserRole struct {
	UserID string `gorm:"primaryKey" validate:"required,user_"`
	RoleID string `gorm:"primaryKey" validate:"required,role_"`
}

type Selection struct {
	ID          string `json:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	UserID      string    `json:"user_id" validate:"required,user_"`
	IsGenerated bool   `json:"is_generated"`
}

type GiftToSelection struct {
	SelectionID string `json:"selection_id" gorm:"primaryKey" validate:"required,selection_"`
	GiftID      string `json:"gift_id" gorm:"primaryKey" validate:"required,gift_"`
}

type SelectionCategory struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type LikeToSelection struct {
	UserID      string `json:"user_id" gorm:"primaryKey" validate:"required"`
	SelectionID string `json:"selection_id" gorm:"primaryKey" validate:"required"`
}

type CommentToSelection struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id" validate:"required"`
	SelectionID string    `json:"selection_id" validate:"required"`
	Text        string    `json:"text" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
}

type Seller struct {
	SellerID string `json:"id" gorm:"primaryKey" swaggerignore:"true"`
	Name     string `json:"name" validate:"required,max=50"`
	Photo    string `json:"photo"`
	RoleName string `json:"role_name"`
	Login    string `json:"login" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}

type SellerSession struct {
	ID		string	`json:"id"`
	SellerID	string	`json:"seller_id"`
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
	SubquestID string `json:"subquest_id" validate:"required,subquest_"`
	UserID     string `json:"user_id" validate:"required,user_"`
	IsDone     bool   `json:"is_done" validate:"required"`
}

type Subquest struct {
	ID     string `json:"id"`
	TaskID string `json:"task_id" validate:"required,task_"`
	Reward int    `json:"reward" validate:"required"`
	IsDone int    `json:"is_done"`
}

type Tasks struct {
	ID          string `json:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type OfflineShops struct {
	ID       string `json:"id"`
	Name     string `json:"name" validate:"required"`
	Location string `json:"location" validate:"required"`
}
