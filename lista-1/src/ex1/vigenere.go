package main

import (
	"fmt"
	"os"
)

// Só funfa com tudo minúsculo
func main() {
	key, message := parse()
	fmt.Printf("%s %s", key, message)

	fmt.Println("")
	result := encipher(key, message)
	fmt.Printf("%s\n", result)
}

func parse() (string, string) {
	args := os.Args
	return args[1], args[2]
}

func encipher(key, message string) string {
	out := make([]rune, 0, len(message))
	for i, v := range message {
		out = append(out, encodePair(v, rune(key[i%len(key)])))
	}
	return string(out)
}

func encodePair(a, b rune) rune {
	return (((a - 'a') + (b - 'a')) % 26) + 'a'
}
