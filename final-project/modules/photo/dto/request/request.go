package request

type PhotoRequest struct {
	PhotoURL string `json:"photo_url" form:"photo_url"`
	Title    string `json:"title" form:"title"`
	Caption  string `json:"caption" form:"caption"`
}
