package response

import (
	or "assignment-2/modules/order/dto/response"
	"assignment-2/modules/user/entity"
)

func UserCoreToUserRegisterResponse(response entity.User) UserRegisterResponse {
	return UserRegisterResponse{
		ID:       response.ID,
		Fullname: response.Fullname,
		Email:    response.Email,
	}
}

func UserCoreToUserLoginResponse(response entity.User, token string) UserLoginResponse {
	return UserLoginResponse{
		ID:       response.ID,
		Fullname: response.Fullname,
		Email:    response.Email,
		Token:    token,
	}
}

func UserCoreToUserResponse(response entity.User) UserResponse {
	return UserResponse{
		ID:       response.ID,
		Fullname: response.Fullname,
		Email:    response.Email,
		Orders:   or.ListOrdeCoreToOrderResponse(response.Orders),
	}
}
