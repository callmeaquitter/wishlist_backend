package db

import (
	"time"

	"github.com/shopspring/decimal"
)

// import "github.com/go-delve/delve/service"


type Gift struct {
	//LIFEHACK: use string id like 'gift_ajdsjanjklsnls'
	ID		string	`json:"id"`
	// UserID      string `json:"user_id"`
	Name		string	`json:"name"`
	Price		int   	`json:"price"` //TODO: use decimal.Decimal instead of int
	Photo		string	`json:"photo"`
	Description	string	`json:"description"`
	IsFavorite	bool  	`json:"is_favorite"`
	Link		string	`json:"link"`
	Comments	string	`json:"comments"`
}

type Seller struct {
	SellerID	string		`json:"id" gorm:"primaryKey" swaggerignore:"true"`
	Name		string		`json:"name" validate:"required,min=5,max=50"`
	Email		string		`json:"email" validate:"required,min=5,email"`
	Photo		string		`json:"photo"`
}

type SellerToService struct {
	SellerID	string		`json:"seller_id" gorm:"primaryKey" validate:"required,seller_"`
	ServiceID	string		`json:"service_id" gorm:"primaryKey" validate:"required,service_"`
}

type Service struct {
	ServiceID	string		`json:"id" gorm:"primaryKey" swaggerignore:"true"`
	Name		string		`json:"name" validate:"required,min=5,max=50"`
	Price		decimal.Decimal	`json:"price" validate:"required"`
	Location	string		`json:"location" validate:"required,min=5"`
	Photos		string		`json:"photos"`
}

type ServiceReview struct {
	ID		string		`json:"id" swaggerignore:"true"`
	ServiceID	string		`json:"service_id" validate:"required,service_"`
	Mark		decimal.Decimal	`json:"mark" validate:"required"`
	Comment		string		`json:"comment" validate:"required,max=5000"`
	UserID		string		`json:"user_id" validate:"required,user_"`
	CreateDate	time.Time	`json:"create_date" swaggerignore:"true"`
	UpdateDate	time.Time	`json:"update_date" swaggerignore:"true"`
}

