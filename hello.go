package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("My name is Dianita!")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
