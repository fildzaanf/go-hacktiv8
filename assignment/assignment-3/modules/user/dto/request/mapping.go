package request

import "assignment-2/modules/user/entity"

func UserRegisterRequestToUserCore(request UserRegisterRequest) entity.User {
	return entity.User{
		Fullname: request.Fullname,
		Email:    request.Email,
		Password: request.Password,
	}
}

func UserLoginRequestToUserCore(request UserLoginRequest) entity.User {
	return entity.User{
		Email:    request.Email,
		Password: request.Password,
	}
}

func UserRequestToUserCore(request UserRequest) entity.User {
	return entity.User{
		Fullname: request.Fullname,
		Email:    request.Email,
		Password: request.Password,
	}
}
