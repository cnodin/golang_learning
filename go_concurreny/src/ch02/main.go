package main

import (
	"bufio"
	"ch02/chatbot"
	"flag"
	"fmt"
	"os"
	"runtime/debug"
)

var chatbotName string

func init() {
	flag.StringVar(&chatbotName, "chatbot", "simple.cn", "The chatbot's name for dialogue.")
}

func main() {
	flag.Parse()
	chatbot.Register(chatbot.NewSimpleCN("simple.cn", nil))
	chatbot.Register(chatbot.NewSimpleEN("simple.en", nil))

	mybot := chatbot.Get(chatbotName)
	if mybot == nil {
		err := fmt.Errorf("Fatal Error: ")
		checkError(nil, err, true)
	}
	inputReader := bufio.NewReader(os.Stdin)
	begin, err := mybot.Begin()
	checkError(mybot, err, true)
	fmt.Println(begin)
	input, err := inputReader.ReadString('\n')
	checkError(mybot, err, true)
	fmt.Println(mybot.Hello(input[:len(input) - 1]))

	for {
		input, err = inputReader.ReadString('\n')
		if checkError(mybot, err, false) {
			continue
		}
		output, end, err := mybot.Talk(input)
		if checkError(mybot, err, false) {
			continue
		}
		if output != "" {
			fmt.Println(output)
		}

		if end {
			err = mybot.End()
			checkError(mybot, err, false)
			os.Exit(0)
		}
	}
}

func checkError(chatbot chatbot.Chatbot, err error, exit bool) bool {
	if err == nil {
		return false
	}

	if chatbot != nil {
		fmt.Println(chatbot.ReportError(err))
	} else {
		fmt.Println(err)
	}

	if exit {
		debug.PrintStack()
		os.Exit(1)
	}
	return true
}
