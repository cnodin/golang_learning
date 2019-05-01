package main

type Talk interface {
	Hello(userName string) string
	Talk(heard string) (saying string, end bool, err error)
}

type simpleCN struct {
	name string
	talk Talk
}

type Chatbot interface {
	Name() string
	Begin() (string, error)
	Talk
	ReportError(err error) string
	End() error
}
