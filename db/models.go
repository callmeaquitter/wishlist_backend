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

type Quest struct {
	ID         int 		`json:"id"`
	SubquestID int    	`json:"subquest_id"`
	UserID     string   `json:"user_id"`
	IsDone     bool   	`json:"is_done"`
}

type Subquest struct {
	ID			int		`json:"id"`
	TaskID		int		`json:"task_id"`
	Reward		int		`json:"reward"`
	IsDone		int		`json:"is_done"`
}

type Tasks struct {
	ID				int			`json:"id"`
	Name			string		`json:"name"`
	Description		string 		`json:"description"`
}