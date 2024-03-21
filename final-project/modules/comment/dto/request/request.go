package request

type CommentRequest struct {
	Message string `json:"message" form:"message" example:"this is a sample comment message"`
}
