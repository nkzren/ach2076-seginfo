package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	key, text := parse()
	cipher := make(map[rune]rune)
	a := 'a'
	for _, char := range key {
		cipher[a] = char
		a++
	}
	fmt.Printf(decipher(cipher, text))

}

func decipher(cipher map[rune]rune, text string) string {
	var b strings.Builder
	for _, char := range text {
		c := cipher[char]
		b.WriteRune(c)
	}
	return b.String()
}

func parse() (string, string) {
	args := os.Args
	key := args[1]
	text := args[2]
	return key, text
}
