package request

import "final-project/modules/media-social/entity"

func MediaSocialRequestToMediaSocialCore(request MediaSocialRequest) entity.MediaSocial {
	return entity.MediaSocial{
		Name:           request.Name,
		MediaSocialURL: request.MediaSocialURL,
	}
}
