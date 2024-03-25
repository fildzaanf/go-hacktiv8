package service

import (
	"errors"
	"final-project/modules/photo/entity"
	"final-project/utils/validator"
)

type photoService struct {
	photoRepository entity.PhotoRepositoryInterface
}

func NewPhotoService(photoRepository entity.PhotoRepositoryInterface) entity.PhotoServiceInterface {
	return &photoService{
		photoRepository: photoRepository,
	}
}

func (ps *photoService) CreatePhoto(photoCore entity.Photo) (entity.Photo, error) {

	errEmpty := validator.IsDataEmpty([]string{"photo_url", "title", "caption"}, photoCore.PhotoURL, photoCore.Title, photoCore.Caption)
	if errEmpty != nil {
		return entity.Photo{}, errEmpty
	}
	

	errLength := validator.IsMaxLengthValid(100, map[string]string{"caption": photoCore.Caption})
	if errLength != nil {
		return entity.Photo{}, errLength
	}

	photosData, errCreate := ps.photoRepository.CreatePhoto(photoCore)
	if errCreate != nil {
		return entity.Photo{}, errCreate
	}

	return photosData, nil
}

func (ps *photoService) GetAllPhotos(userID string) ([]entity.Photo, error) {
	photosData, errGetAll := ps.photoRepository.GetAllPhotos(userID)
	if errGetAll != nil {
		return nil, errGetAll
	}

	return photosData, nil
}

func (ps *photoService) GetPhotoByID(photoID string) (entity.Photo, error) {
	if photoID == "" {
		return entity.Photo{}, errors.New("invalid photo id")
	}

	photosData, errGetID := ps.photoRepository.GetPhotoByID(photoID)
	if errGetID != nil {
		return entity.Photo{}, errors.New("photos data empty")
	}

	return photosData, nil
}

func (ps *photoService) UpdatePhotoByID(photoID string, photoCore entity.Photo) (entity.Photo, error) {
	if photoID == "" {
		return entity.Photo{}, errors.New("invalid photo id")
	}

	_, errGetID := ps.photoRepository.GetPhotoByID(photoID)
	if errGetID != nil {
		return entity.Photo{}, errors.New("photos data empty")
	}

	photosData, errUpdate := ps.photoRepository.UpdatePhotoByID(photoID, photoCore)
	if errUpdate != nil {
		return entity.Photo{}, errUpdate
	}

	return photosData, nil
}

func (ps *photoService) DeletePhotoByID(photoID string) error {
	if photoID == "" {
		return errors.New("invalid photo id")
	}

	errDelete := ps.photoRepository.DeletePhotoByID(photoID)
	if errDelete != nil {
		return errDelete
	}

	return nil
}
