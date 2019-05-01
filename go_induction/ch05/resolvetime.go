package main

import "fmt"

const (
	//定义每分钟的秒数
	SecondsPerMinute = 60

	//定义每小时的秒数
	SecondPerHour = SecondsPerMinute * 60

	//定义每天的秒数
	SecondPerDay = SecondPerHour * 24
)

func resolveTime(seconds int) (day int, hour int, minute int) {
	day = seconds / SecondPerDay
	hour = seconds / SecondPerHour
	minute = seconds / SecondsPerMinute

	return
}

func main() {

	fmt.Println(resolveTime(1000))

	_, hour, minute := resolveTime(18000)
	fmt.Println(hour, minute)


}
