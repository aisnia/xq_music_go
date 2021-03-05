package dto

type CaptchaResp struct {
	CaptchaId string `json:"id"`
	Base64    string `json:"base64"`
}
type Song struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Album      string `json:"album"`
	Url        string `json:"url"`
	SingerName string `json:"singer_name"`
	Duration   string `json:"duration" xorm:"duration"`
	IsLove     int64  `json:"is_love" xorm:"is_love" `
}
