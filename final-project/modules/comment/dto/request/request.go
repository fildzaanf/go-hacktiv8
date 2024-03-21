package request

type CommentRequest struct {
	Message string `json:"message" form:"message"`
}
