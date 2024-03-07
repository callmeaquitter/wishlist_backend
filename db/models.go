package db

type Gift struct {
	//LIFEHACK: use string id like 'gift_ajdsjanjklsnls'
	ID 			string `json:"id"`
	// UserID      string `json:"user_id"`
	Name        string `json:"name"`
	Price       int    `json:"price"` //TODO: use decimal.Decimal instead of int
	Photo       string `json:"photo"`
	Description string `json:"description"`
	Link     	string `json:"link"`
	Comments 	string `json:"comments"`
	Category	string `json:"category"`
}

type BookedGiftInWishlist struct {
	// ID     string `json:"id"`
	UserID string `json:"user_id" gorm:"primaryKey"`
	GiftID string `json:"gift_id" gorm:"primaryKey"`
	
}

type GiftCategory struct {
	ID		string `json:"id"`
	Name 	string `json:"name"`
}
