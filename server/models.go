package server

//LIFEHACK: use inline todos

type AuthCredentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Session string `json:"session"`
}
