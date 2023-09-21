package main

import (
	"fmt"
	"os"
)

func main() {
	var commitMessage []string

	if len(os.Args) >= 2 {
		commitMessage = append(commitMessage, os.Args[1:]...)
	}
	fmt.Println([]byte(commitMessage[1]))
}
