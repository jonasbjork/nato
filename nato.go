package main

/*
Nato translates text to the NATO phonetic alphabet.

Usage:
	nato <text>
	go run nato.go <text>

Author : Jonas Björk <jonas.bjork@gmail.com>
Date   : 2024-02-27
License: MIT
*/

import (
	"fmt"
	"strings"
	"os"
)

var natoAlphabet = map[string]string{
	"a": "Alpha", "b": "Bravo", "c": "Charlie",
	"d": "Delta", "e": "Echo", "f": "Foxtrot",
	"g": "Golf", "h": "Hotel", "i": "India",
	"j": "Juliett", "k": "Kilo", "l": "Lima",
	"m": "Mike", "n": "November", "o": "Oscar",
	"p": "Papa", "q": "Quebec", "r": "Romeo",
	"s": "Sierra", "t": "Tango", "u": "Uniform",
	"v": "Victor", "w": "Whiskey", "x": "X-ray",
	"y": "Yankee", "z": "Zulu",
}

func translate(text string) string {
    var translated []string
    for _, char := range text {

        if natoWord, exists := natoAlphabet[strings.ToLower(string(char))]; exists {
            translated = append(translated, natoWord)
        } else {
            translated = append(translated, string(char))
        }
    }
    return strings.Join(translated, " ")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: nato <text>")
		os.Exit(1)
	}
	text := os.Args[1]
	fmt.Println(translate(text))
}

