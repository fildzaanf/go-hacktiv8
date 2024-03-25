package request

type UserRegisterRequest struct {
	Email           string `json:"email" form:"email" example:"hanisahfildza@gmail.com"`
	Password        string `json:"password" form:"password" example:"password123456789"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"  example:"password123456789"`
}

type UserLoginRequest struct {
	Email    string `json:"email" form:"email" example:"hanisahfildza@gmail.com"`
	Password string `json:"password" form:"password" example:"password123456789"`
}

type UserUpdateProfileRequest struct {
	Fullname        string `json:"fullname" form:"fullname" example:"hanisah fildza annafisah"`
	Username        string `json:"username" form:"username" example:"fildzaanf"`
	Age             int    `json:"age" form:"age" example:"20"`
	Email           string `json:"email" form:"email" example:"hanisahfildza@gmail.com"`
	Password        string `json:"password" form:"password" example:"password123456789"`
	NewPassword     string `json:"new_password" form:"new_password"  example:"password12345678910"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"  example:"password12345678910"`
}

type UserNewPasswordRequest struct {
	Password        string `json:"password" form:"password" example:"password123456789"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"  example:"password123456789"`
}

type UserSendOTPRequest struct {
	Email string `json:"email" form:"email" example:"hanisahfildza@gmail.com"`
}

type UserVerifyOTPRequest struct {
	Email string `json:"email" form:"email" example:"hanisahfildza@gmail.com"`
	OTP   string `json:"otp" form:"otp" example:"7777"`
}
