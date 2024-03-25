package response

import (
	rc "final-project/modules/comment/dto/response"
	rm "final-project/modules/media-social/dto/response"
	rp "final-project/modules/photo/dto/response"
	"final-project/modules/user/entity"
)

func UserCoreToUserRegisterResponse(userCore entity.User) UserRegisterResponse {
	return UserRegisterResponse{
		ID:           userCore.ID,
		Fullname:     userCore.Fullname,
		Username:     userCore.Username,
		Age:          userCore.Age,
		Email:        userCore.Email,
		Role:         userCore.Role,
		CreatedAt:    userCore.CreatedAt,
		UpdatedAt:    userCore.UpdatedAt,
		DeletedAt:    userCore.DeletedAt,
		Photos:       rp.ListPhotoCoreToPhotoResponse(userCore.Photos),
		Comments:     rc.ListCommentCoreToCommentResponse(userCore.Comments),
		MediaSocials: rm.ListMediaSocialCoreToMediaSocialResponse(userCore.MediaSocials),
	}
}

func UserCoreToUserLoginResponse(userCore entity.User, token string) UserLoginResponse {
	return UserLoginResponse{
		ID:           userCore.ID,
		Fullname:     userCore.Fullname,
		Username:     userCore.Username,
		Age:          userCore.Age,
		Email:        userCore.Email,
		Role:         userCore.Role,
		CreatedAt:    userCore.CreatedAt,
		UpdatedAt:    userCore.UpdatedAt,
		DeletedAt:    userCore.DeletedAt,
		Token:        token,
		Photos:       rp.ListPhotoCoreToPhotoResponse(userCore.Photos),
		Comments:     rc.ListCommentCoreToCommentResponse(userCore.Comments),
		MediaSocials: rm.ListMediaSocialCoreToMediaSocialResponse(userCore.MediaSocials),
	}
}

func UserCoreToUserResponse(userCore entity.User) UserResponse {
	return UserResponse{
		ID:           userCore.ID,
		Fullname:     userCore.Fullname,
		Username:     userCore.Username,
		Age:          userCore.Age,
		Email:        userCore.Email,
		Role:         userCore.Role,
		CreatedAt:    userCore.CreatedAt,
		UpdatedAt:    userCore.UpdatedAt,
		DeletedAt:    userCore.DeletedAt,
		Photos:       rp.ListPhotoCoreToPhotoResponse(userCore.Photos),
		Comments:     rc.ListCommentCoreToCommentResponse(userCore.Comments),
		MediaSocials: rm.ListMediaSocialCoreToMediaSocialResponse(userCore.MediaSocials),
	}
}

func UserCoreToUserVerifyOTPResponse(token string) UserVerifyOTPResponse {
	return UserVerifyOTPResponse{
		Token: token,
	}
}

func ListUserCoreToUserResponse(userCores []entity.User) []UserResponse {
	var userResponses []UserResponse
	for _, userCore := range userCores {
		userResponse := UserCoreToUserResponse(userCore)
		userResponses = append(userResponses, userResponse)
	}
	return userResponses
}
