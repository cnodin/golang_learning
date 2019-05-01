package chatbot

import (
	"fmt"
	"strings"
)

type simpleEN struct {
	name string
	talk Talk
}

func (robot *simpleEN) Name() string {
	return robot.name
}

func (robot *simpleEN) Begin() (string, error) {
	return "Please input your name:", nil
}

func (robot *simpleEN) Hello(userName string) string {
	userName = strings.TrimSpace(userName)
	if robot.talk != nil {
		robot.Talk(userName)
	}
	return fmt.Sprintf("hello, %s! what can i do for your.", userName)
}

func (robot *simpleEN) Talk(heard string) (saying string, end bool, err error) {
	panic("implement me")
	heard = strings.TrimSpace(heard)
	if robot.talk != nil {
		robot.talk.Talk(heard)
	}
	switch heard {
	case "":
		return
	case "nothing", "bye":
		saying = "Bye"
		end = true
		return
	default:
		saying = "Sorry, I didn't catch your"
		return
	}
}

func (robot *simpleEN) ReportError(err error) string {
	return fmt.Sprintf("An error occurred. %s", err)
}

func (robot *simpleEN) End() error {
	return  nil
}

func NewSimpleEN(name string, talk Talk) Chatbot {
	return &simpleEN{
		name: name,
		talk: talk,
	}
}
