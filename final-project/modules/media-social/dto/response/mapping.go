package response

import "final-project/modules/media-social/entity"

func MediaSocialCoreToMediaSocialResponse(mediaSocialCore entity.MediaSocial) MediaSocialResponse {
	return MediaSocialResponse{
		ID:             mediaSocialCore.ID,
		Name:           mediaSocialCore.Name,
		MediaSocialURL: mediaSocialCore.MediaSocialURL,
		UserID:         mediaSocialCore.UserID,
		CreatedAt:      mediaSocialCore.CreatedAt,
		UpdatedAt:      mediaSocialCore.UpdatedAt,
		DeletedAt:      mediaSocialCore.DeletedAt,
	}
}

func ListMediaSocialCoreToMediaSocialResponse(mediaSocialCore []entity.MediaSocial) []MediaSocialResponse {
	mediaSocialResponses := []MediaSocialResponse{}
	for _, mediaSocialCore := range mediaSocialCore {
		mediaSocialResponse := MediaSocialCoreToMediaSocialResponse(mediaSocialCore)
		mediaSocialResponses = append(mediaSocialResponses, mediaSocialResponse)
	}
	return mediaSocialResponses
}