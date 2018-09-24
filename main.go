package main

import (
	"fmt"
	"time"
)

func main() {
	//reader := bufio.NewReader(os.Stdin)
	fmt.Println(">>> Start Game Development >>> ")
	command := "#"

	for {
		fmt.Println("Command: ")
		fmt.Scanln(command)

		time.Sleep(5 * time.Second)

		if command == "#" {
			break
		}

		//switch expression {
		//case condition:

		//default condition:
		//continue;
		//}
	}

	fmt.Println("Bye!")
}
