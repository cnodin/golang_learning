package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println(os.Getegid())
	fmt.Println(os.Getppid())

	cmd0 := exec.Command("echo", "-n", "My first command comes from golang1.你好")
	//if err := cmd0.Start(); err != nil {
	//	fmt.Printf("Error: The command No.0 can not be run: %s\n", err)
	//	return
	//}

	stdout0, err := cmd0.StdoutPipe()
	if err != nil {
		fmt.Printf("Error: Couldn't obtain the stdout pipe for command No.0: %s\n", err)
		return
	}

	if err := cmd0.Start(); err != nil {
		fmt.Printf("Error: The command No.0 can not be run: %s\n", err)
		return
	}

	//var outputBuf0 bytes.Buffer
	//for {
	//	tempOutput := make([]byte, 5)
	//	n, err := stdout0.Read(tempOutput)
	//	if err != nil {
	//		if err == io.EOF {
	//			break
	//		} else {
	//			fmt.Printf("Error: Couldn't read data from the pipe: %s\n", err)
	//			return
	//		}
	//	}
	//	if n > 0 {
	//		outputBuf0.Write(tempOutput[:n])
	//	}
	//}
	//fmt.Printf("%s\n", outputBuf0.String())

	//output0 := make([]byte, 60)
	//n, err := stdout0.Read(output0)
	//if err != nil {
	//	fmt.Printf("Error: Couldn't read data from the pipe: %s\n", err)
	//	return
	//}
	//fmt.Printf("%s\n", output0[:n])

	outputBuf0 := bufio.NewReader(stdout0)
	output0, _, err := outputBuf0.ReadLine()
	if err != nil {
		fmt.Printf("Error: Couldn't read data from the pipe: %s\n", output0)
		return
	}
	fmt.Printf("%s\n", output0)
}
