package domain

type UserInfo struct {
	UserID      uint   `json:"user_id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type RegisterRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type RegisterResponse struct {
	UserInfo
}

type ProfileRequest struct {
	UserID uint
}

type ProfileResponse struct {
	UserInfo
}

type DelAccountRequest struct {
	UserID uint
}

type DelAccountResponse struct {
}
