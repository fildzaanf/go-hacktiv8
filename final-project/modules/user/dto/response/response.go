package response

import (
	rc "final-project/modules/comment/dto/response"
	rm "final-project/modules/media-social/dto/response"
	rp "final-project/modules/photo/dto/response"
	"time"

	"gorm.io/gorm"
)

type UserRegisterResponse struct {
	ID           string                   `json:"id"`
	Fullname     string                   `json:"fullname"`
	Username     string                   `json:"username"`
	Age          int                      `json:"age"`
	Email        string                   `json:"email"`
	Role         string                   `json:"role"`
	Photos       []rp.PhotoResponse       `json:"photos"`
	Comments     []rc.CommentResponse     `json:"comments"`
	MediaSocials []rm.MediaSocialResponse `json:"media_socials"`
	CreatedAt    time.Time                `json:"created_at"`
	UpdatedAt    time.Time                `json:"updated_at"`
	DeletedAt    gorm.DeletedAt           `json:"deleted_at"`
}

type UserLoginResponse struct {
	ID           string                   `json:"id"`
	Fullname     string                   `json:"fullname"`
	Username     string                   `json:"username"`
	Age          int                      `json:"age"`
	Email        string                   `json:"email"`
	Role         string                   `json:"role"`
	Token        string                   `json:"token"`
	CreatedAt    time.Time                `json:"created_at"`
	UpdatedAt    time.Time                `json:"updated_at"`
	DeletedAt    gorm.DeletedAt           `json:"deleted_at"`
	Photos       []rp.PhotoResponse       `json:"photos"`
	Comments     []rc.CommentResponse     `json:"comments"`
	MediaSocials []rm.MediaSocialResponse `json:"media_socials"`
}

type UserResponse struct {
	ID           string                   `json:"id"`
	Fullname     string                   `json:"fullname"`
	Username     string                   `json:"username"`
	Age          int                      `json:"age"`
	Email        string                   `json:"email"`
	Role         string                   `json:"role"`
	CreatedAt    time.Time                `json:"created_at"`
	UpdatedAt    time.Time                `json:"updated_at"`
	DeletedAt    gorm.DeletedAt           `json:"deleted_at"`
	Photos       []rp.PhotoResponse       `json:"photos"`
	Comments     []rc.CommentResponse     `json:"comments"`
	MediaSocials []rm.MediaSocialResponse `json:"media_socials"`
}

type UserVerifyOTPResponse struct {
	Token string `json:"token"`
}
