package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	c, err := os.Getwd()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(c)
		return
	}
}
