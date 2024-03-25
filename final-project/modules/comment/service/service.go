package service

import (
	"errors"
	"final-project/modules/comment/entity"
	"final-project/utils/validator"
)

type commentService struct {
	commentRepository entity.CommentRepositoryInterface
}

func NewCommentService(commentRepository entity.CommentRepositoryInterface) entity.CommentServiceInterface {
	return &commentService{
		commentRepository: commentRepository,
	}
}

func (cs *commentService) CreateComment(photoID string, commentCore entity.Comment) (entity.Comment, error) {

	errEmpty := validator.IsDataEmpty([]string{"message"}, commentCore.Message)
	if errEmpty != nil {
		return entity.Comment{}, errEmpty
	}	

	errLength := validator.IsMaxLengthValid(100, map[string]string{"message":commentCore.Message})
	if errLength != nil {
		return entity.Comment{}, errLength
	}

	commentsData, errCreate := cs.commentRepository.CreateComment(photoID, commentCore)
	if errCreate != nil {
		return entity.Comment{}, errCreate
	}

	return commentsData, nil
}

func (cs *commentService) GetAllComments(userID string) ([]entity.Comment, error) {
	commentsData, errGetAll := cs.commentRepository.GetAllComments(userID)
	if errGetAll != nil {
		return nil, errGetAll
	}

	return commentsData, nil
}

func (cs *commentService) GetCommentByID(commentID string) (entity.Comment, error) {
	if commentID == "" {
		return entity.Comment{}, errors.New("invalid comment id")
	}

	commentsData, errGetID := cs.commentRepository.GetCommentByID(commentID)
	if errGetID != nil {
		return entity.Comment{}, errors.New("comments data empty")
	}

	return commentsData, nil
}

func (cs *commentService) UpdateCommentByID(commentID string, commentCore entity.Comment) (entity.Comment, error) {
	if commentID == "" {
		return entity.Comment{}, errors.New("invalid comment id")
	}

	_, errGetID := cs.commentRepository.GetCommentByID(commentID)
	if errGetID != nil {
		return entity.Comment{}, errors.New("comments data empty")
	}

	commentsData, errUpdate := cs.commentRepository.UpdateCommentByID(commentID, commentCore)
	if errUpdate != nil {
		return entity.Comment{}, errUpdate
	}

	return commentsData, nil
}

func (cs *commentService) DeleteCommentByID(commentID string) error {
	if commentID == "" {
		return errors.New("invalid comment id")
	}

	errDelete := cs.commentRepository.DeleteCommentByID(commentID)
	if errDelete != nil {
		return errDelete
	}

	return nil
}
