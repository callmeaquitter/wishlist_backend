package db

type Gift struct {
	//LIFEHACK: use string id like 'gift_ajdsjanjklsnls'
	ID string `json:"id"`
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
	ID          string `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Birthday    string `json:"birthday"`
	Coins       int    `json:"coins"`
	Role_name   Role   `gorm:"foreignKey:Name"`
}

type UserWishlist struct {
	Name   string `json:"name"`
	UserID int    `json:"user_id" gorm:"primaryKey"`
}

type Wishes struct {
	GiftID	string `json:"gift_id" gorm:"primaryKey"`
	WishlistID string `json:"wishlist_id" gorm:"primaryKey"`
}

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserRole struct {
	UserID int `gorm:"primaryKey"`
	RoleID int `gorm:"primaryKey"`
}
