package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(">>> Start Game Development >>> ")

	for {
		fmt.Print("\nCommand: ")
		command, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Println(err)
		}
		if string(command) == "exit\n" {
			break
		}
	}

	fmt.Println("Bye!")
}
