package repository

import (
	"errors"
	"final-project/modules/photo/entity"
	"final-project/modules/photo/model"

	"gorm.io/gorm"
)

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) entity.PhotoRepositoryInterface {
	return &photoRepository{
		db: db,
	}
}

func (pr *photoRepository) CreatePhoto(photoCore entity.Photo) (entity.Photo, error) {
	photoModelRequest := entity.PhotoCoreToPhotoModel(photoCore)

	tx := pr.db.Create(&photoModelRequest)
	if tx.Error != nil {
		return entity.Photo{}, tx.Error
	}

	photoCoreResponse := entity.PhotoModelToPhotoCore(photoModelRequest)
	return photoCoreResponse, nil
}
func (pr *photoRepository) GetAllPhotos(userID string) ([]entity.Photo, error) {
	photoModel := []model.Photo{}

	tx := pr.db.Model(&photoModel).Where("user_id = ?", userID).Preload("Comments").Find(&photoModel)
	if tx.Error != nil {
		return nil, tx.Error
	}

	photoCoreResponse := entity.ListPhotoModelToPhotoCore(photoModel)
	return photoCoreResponse, nil
}

func (pr *photoRepository) GetPhotoByID(photoID string) (entity.Photo, error) {

	photoModel := model.Photo{}

	tx := pr.db.Where("id = ?", photoID).Preload("Comments").First(&photoModel)
	if tx.Error != nil {
		return entity.Photo{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.Photo{}, errors.New("data not found")
	}

	photoCoreResponse := entity.PhotoModelToPhotoCore(photoModel)
	return photoCoreResponse, nil
}

func (pr *photoRepository) UpdatePhotoByID(photoID string, photoCore entity.Photo) (entity.Photo, error) {
	photoModel := model.Photo{}

	tx := pr.db.Where("id = ?", photoID).Preload("Comments").First(&photoModel)
	if tx.Error != nil {
		return entity.Photo{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.Photo{}, errors.New("data not found")
	}

	updatedPhotoModel := entity.PhotoCoreToPhotoModel(photoCore)

	tx = pr.db.Model(&photoModel).Updates(updatedPhotoModel)
	if tx.Error != nil {
		return entity.Photo{}, tx.Error
	}

	photoCoreResponse := entity.PhotoModelToPhotoCore(photoModel)
	return photoCoreResponse, nil
}

func (pr *photoRepository) DeletePhotoByID(photoID string) error {
	photoModel := model.Photo{}

	tx := pr.db.Where("id = ?", photoID).Delete(&photoModel)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("data not found")
	}
	return nil
}
