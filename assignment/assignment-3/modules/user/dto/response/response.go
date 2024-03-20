package response

import "assignment-2/modules/order/dto/response"

type UserRegisterResponse struct {
	ID       string `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type UserLoginResponse struct {
	ID       string `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

type UserResponse struct {
	ID       string                   `json:"id"`
	Fullname string                   `json:"fullname"`
	Email    string                   `json:"email"`
	Role     string                   `json:"role"`
	Orders   []response.OrderResponse `json:"orders"`
}
