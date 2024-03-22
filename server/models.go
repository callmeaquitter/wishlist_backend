package server

//LIFEHACK: use inline todos

type AuthCredentials struct {
	Login    string `json:"login" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}

type SellerAuthCredentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Session string `json:"session"`
}
