package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	text := args[1]
	count := 0
	for _, char := range text {
		newChar := char + 3
		fmt.Printf("%c", newChar)
		count++
	}
	fmt.Printf("%d %d", len(text), count)
}
