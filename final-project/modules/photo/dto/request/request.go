package request

type PhotoRequest struct {
	PhotoURL string `json:"photo_url" form:"photo_url" example:"http://example.com/photo.jpg"`
	Title    string `json:"title" form:"title" example:"title photo"`
	Caption  string `json:"caption" form:"caption" example:"this is example caption photo"`
}

type PhotoUpdateRequest struct {
	Title    string `json:"title" form:"title" example:"title photo"`
	Caption  string `json:"caption" form:"caption" example:"this is example caption photo"`
}