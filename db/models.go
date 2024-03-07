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

type Selection struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID       int   `json:"user_id"`
}

type GiftToSelection struct {
	SelectionID  int   `json:"selection_id" gorm:"primaryKey"`
	GiftID       int   `json:"gift_id" gorm:"primaryKey"`
}

type SelectionCategory struct {
	ID   int	       `json:"id"`
	Name string        `json:"name"`
}
