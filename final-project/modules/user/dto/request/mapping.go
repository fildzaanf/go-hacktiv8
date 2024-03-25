package request

import "final-project/modules/user/entity"

func UserRegisterRequestToUserCore(request UserRegisterRequest) entity.User {
	return entity.User{
		Email:           request.Email,
		Password:        request.Password,
		ConfirmPassword: request.ConfirmPassword,
	}
}

func UserLoginRequestToUserCore(request UserLoginRequest) entity.User {
	return entity.User{
		Email:    request.Email,
		Password: request.Password,
	}
}

func UserUpdateProfileRequestToUserCore(request UserUpdateProfileRequest) entity.User {
	return entity.User{
		Fullname:        request.Fullname,
		Username:        request.Username,
		Email:           request.Email,
		Age:             request.Age,
		Password:        request.Password,
		NewPassword:     request.NewPassword,
		ConfirmPassword: request.ConfirmPassword,
	}
}

func UserNewPasswordRequestToUserCore(request UserNewPasswordRequest) entity.User {
	return entity.User{
		Password:        request.Password,
		ConfirmPassword: request.ConfirmPassword,
	}
}

func UserSendOTPRequestToUserCore(request UserSendOTPRequest) entity.User {
	return entity.User{
		Email: request.Email,
	}
}

func UserVerifyOTPRequestToUserCore(request UserVerifyOTPRequest) entity.User {
	return entity.User{
		Email: request.Email,
		OTP:   request.OTP,
	}
}
