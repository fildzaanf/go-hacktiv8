package entity

import (
	"assignment-2/modules/order/entity"
	"assignment-2/modules/user/model"
)

func UserCoreToUserModel(userCore User) model.User {
	userModel := model.User{
		ID:        userCore.ID,
		Fullname:  userCore.Fullname,
		Email:     userCore.Email,
		Password:  userCore.Password,
		Role:      userCore.Role,
		Orders:    entity.ListOrderCoreToOrderModel(userCore.Orders),
		CreatedAt: userCore.CreatedAt,
		UpdatedAt: userCore.UpdatedAt,
		DeletedAt: userCore.DeletedAt,
	}
	return userModel
}

func ListUserCoreToUserModel(userCore []User) []model.User {
	listUserModel := []model.User{}
	for _, user := range userCore {
		userModel := UserCoreToUserModel(user)
		listUserModel = append(listUserModel, userModel)
	}
	return listUserModel
}

func UserModelToUserCore(userModel model.User) User {
	userCore := User{
		ID:        userModel.ID,
		Fullname:  userModel.Fullname,
		Email:     userModel.Email,
		Password:  userModel.Password,
		Role:      userModel.Role,
		Orders:    entity.ListOrderModelToOrderCore(userModel.Orders),
		CreatedAt: userModel.CreatedAt,
		UpdatedAt: userModel.UpdatedAt,
		DeletedAt: userModel.DeletedAt,
	}
	return userCore
}

func ListUserModelToUserCore(userModel []model.User) []User {
	listUserCore := []User{}
	for _, user := range userModel {
		userCore := UserModelToUserCore(user)
		listUserCore = append(listUserCore, userCore)
	}
	return listUserCore
}
