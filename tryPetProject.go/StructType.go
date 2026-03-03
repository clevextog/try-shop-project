package main

type RegisterResponse struct {
	Email        string `json:"email"`
	PasswordHash string `json:"passwordhash"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type LoginResponse struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
