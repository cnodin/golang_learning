package model

type LoginRequest struct {
	AppID uint8 `json:"appid"`
	Captcha string `json:"captcha"`
	Cellphone string `json:"cellphone"`
	Country uint8 `json:"country"`
	Password string `json:"password"`
	Platform uint8 `json:"platform"`
	Remember uint8 `json:remember`
}
