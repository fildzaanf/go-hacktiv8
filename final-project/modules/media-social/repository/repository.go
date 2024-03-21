package repository

import (
	"errors"
	"final-project/modules/media-social/entity"
	"final-project/modules/media-social/model"

	"gorm.io/gorm"
)

type mediaSocialRepository struct {
	db *gorm.DB
}

func NewMediaSocialRepository(db *gorm.DB) entity.MediaSocialRepositoryInterface {
	return &mediaSocialRepository{
		db: db,
	}
}

func (mr *mediaSocialRepository) CreateMediaSocial(mediaSocialCore entity.MediaSocial) (entity.MediaSocial, error) {
	mediaSocialModelRequest := entity.MediaSocialCoreToMediaSocialModel(mediaSocialCore)

	tx := mr.db.Create(&mediaSocialModelRequest)
	if tx.Error != nil {
		return entity.MediaSocial{}, tx.Error
	}

	mediaSocialCoreResponse := entity.MediaSocialModelToMediaSocialCore(mediaSocialModelRequest)
	return mediaSocialCoreResponse, nil
}

func (mr *mediaSocialRepository) GetAllMediaSocials(userID string) ([]entity.MediaSocial, error) {
	mediaSocialModel := []model.MediaSocial{}

	tx := mr.db.Model(&mediaSocialModel).Where("user_id = ?", userID).Find(&mediaSocialModel)
	if tx.Error != nil {
		return nil, tx.Error
	}

	mediaSocialCoreResponse := entity.ListMediaSocialModelToMediaSocialCore(mediaSocialModel)
	return mediaSocialCoreResponse, nil
}

func (mr *mediaSocialRepository) GetMediaSocialByID(mediaSocialID string) (entity.MediaSocial, error) {

	mediaSocialModel := model.MediaSocial{}

	tx := mr.db.Where("id = ?", mediaSocialID).First(&mediaSocialModel)
	if tx.Error != nil {
		return entity.MediaSocial{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.MediaSocial{}, errors.New("data not found")
	}

	mediaSocialCoreResponse := entity.MediaSocialModelToMediaSocialCore(mediaSocialModel)
	return mediaSocialCoreResponse, nil
}

func (mr *mediaSocialRepository) UpdateMediaSocialByID(mediaSocialID string, mediaSocialCore entity.MediaSocial) (entity.MediaSocial, error) {
	mediaSocialModel := model.MediaSocial{}

	tx := mr.db.Where("id = ?", mediaSocialID).First(&mediaSocialModel)
	if tx.Error != nil {
		return entity.MediaSocial{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.MediaSocial{}, errors.New("data not found")
	}

	updatedMediaSocialModel := entity.MediaSocialCoreToMediaSocialModel(mediaSocialCore)

	tx = mr.db.Model(&mediaSocialModel).Updates(updatedMediaSocialModel)
	if tx.Error != nil {
		return entity.MediaSocial{}, tx.Error
	}

	mediaSocialCoreResponse := entity.MediaSocialModelToMediaSocialCore(mediaSocialModel)
	return mediaSocialCoreResponse, nil
}

func (mr *mediaSocialRepository) DeleteMediaSocialByID(mediaSocialID string) error {
	mediaSocialModel := model.MediaSocial{}

	tx := mr.db.Where("id = ?", mediaSocialID).Delete(&mediaSocialModel)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("data not found")
	}
	return nil
}
