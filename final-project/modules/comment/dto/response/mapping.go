package response

import "final-project/modules/comment/entity"

func CommentCoreToCommentResponse(commentCore entity.Comment) CommentResponse {
	return CommentResponse{
		ID:        commentCore.ID,
		UserID:    commentCore.UserID,
		PhotoID:   commentCore.PhotoID,
		Message:   commentCore.Message,
		CreatedAt: commentCore.CreatedAt,
		UpdatedAt: commentCore.UpdatedAt,
		DeletedAt: commentCore.DeletedAt,
	}
}

func ListCommentCoreToCommentResponse(commentCore []entity.Comment) []CommentResponse {
	commentResponses := []CommentResponse{}
	for _, comment := range commentCore {
		commentResponse := CommentCoreToCommentResponse(comment)
		commentResponses = append(commentResponses, commentResponse)
	}
	return commentResponses
}
