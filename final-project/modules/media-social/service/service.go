package service

import (
	"errors"
	"final-project/modules/media-social/entity"
	"final-project/utils/validator"
)

type mediaSocialService struct {
	mediaSocialRepository entity.MediaSocialRepositoryInterface
}

func NewMediaSocialService(mediaSocialRepository entity.MediaSocialRepositoryInterface) entity.MediaSocialServiceInterface {
	return &mediaSocialService{
		mediaSocialRepository: mediaSocialRepository,
	}
}

func (ms *mediaSocialService) CreateMediaSocial(mediaSocialCore entity.MediaSocial) (entity.MediaSocial, error) {

	errEmpty := validator.IsDataEmpty([]string{"name", "media_social_url"}, mediaSocialCore.Name, mediaSocialCore.MediaSocialURL)
	if errEmpty != nil {
		return entity.MediaSocial{}, errEmpty
	}
	
	errLength := validator.IsMaxLengthValid(30, map[string]string{"name": mediaSocialCore.Name})
	if errLength != nil {
		return entity.MediaSocial{}, errLength
	}

	mediaSocialData, errCreate := ms.mediaSocialRepository.CreateMediaSocial(mediaSocialCore)
	if errCreate != nil {
		return entity.MediaSocial{}, errCreate
	}

	return mediaSocialData, nil
}

func (ms *mediaSocialService) GetAllMediaSocials(userID string) ([]entity.MediaSocial, error) {
	mediaSocialData, errGetAll := ms.mediaSocialRepository.GetAllMediaSocials(userID)
	if errGetAll != nil {
		return nil, errGetAll
	}

	return mediaSocialData, nil
}

func (ms *mediaSocialService) GetMediaSocialByID(mediaSocialID string) (entity.MediaSocial, error) {
	if mediaSocialID == "" {
		return entity.MediaSocial{}, errors.New("invalid media social id")
	}

	mediaSocialData, errGetID := ms.mediaSocialRepository.GetMediaSocialByID(mediaSocialID)
	if errGetID != nil {
		return entity.MediaSocial{}, errors.New("media social data empty")
	}

	return mediaSocialData, nil
}

func (ms *mediaSocialService) UpdateMediaSocialByID(mediaSocialID string, mediaSocialCore entity.MediaSocial) (entity.MediaSocial, error) {
	if mediaSocialID == "" {
		return entity.MediaSocial{}, errors.New("invalid media social id")
	}

	_, errGetID := ms.mediaSocialRepository.GetMediaSocialByID(mediaSocialID)
	if errGetID != nil {
		return entity.MediaSocial{}, errors.New("media social data empty")
	}

	mediaSocialData, errUpdate := ms.mediaSocialRepository.UpdateMediaSocialByID(mediaSocialID, mediaSocialCore)
	if errUpdate != nil {
		return entity.MediaSocial{}, errUpdate
	}

	return mediaSocialData, nil
}

func (ms *mediaSocialService) DeleteMediaSocialByID(mediaSocialID string) error {
	if mediaSocialID == "" {
		return errors.New("invalid media social id")
	}

	errDelete := ms.mediaSocialRepository.DeleteMediaSocialByID(mediaSocialID)
	if errDelete != nil {
		return errDelete
	}

	return nil
}
