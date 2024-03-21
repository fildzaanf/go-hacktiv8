package request

type UserRegisterRequest struct {
	Fullname string `json:"fullname" form:"fullname" example:"hanisah fildza annafisah"`
	Username string `json:"username" form:"username" example:"fildzaanf"`
	Age      int    `json:"age" form:"age" example:"20"`
	Email    string `json:"email" form:"email" example:"hanisahfildza@gmail.com"`
	Password string `json:"password" form:"password" example:"password123456789"`
}

type UserLoginRequest struct {
	Email    string `json:"email" form:"email" example:"hanisahfildza@gmail.com"`
	Password string `json:"password" form:"password" example:"password123456789"`
}

type UserRequest struct {
	Fullname string `json:"fullname" form:"fullname" example:"hanisah fildza annafisah"`
	Username string `json:"username" form:"username" example:"fildzaanf"`
	Age      int    `json:"age" form:"age" example:"20"`
	Email    string `json:"email" form:"email" example:"hanisahfildza@gmail.com"`
	Password string `json:"password" form:"password" example:"password123456789"`
}
