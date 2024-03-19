package response

import "assignment-2/modules/order/dto/response"

type UserRegisterResponse struct {
	ID       string `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

type UserLoginResponse struct {
	ID       string `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type UserResponse struct {
	ID       string                   `json:"id"`
	Fullname string                   `json:"fullname"`
	Email    string                   `json:"email"`
	Orders   []response.OrderResponse `json:"orders"`
}
