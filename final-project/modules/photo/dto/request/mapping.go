package request

import "final-project/modules/photo/entity"

func PhotoRequestToPhotoCore(request PhotoRequest) entity.Photo {
	return entity.Photo{
		PhotoURL: request.PhotoURL,
		Title:    request.Title,
		Caption:  request.Caption,
	}
}

func PhotoUpdateRequestToPhotoCore(request PhotoUpdateRequest) entity.Photo {
	return entity.Photo{
		Title:    request.Title,
		Caption:  request.Caption,
	}
}

