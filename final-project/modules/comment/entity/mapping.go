package entity

import (
	"final-project/modules/comment/model"
)

func CommentCoreToCommentModel(commentCore Comment) model.Comment {
	commentModel := model.Comment{
		ID:        commentCore.ID,
		UserID:    commentCore.UserID,
		PhotoID:   commentCore.PhotoID,
		Message:   commentCore.Message,
		CreatedAt: commentCore.CreatedAt,
		UpdatedAt: commentCore.UpdatedAt,
		DeletedAt: commentCore.DeletedAt,
	}
	return commentModel
}

func ListCommentCoreToCommentModel(commentCore []Comment) []model.Comment {
	listCommentModel := make([]model.Comment, len(commentCore))
	for i, comment := range commentCore {
		listCommentModel[i] = CommentCoreToCommentModel(comment)
	}
	return listCommentModel
}

func CommentModelToCommentCore(commentModel model.Comment) Comment {
	commentCore := Comment{
		ID:        commentModel.ID,
		UserID:    commentModel.UserID,
		PhotoID:   commentModel.PhotoID,
		Message:   commentModel.Message,
		CreatedAt: commentModel.CreatedAt,
		UpdatedAt: commentModel.UpdatedAt,
		DeletedAt: commentModel.DeletedAt,
	}
	return commentCore
}

func ListCommentModelToCommentCore(commentModel []model.Comment) []Comment {
	listCommentCore := make([]Comment, len(commentModel))
	for i, comment := range commentModel {
		listCommentCore[i] = CommentModelToCommentCore(comment)
	}
	return listCommentCore
}
