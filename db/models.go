package db

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
	SellerID	string	`json:"id"`
	Name		string	`json:"name"`
	Email		string	`json:"email"`
	Photo		string	`json:"photo"`
}

type Sellers_services struct {
	ID		string	`json:"id" gorm:"primaryKey"`
	SellerID	string	`json:"seller_id" gorm:"foreignKey:Seller_ID;constraint:OnUpdate:RESTRICT,OnDelete:CASCADE;"`
	ServiceID	string	`json:"service_id" gorm:"foreignKey:Service_ID;constraint:OnUpdate:RESTRICT,OnDelete:CASCADE;"`
}

type Service struct {
	ServiceID	string	`json:"id"`
	Name		string	`json:"name"`
	Price		int	`json:"price"` //TODO: use decimal.Decimal instead of int
	Location	string	`json:"location"`
	Photos		string	`json:"photos"`
}

