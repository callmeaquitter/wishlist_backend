package db

// import "gorm.io/gorm"
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
	ID		string	`json:"id"`
	Name		string	`json:"name"`
	Email		string	`json:"email"`
	Photo		string	`json:"photo"`
}

type SellerToService struct {
	SellerID	string	`json:"seller_id" gorm:"primaryKey"`
	ServiceID	string	`json:"service_id" gorm:"primaryKey"`
}

type Service struct {
	ID		string	`json:"id"`
	Name		string	`json:"name"`
	Price		int	`json:"price"` //TODO: use decimal.Decimal instead of int
	Location	string	`json:"location"`
	Photos		string	`json:"photos"`
}

