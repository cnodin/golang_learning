package chatbot

import "github.com/pkg/errors"

type Talk interface {
	Hello(userName string) string
	Talk(heard string) (saying string, end bool, err error)
}

type Chatbot interface {
	Name() string
	Begin() (string, error)
	Talk
	ReportError(err error) string
	End() error
}

var (
	//无效的机器人名称
	ErrInvalidChatbotName = errors.New("Invalid chatbot name")
	//无效的机器人
	ErrInvalidChatbot = errors.New("Invalid chatbot")
	//机器人重复注册
	ErrExistingChatbot = errors.New("Existing chatbot")
)

var chatbotMap = map[string]Chatbot{}

//用于注册机器人
func Register(chatbot Chatbot) error {
	if chatbot == nil {
		return ErrInvalidChatbot
	}
	name := chatbot.Name()
	if name == "" {
		return ErrInvalidChatbotName
	}
	if _, ok := chatbotMap[name]; ok {
		return ErrExistingChatbot
	}
	chatbotMap[name] = chatbot
	return nil
}

func Get(name string) Chatbot {
	return chatbotMap[name]
}


