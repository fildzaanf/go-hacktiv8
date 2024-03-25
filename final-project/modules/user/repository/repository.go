package repository

import (
	"final-project/modules/user/entity"
	"final-project/modules/user/model"
	"final-project/utils/helper/bcrypt"
	"errors"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) entity.UserRepositoryInterface {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Register(userCore entity.User) (entity.User, error) {
	request := entity.UserCoreToUserModel(userCore)

	tx := ur.db.Create(&request)
	if tx.Error != nil {
		return entity.User{}, tx.Error
	}

	response := entity.UserModelToUserCore(request)
	return response, nil
}

func (ur *userRepository) Login(email, password string) (entity.User, error) {
	userModel := model.User{}

	tx := ur.db.Where("email = ?", email).Preload("Photos").Preload("Comments").Preload("MediaSocials").First(&userModel) 
	if tx.Error != nil {
		return entity.User{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.User{}, errors.New("email not found")
	}

	if errComparePass := bcrypt.ComparePassword(userModel.Password, password); errComparePass != nil {
		return entity.User{}, errors.New("invalid password")
	}

	response := entity.UserModelToUserCore(userModel)
	return response, nil
}

func (ur *userRepository) GetUserByID(userID string) (entity.User, error) {

	userModel := model.User{}

	tx := ur.db.Where("id = ?", userID).Preload("Photos").Preload("Comments").Preload("MediaSocials").First(&userModel) 
	if tx.Error != nil {
		return entity.User{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.User{}, errors.New("data not found")
	}

	response := entity.UserModelToUserCore(userModel)
	return response, nil
}

func (ur *userRepository) UpdateUserByID(userID string, userCore entity.User) (entity.User, error) {
	userModel := model.User{}

	tx := ur.db.Where("id = ?", userID).Preload("Photos").Preload("Comments").Preload("MediaSocials").First(&userModel) 
	if tx.Error != nil {
		return entity.User{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.User{}, errors.New("id not found")
	}

	updatedUserModel := entity.UserCoreToUserModel(userCore)

	tx = ur.db.Model(&userModel).Updates(updatedUserModel)
	if tx.Error != nil {
		return entity.User{}, tx.Error
	}

	response := entity.UserModelToUserCore(userModel)
	return response, nil
}

func (ur *userRepository) GetUserByEmail(email string) (entity.User, error) {
	userModel := model.User{}

	tx := ur.db.Where("email = ?", email).Preload("Photos").Preload("Comments").Preload("MediaSocials").First(&userModel) 

	if tx.RowsAffected == 0 {
		return entity.User{}, errors.New("email not found")
	}

	if tx.Error != nil {
		return entity.User{}, tx.Error
	}

	response := entity.UserModelToUserCore(userModel)
	return response, nil
}

func (ur *userRepository) GetAllUsers() ([]entity.User, error) {
	userModel := []model.User{}

	tx := ur.db.Model(&userModel).Preload("Photos").Preload("Comments").Preload("MediaSocials").Find(&userModel)
	if tx.Error != nil {
		return nil, tx.Error
	}

	response := entity.ListUserModelToUserCore(userModel)
	return response, nil
}

func (ur *userRepository) GetUserByUsername(username string) (entity.User, error) {
	userModel := model.User{}

	tx := ur.db.Where("username = ?", username).Preload("Photos").Preload("Comments").Preload("MediaSocials").First(&userModel) 

	if tx.RowsAffected == 0 {
		return entity.User{}, errors.New("username not found")
	}

	if tx.Error != nil {
		return entity.User{}, tx.Error
	}

	response := entity.UserModelToUserCore(userModel)
	return response, nil
}

func (ur *userRepository) DeleteUserByID(userID string) error {
	userModel := model.User{}

	tx := ur.db.Where("id = ?", userID).Delete(&userModel)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("data not found")
	}
	return nil
}


func (ur *userRepository) SendOTP(email string, otp string, expired int64) (userCore entity.User, err error) {
	userModel := model.User{}

	tx := ur.db.Where("email = ?", email).First(&userModel)
	if tx.Error != nil {
		if tx.RowsAffected == 0 {
			return entity.User{}, errors.New("email not found")
		}
		return entity.User{}, tx.Error
	}

	userModel.OTP = otp
	userModel.OTPExpiration = expired

	errUpdate := ur.db.Save(&userModel).Error
	if errUpdate != nil {
		return entity.User{}, errUpdate
	}

	userCore = entity.UserModelToUserCore(userModel)

	return userCore, nil
}

func (ur *userRepository) VerifyOTP(email, otp string) (entity.User, error) {
	userModel := model.User{}

	tx := ur.db.Where("otp = ? AND email = ?", otp, email).First(&userModel)
	if tx.Error != nil {
		return entity.User{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.User{}, errors.New("email or otp not found")
	}

	userCore := entity.UserModelToUserCore(userModel)
	return userCore, nil
}

func (ur *userRepository) ResetOTP(otp string) (userCore entity.User, err error) {
	userModel := model.User{}

	tx := ur.db.Where("otp = ?", otp).First(&userModel)
	if tx.Error != nil {
		return entity.User{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.User{}, errors.New("otp not found")
	}

	userModel.OTP = ""
	userModel.OTPExpiration = 0

	errUpdate := ur.db.Save(&userModel).Error
	if errUpdate != nil {
		return entity.User{}, errUpdate
	}

	userCore = entity.UserModelToUserCore(userModel)
	return userCore, nil
}


func (ur *userRepository) NewPassword(email string, userCore entity.User) (entity.User, error) {
	userModel := model.User{}

	tx := ur.db.Where("email = ?", email).First(&userModel)
	if tx.Error != nil {
		return entity.User{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.User{}, errors.New("email not found")
	}

	userModelUpdate := entity.UserCoreToUserModel(userCore)

	errUpdate := ur.db.Model(&userModel).Updates(userModelUpdate)
	if errUpdate != nil {
		return entity.User{}, errUpdate.Error
	}

	userCore = entity.UserModelToUserCore(userModel)

	return userCore, nil
}


