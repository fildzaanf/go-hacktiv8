package response

import "assignment-2/modules/user/entity"

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

func UserCoreToUserProfileResponse(response entity.User) UserProfileResponse {
	return UserProfileResponse{
		ID:       response.ID,
		Fullname: response.Fullname,
		Email:    response.Email,
	}
}
