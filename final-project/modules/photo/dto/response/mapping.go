package response

import (
	rc "final-project/modules/comment/dto/response"
	"final-project/modules/photo/entity"
)

func PhotoCoreToPhotoResponse(photoCore entity.Photo) PhotoResponse {
	return PhotoResponse{
		ID:        photoCore.ID,
		PhotoURL:  photoCore.PhotoURL,
		Title:     photoCore.Title,
		Caption:   photoCore.Caption,
		UserID:    photoCore.UserID,
		Comments:  rc.ListCommentCoreToCommentResponse(photoCore.Comments),
		CreatedAt: photoCore.CreatedAt,
		UpdatedAt: photoCore.UpdatedAt,
		DeletedAt: photoCore.DeletedAt,
	}
}

func ListPhotoCoreToPhotoResponse(photos []entity.Photo) []PhotoResponse {
    photoResponses := []PhotoResponse{}
    for _, photo := range photos {
        photoResponse := PhotoCoreToPhotoResponse(photo)
        photoResponses = append(photoResponses, photoResponse)
    }
    return photoResponses
}
