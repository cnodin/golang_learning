package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//r := bufio.NewReader(os.Stdin)
	r := bufio.NewScanner(os.Stdin)
	for r.Scan() {
		fmt.Println(r.Text())
	}
	//for {
		//c, err := r.ReadString('\n')

		//if err == nil {
		//	c = strings.Replace(c, "吗", "", -1)
		//	c = strings.Replace(c, "?", "!", -1)
		//	c = strings.Replace(c, "？", "!", -1)
		//	fmt.Println(c)
		//}
	//}
}
