package db

import "time"

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
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      string `json:"user_id"`
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
