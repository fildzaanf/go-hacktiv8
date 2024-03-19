package service

import (
	"assignment-2/middlewares"
	"assignment-2/modules/user/entity"
	"assignment-2/utils/helper/bcrypt"
	"assignment-2/utils/validator"
	"errors"
)

type userService struct {
	userRepository entity.UserRepositoryInterface
}

func NewUserService(userRepository entity.UserRepositoryInterface) entity.UserServiceInterface {
	return &userService{
		userRepository: userRepository,
	}
}

func (us *userService) Register(userCore entity.User) (entity.User, error) {

	errEmpty := validator.IsDataEmpty(userCore.Fullname, userCore.Email, userCore.Password)
	if errEmpty != nil {
		return entity.User{}, errEmpty
	}

	errEmailValid := validator.IsEmailValid(userCore.Email)
	if errEmailValid != nil {
		return entity.User{}, errEmailValid
	}

	errLength := validator.IsMinLengthValid(userCore.Password, 10)
	if errLength != nil {
		return entity.User{}, errLength
	}

	_, errGetEmail := us.userRepository.GetUserByEmail(userCore.Email)
	if errGetEmail == nil {
		return entity.User{}, errors.New("email already exists")
	}

	hashedPassword, errHash := bcrypt.HashPassword(userCore.Password)
	if errHash != nil {
		return entity.User{}, errors.New("error hashing password")
	}

	userCore.Password = hashedPassword

	userData, errRegister := us.userRepository.Register(userCore)
	if errRegister != nil {
		return entity.User{}, errRegister
	}

	return userData, nil
}

func (us *userService) Login(email, password string) (entity.User, string, error) {

	errEmpty := validator.IsDataEmpty(email, password)
	if errEmpty != nil {
		return entity.User{}, "", errEmpty
	}

	errEmailValid := validator.IsEmailValid(email)
	if errEmailValid != nil {
		return entity.User{}, "", errEmailValid
	}

	userData, errGetEmail := us.userRepository.GetUserByEmail(email)
	if errGetEmail != nil {
		return entity.User{}, "", errors.New("email not registered")
	}

	comparePassword := bcrypt.ComparePassword(userData.Password, password)
	if comparePassword != nil {
		return entity.User{}, "", errors.New("incorrect email or password")
	}

	token, errCreateToken := middlewares.GenerateToken(userData.ID, "")
	if errCreateToken != nil {
		return entity.User{}, "", errors.New("generate token failed")
	}
	return userData, token, nil
}

func (us *userService) GetUserByID(userID string) (entity.User, error) {
	if userID == "" {
		return entity.User{}, errors.New("invalid id")
	}

	userData, errGetID := us.userRepository.GetUserByID(userID)
	if errGetID != nil {
		return entity.User{}, errors.New("data is empty")
	}
	return userData, nil
}

func (us *userService) UpdateUserByID(userID string, userCore entity.User) (entity.User, error) {
	if userID == "" {
		return entity.User{}, errors.New("invalid id")
	}

	_, errGetID := us.userRepository.GetUserByID(userID)
	if errGetID != nil {
		return entity.User{}, errGetID
	}

	errEmailValid := validator.IsEmailValid(userCore.Email)
	if errEmailValid != nil {
		return  entity.User{}, errEmailValid
	}

	userData, errUpdate := us.userRepository.UpdateUserByID(userID, userCore)
	if errUpdate != nil {
		return  entity.User{}, errUpdate
	}

	return userData, nil
}

func (us *userService) DeleteUserByID(userID string) error {
	if userID == "" {
		return errors.New("id not found")
	}

	errDeleteUserByID := us.userRepository.DeleteUserByID(userID)
	if errDeleteUserByID != nil {
		return errDeleteUserByID
	}

	return nil
}
