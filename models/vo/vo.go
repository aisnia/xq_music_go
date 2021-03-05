package vo

//CaptchaBody 获取验证码的请求正文
type CaptchaBody struct {
	Id          string `json:"id" form:"id" query:"id" query:"id" validate:"required"`
	VerifyValue string `json:"verify_value" form:"verify_value" query:"verify_value" validate:"required"`
}

//注册表单
type RegisterForm struct {
	CId      string `json:"c_id" form:"c_id" query:"c_id" validate:"required"`
	Username string `json:"username" form:"username" query:"username" validate:"required"`
	Name     string `json:"name" form:"name" query:"name" validate:"required"`
	Sex      int8   `json:"sex" form:"sex" query:"sex" validate:"required"`
	Phone    string `json:"phone" form:"phone" query:"phone" validate:"required"`
	Email    string `json:"email" form:"email" query:"email" validate:"required"`
	Password string `json:"password" form:"password" query:"password" validate:"required"`
	Code     string `json:"code" form:"code" query:"code" validate:"required"`
}
