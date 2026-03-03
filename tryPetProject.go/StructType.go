package main

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"passwordhash"`
	Message  string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type LoginResponse struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Message  string `json:"message"`
}
