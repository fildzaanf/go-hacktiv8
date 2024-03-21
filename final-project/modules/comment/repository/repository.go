package repository

import (
	"errors"
	"final-project/modules/comment/entity"
	"final-project/modules/comment/model"
	ep "final-project/modules/photo/entity"

	"gorm.io/gorm"
)

type commentRepository struct {
	db              *gorm.DB
	photoRepository ep.PhotoRepositoryInterface
}

func NewCommentRepository(db *gorm.DB, photoRepository ep.PhotoRepositoryInterface) entity.CommentRepositoryInterface {
	return &commentRepository{
		db:              db,
		photoRepository: photoRepository,
	}
}

func (cr *commentRepository) CreateComment(photoID string, commentCore entity.Comment) (entity.Comment, error) {

	photo, err := cr.photoRepository.GetPhotoByID(photoID)
	if err != nil {
		return entity.Comment{}, err
	}

	commentModel := entity.CommentCoreToCommentModel(commentCore)

	commentModel.PhotoID = photo.ID

	tx := cr.db.Create(&commentModel)
	if tx.Error != nil {
		return entity.Comment{}, tx.Error
	}

	commentCoreResponse := entity.CommentModelToCommentCore(commentModel)
	return commentCoreResponse, nil
}

func (cr *commentRepository) GetAllComments(userID string) ([]entity.Comment, error) {
	commentModel := []model.Comment{}

	tx := cr.db.Model(&commentModel).Where("user_id = ?", userID).Find(&commentModel)
	if tx.Error != nil {
		return nil, tx.Error
	}

	commentCoreResponse := entity.ListCommentModelToCommentCore(commentModel)
	return commentCoreResponse, nil
}

func (cr *commentRepository) GetCommentByID(commentID string) (entity.Comment, error) {
	commentModel := model.Comment{}

	tx := cr.db.Where("id = ?", commentID).First(&commentModel)
	if tx.Error != nil {
		return entity.Comment{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.Comment{}, errors.New("photo data not found")
	}

	commentCoreResponse := entity.CommentModelToCommentCore(commentModel)
	return commentCoreResponse, nil
}

func (cr *commentRepository) UpdateCommentByID(commentID string, commentCore entity.Comment) (entity.Comment, error) {
	commentModel := model.Comment{}

	tx := cr.db.Where("id = ?", commentID).First(&commentModel)
	if tx.Error != nil {
		return entity.Comment{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.Comment{}, errors.New("photo data not found")
	}

	updatedCommentModel := entity.CommentCoreToCommentModel(commentCore)

	tx = cr.db.Model(&commentModel).Updates(updatedCommentModel)
	if tx.Error != nil {
		return entity.Comment{}, tx.Error
	}

	commentCoreResponse := entity.CommentModelToCommentCore(commentModel)
	return commentCoreResponse, nil
}

func (cr *commentRepository) DeleteCommentByID(commentID string) error {
	commentModel := model.Comment{}

	tx := cr.db.Where("id = ?", commentID).Delete(&commentModel)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("photo data not found")
	}
	return nil
}
