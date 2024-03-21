package entity

import (
	"final-project/modules/media-social/model"
)

func MediaSocialCoreToMediaSocialModel(mediaSocialCore MediaSocial) model.MediaSocial {
	mediaSocialModel := model.MediaSocial{
		ID:             mediaSocialCore.ID,
		Name:           mediaSocialCore.Name,
		MediaSocialURL: mediaSocialCore.MediaSocialURL,
		UserID:         mediaSocialCore.UserID,
		CreatedAt:      mediaSocialCore.CreatedAt,
		UpdatedAt:      mediaSocialCore.UpdatedAt,
		DeletedAt:      mediaSocialCore.DeletedAt,
	}
	return mediaSocialModel
}

func ListMediaSocialCoreToMediaSocialModel(mediaSocialCore []MediaSocial) []model.MediaSocial {
	listMediaSocialModel := make([]model.MediaSocial, len(mediaSocialCore))
	for i, mediaSocial := range mediaSocialCore {
		listMediaSocialModel[i] = MediaSocialCoreToMediaSocialModel(mediaSocial)
	}
	return listMediaSocialModel
}

func MediaSocialModelToMediaSocialCore(mediaSocialModel model.MediaSocial) MediaSocial {
	mediaSocialCore := MediaSocial{
		ID:             mediaSocialModel.ID,
		Name:           mediaSocialModel.Name,
		MediaSocialURL: mediaSocialModel.MediaSocialURL,
		UserID:         mediaSocialModel.UserID,
		CreatedAt:      mediaSocialModel.CreatedAt,
		UpdatedAt:      mediaSocialModel.UpdatedAt,
		DeletedAt:      mediaSocialModel.DeletedAt,
	}
	return mediaSocialCore
}

func ListMediaSocialModelToMediaSocialCore(mediaSocialModel []model.MediaSocial) []MediaSocial {
	listMediaSocialCore := make([]MediaSocial, len(mediaSocialModel))
	for i, mediaSocial := range mediaSocialModel {
		listMediaSocialCore[i] = MediaSocialModelToMediaSocialCore(mediaSocial)
	}
	return listMediaSocialCore
}
