package entity

import (
	ec "final-project/modules/comment/entity"
	em "final-project/modules/media-social/entity"
	ep "final-project/modules/photo/entity"
	"final-project/modules/user/model"
)

func UserCoreToUserModel(userCore User) model.User {
	userModel := model.User{
		ID:            userCore.ID,
		Fullname:      userCore.Fullname,
		Email:         userCore.Email,
		Password:      userCore.Password,
		Role:          userCore.Role,
		Username:      userCore.Username,
		Age:           userCore.Age,
		OTP:           userCore.OTP,
		OTPExpiration: userCore.OTPExpiration,
		CreatedAt:     userCore.CreatedAt,
		UpdatedAt:     userCore.UpdatedAt,
		DeletedAt:     userCore.DeletedAt,
		Photos:        ep.ListPhotoCoreToPhotoModel(userCore.Photos),
		Comments:      ec.ListCommentCoreToCommentModel(userCore.Comments),
		MediaSocials:  em.ListMediaSocialCoreToMediaSocialModel(userCore.MediaSocials),
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
		ID:            userModel.ID,
		Fullname:      userModel.Fullname,
		Email:         userModel.Email,
		Password:      userModel.Password,
		Role:          userModel.Role,
		Age:           userModel.Age,
		Username:      userModel.Username,
		OTP:           userModel.OTP,
		OTPExpiration: userModel.OTPExpiration,
		CreatedAt:     userModel.CreatedAt,
		UpdatedAt:     userModel.UpdatedAt,
		DeletedAt:     userModel.DeletedAt,
		Photos:        ep.ListPhotoModelToPhotoCore(userModel.Photos),
		Comments:      ec.ListCommentModelToCommentCore(userModel.Comments),
		MediaSocials:  em.ListMediaSocialModelToMediaSocialCore(userModel.MediaSocials),
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
