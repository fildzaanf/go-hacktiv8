package entity

import (
	ec "final-project/modules/comment/entity"
	"final-project/modules/photo/model"
)

func PhotoCoreToPhotoModel(photoCore Photo) model.Photo {
	photoModel := model.Photo{
		ID:        photoCore.ID,
		PhotoURL:  photoCore.PhotoURL,
		Title:     photoCore.Title,
		Caption:   photoCore.Caption,
		UserID:    photoCore.UserID,
		Comments:  ec.ListCommentCoreToCommentModel(photoCore.Comments),
		CreatedAt: photoCore.CreatedAt,
		UpdatedAt: photoCore.UpdatedAt,
		DeletedAt: photoCore.DeletedAt,
	}
	return photoModel
}

func ListPhotoCoreToPhotoModel(photoCore []Photo) []model.Photo {
	listPhotoModel := make([]model.Photo, len(photoCore))
	for i, photo := range photoCore {
		listPhotoModel[i] = PhotoCoreToPhotoModel(photo)
	}
	return listPhotoModel
}

func PhotoModelToPhotoCore(photoModel model.Photo) Photo {
	photoCore := Photo{
		ID:        photoModel.ID,
		PhotoURL:  photoModel.PhotoURL,
		Title:     photoModel.Title,
		Caption:   photoModel.Caption,
		UserID:    photoModel.UserID,
		Comments:  ec.ListCommentModelToCommentCore(photoModel.Comments),
		CreatedAt: photoModel.CreatedAt,
		UpdatedAt: photoModel.UpdatedAt,
		DeletedAt: photoModel.DeletedAt,
	}
	return photoCore
}

func ListPhotoModelToPhotoCore(photoModel []model.Photo) []Photo {
	listPhotoCore := make([]Photo, len(photoModel))
	for i, photo := range photoModel {
		listPhotoCore[i] = PhotoModelToPhotoCore(photo)
	}
	return listPhotoCore
}
