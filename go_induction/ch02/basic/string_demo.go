package main

import (
	"fmt"
	"unicode/utf8"
	"strings"
	"bytes"
)

func showLen() {
	tip1 := "genji is a ninjia"
	fmt.Println(len(tip1))

	tip2 := "忍者"
	fmt.Println(len(tip2))
}

func showLenOnUnicode()  {
	tip1 := "忍者"
	fmt.Println(utf8.RuneCountInString(tip1))

	tip2 := "龙刃出鞘, fight"
	fmt.Println(utf8.RuneCountInString(tip2))
}

func traversalString()  {
	theme := "狙击 start"

	for i := 0; i < len(theme); i++ {
		fmt.Printf("ascii: %c %d\n", theme[i], theme[i])
	}
}

func traversalStringOnUnicode() {
	theme := "狙击 start"

	for _, s := range theme {
		fmt.Printf("Unicode: %c %d\n", s, s)
	}
}

func seachSubstring() {
	tracer := "死神来了，死神bye bye"
	comma := strings.Index(tracer, "，")

	pos := strings.Index(tracer[comma:], "死神")
	fmt.Println(comma, pos, tracer[comma + pos])
}

func modifyString() {
	angel := "Heros never die"

	angleBytes := []byte(angel)

	for i := 5; i <= 10; i++ {
		angleBytes[i] = ' '
	}

	fmt.Println(string(angleBytes))
}

func stringBuilder() {
	hammer := "吃我一锤"
	sickle := "死吧"

	var stringBuilder bytes.Buffer

	stringBuilder.WriteString(hammer)
	stringBuilder.WriteString(sickle)

	fmt.Println(stringBuilder.String())
}

func main() {
	//showLen()
	//showLenOnUnicode()

	//traversalString()
	//traversalStringOnUnicode()

	//seachSubstring()

	//modifyString()

	stringBuilder()
}


