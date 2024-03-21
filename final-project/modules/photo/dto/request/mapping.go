package request

import "final-project/modules/photo/entity"

func PhotoRequestToPhotoCore(request PhotoRequest) entity.Photo {
	return entity.Photo{
		PhotoURL: request.PhotoURL,
		Title:    request.Title,
		Caption:  request.Caption,
	}
}
