package dto

type RegisterRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type UserInfo struct {
	UserID      uint   `json:"user_id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type RegisterResponse struct {
	UserInfo
}