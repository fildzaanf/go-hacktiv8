package service

import (
	"errors"
	"final-project/middlewares"
	"final-project/modules/user/entity"
	"final-project/utils/helper/bcrypt"
	"final-project/utils/helper/email/mailer/notification"
	"final-project/utils/helper/email/mailer/onetimepassword"
	"final-project/utils/helper/generator"
	"final-project/utils/validator"
	"time"
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

	errEmpty := validator.IsDataEmpty([]string{"email", "password", "confirm_password"}, userCore.Email, userCore.Password, userCore.ConfirmPassword)
	if errEmpty != nil {
		return entity.User{}, errEmpty
	}

	errEmailValid := validator.IsEmailValid(userCore.Email)
	if errEmailValid != nil {
		return entity.User{}, errEmailValid
	}

	errLength := validator.IsMinLengthValid(10, map[string]string{"password": userCore.Password})
	if errLength != nil {
		return entity.User{}, errLength
	}

	_, errGetEmail := us.userRepository.GetUserByEmail(userCore.Email)
	if errGetEmail == nil {
		return entity.User{}, errors.New("email already exists")
	}

	if userCore.Password != userCore.ConfirmPassword {
		return entity.User{}, errors.New("password and confirm password does not match")
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

	notification.SendEmailNotificationRegisterAccount(userData.Email)

	return userData, nil
}

func (us *userService) Login(email, password string) (entity.User, string, error) {

	errEmpty := validator.IsDataEmpty([]string{"email", "password"}, email, password)
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

	token, errGenerate := middlewares.GenerateToken(userData.ID, userData.Role)
	if errGenerate != nil {
		return entity.User{}, "", errors.New("generate token failed")
	}

	notification.SendEmailNotificationLoginAccount(email)

	return userData, token, nil
}

func (us *userService) GetAllUsers() ([]entity.User, error) {
	users, errGetAllUsersRepo := us.userRepository.GetAllUsers()
	if errGetAllUsersRepo != nil {
		return nil, errGetAllUsersRepo
	}

	return users, nil
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

func (us *userService) GetUserByEmail(email string) (entity.User, error) {
	if email == "" {
		return entity.User{}, errors.New("invalid email")
	}

	userData, errGetEmail := us.userRepository.GetUserByEmail(email)
	if errGetEmail != nil {
		return entity.User{}, errors.New("data is empty")
	}
	return userData, nil
}

func (us *userService) GetUserByUsername(username string) (entity.User, error) {
	if username == "" {
		return entity.User{}, errors.New("invalid username")
	}

	userData, errGetUname := us.userRepository.GetUserByUsername(username)
	if errGetUname != nil {
		return entity.User{}, errors.New("data is empty")
	}
	return userData, nil
}

func (us *userService) UpdateUserByID(userID string, userCore entity.User) (entity.User, error) {
	if userID == "" {
		return entity.User{}, errors.New("invalid id")
	}

	userData, errGetID := us.userRepository.GetUserByID(userID)
	if errGetID != nil {
		return entity.User{}, errors.New("user not found")
	}

	if userCore.Email != "" {

		errEmailValid := validator.IsEmailValid(userCore.Email)
		if errEmailValid != nil {
			return entity.User{}, errEmailValid
		}

		_, errGetEmail := us.userRepository.GetUserByEmail(userCore.Email)
		if errGetEmail == nil {
			return entity.User{}, errors.New("email already exists")
		}
	}

	if userCore.Password != "" {

		errEmpty := validator.IsDataEmpty([]string{"password", "new_password", "confirm_password"}, userCore.Password, userCore.NewPassword, userCore.ConfirmPassword)
		if errEmpty != nil {
			return entity.User{}, errEmpty
		}

		comparePassword := bcrypt.ComparePassword(userData.Password, userCore.Password)
		if comparePassword != nil {
			return entity.User{}, errors.New("incorrect password")
		}

		// if userCore.Password != userModel.Password {
		// 	return entity.User{}, errors.New("password incorrect")
		// }

		if userCore.NewPassword != userCore.ConfirmPassword {
			return entity.User{}, errors.New("new password and confirm password does not match")
		}

		errLength := validator.IsMinLengthValid(10, map[string]string{"password": userCore.Password, "new_password": userCore.NewPassword, "confirm_password": userCore.ConfirmPassword})
		if errLength != nil {
			return entity.User{}, errLength
		}

		hashedNewPassword, errHash := bcrypt.HashPassword(userCore.NewPassword)
		if errHash != nil {
			return entity.User{}, errors.New("error hashing password")
		}
		userCore.Password = hashedNewPassword

	}

	if userCore.Username != "" {
		_, errGetUname := us.userRepository.GetUserByUsername(userCore.Username)
		if errGetUname == nil {
			return entity.User{}, errors.New("username already exists")
		}
	}

	userData, errUpdate := us.userRepository.UpdateUserByID(userID, userCore)
	if errUpdate != nil {
		return entity.User{}, errUpdate
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

func (us *userService) SendOTP(email string) error {

	errEmpty := validator.IsDataEmpty([]string{"email"}, email)
	if errEmpty != nil {
		return errEmpty
	}

	errEmailValid := validator.IsEmailValid(email)
	if errEmailValid != nil {
		return errEmailValid
	}

	otp, errGenerateOTP := generator.GenerateRandomCode()
	if errGenerateOTP != nil {
		return errors.New("failed to generate otp")
	}

	expired := time.Now().Add(7 * time.Minute).Unix()

	_, errSend := us.userRepository.SendOTP(email, otp, expired)
	if errSend != nil {
		return errSend
	}

	onetimepassword.SendEmailOTP(email, otp)
	return nil
}

func (us *userService) VerifyOTP(email, otp string) (string, error) {

	errEmpty := validator.IsDataEmpty([]string{"email", "otp"}, email, otp)
	if errEmpty != nil {
		return "", errEmpty
	}

	userData, err := us.userRepository.VerifyOTP(email, otp)
	if err != nil {
		return "", errors.New("email or otp not found")
	}

	if userData.OTPExpiration <= time.Now().Unix() {
		return "", errors.New("otp has expired")
	}

	if userData.OTP != otp {
		return "", errors.New("invalid otp")
	}

	token, err := middlewares.GenerateVerifyToken(email)
	if err != nil {
		return "", errors.New("generate token failed")
	}

	_, errReset := us.userRepository.ResetOTP(otp)
	if errReset != nil {
		return "", errors.New("failed to reset otp")
	}

	return token, nil
}

func (us *userService) NewPassword(email string, userCore entity.User) error {

	errEmpty := validator.IsDataEmpty([]string{"email", "password", "confirm_password"}, email, userCore.Password, userCore.ConfirmPassword)
	if errEmpty != nil {
		return errEmpty
	}

	errEmailValid := validator.IsEmailValid(email)
	if errEmailValid != nil {
		return errEmailValid
	}

	errLength := validator.IsMinLengthValid(10, map[string]string{"password": userCore.Password})
	if errLength != nil {
		return  errLength
	}

	if userCore.Password != userCore.ConfirmPassword {
		return errors.New("password and confirm password does not match")
	}

	HashPassword, errHash := bcrypt.HashPassword(userCore.Password)
	if errHash != nil {
		return errors.New("error hashing password")
	}
	userCore.Password = HashPassword

	_, errNewPass := us.userRepository.NewPassword(email, userCore)
	if errNewPass != nil {
		return errNewPass
	}

	return nil
}
