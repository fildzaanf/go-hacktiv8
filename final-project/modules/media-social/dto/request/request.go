package request

type MediaSocialRequest struct {
	Name           string `json:"name" form:"name" example:"instagram"`
	MediaSocialURL string `json:"media_social_url" form:"media_social_url" example:"https://www.instagram.com/fildzaanf"`
}
