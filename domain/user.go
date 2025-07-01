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

type DelRequest struct {
	UserID uint
}

type DelResponse struct {
}

type ListRequest struct {
	Page     int
	PageSize int
}

type ListResponse struct {
	UsersInfo []UserInfo
}
