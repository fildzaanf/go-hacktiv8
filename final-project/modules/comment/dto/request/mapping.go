package request

import "final-project/modules/comment/entity"

func CommentRequestToCommentCore(request CommentRequest) entity.Comment {
	return entity.Comment{
		Message: request.Message,
	}
}
