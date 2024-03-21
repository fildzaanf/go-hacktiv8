package request

import "final-project/modules/user/entity"

func UserRegisterRequestToUserCore(request UserRegisterRequest) entity.User {
	return entity.User{
		Fullname: request.Fullname,
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
		Age:      request.Age,
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
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
		Age:      request.Age,
	}
}
