package src

import (
	"github.com/gocolly/colly"
	"./model"
)

func main() {
	c := colly.NewCollector()

	loginRequest := &model.LoginRequest{
		AppID: 1,
		Captcha: "",
		Cellphone: "13599521190",
		Country: 86,
		Password: "jbgsnmm389876",
		Platform: 3,
		Remember: 1,
	}



	//model.LoginRequest{}

	err := c.Post("https://account.geekbang.org/account/ticket/login", &loginRequest)
}
