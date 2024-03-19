package repository

import (
	"assignment-2/modules/user/entity"
	"assignment-2/modules/user/model"
	"assignment-2/utils/helper/bcrypt"
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

	tx := ur.db.Where("email = ?", email).First(&userModel)
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

func (ur *userRepository) GetUserByID(id string) (entity.User, error) {

	userModel := model.User{}

	tx := ur.db.Where("id = ?", id).Preload("Orders").First(&userModel)
	if tx.Error != nil {
		return entity.User{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.User{}, errors.New("id not found")
	}

	response := entity.UserModelToUserCore(userModel)
	return response, nil
}

func (ur *userRepository) UpdateUserByID(id string, userCore entity.User) (entity.User, error) {
	request := entity.UserCoreToUserModel(userCore)

	tx := ur.db.Where("id = ?", id).Updates(&request)
	if tx.Error != nil {
		return entity.User{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.User{}, errors.New("id not found")
	}

	response := entity.UserModelToUserCore(request)
	return response, nil
}

func (ur *userRepository) GetUserByEmail(email string) (entity.User, error) {
	userModel := model.User{}

	tx := ur.db.Where("email = ?", email).First(&userModel)

	if tx.RowsAffected == 0 {
		return entity.User{}, errors.New("email not found")
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
