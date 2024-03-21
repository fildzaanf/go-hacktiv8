package request

type UserRegisterRequest struct {
	Fullname string `json:"fullname" form:"fullname"`
	Username string `json:"username" form:"username"`
	Age      int    `json:"age" form:"age"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserLoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserRequest struct {
	Fullname string `json:"fullname" form:"fullname"`
	Username string `json:"username" form:"username"`
	Age      int    `json:"age" form:"age"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
