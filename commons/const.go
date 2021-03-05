package commons

import "github.com/mojocn/base64Captcha"

const SECRET = "xiao_qiang_music"

var Store = base64Captcha.DefaultMemStore

const (
	PWD_SECRET_KEY = "user_pwd"
	PROJECT_NAME   = "xq_music_go"
	MUSIC_PAYH     = "/music/"
	MAPPING_Path   = "/static"
)
const (
	//普通用户
	DEFAULT_ROLE_NAME = "user"
	//ID 1
	DEFAULT_ROLE_ID = 1
)
