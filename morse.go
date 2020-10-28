package main

import (
	"fmt"
	"strings"
)

var Morse = map[string]string{"a": "*-",
	"b": "-***",
	"c": "-*-*",
	"d": "-**",
	"e": "*",
	"f": "**-*",
	"g": "--*",
	"h": "****",
	"i": "**",
	"j": "*---",
	"k": "-*-",
	"l": "*-**",
	"m": "--",
	"n": "-*",
	"o": "---",
	"p": "*--*",
	"q": "--*-",
	"r": "*-*",
	"s": "***",
	"t": "-",
	"u": "**-",
	"v": "***-",
	"w": "*--",
	"x": "-**-",
	"y": "-*--",
	"z": "--**",
}

func main() {
	words := "****,*,*-**,*-**,---,*--,---,*-*,*-**,-**"
	text := "Hello world"
	fmt.Println(DecodeMorse(words))
	fmt.Println(EncodeMorse(text))
}

func DecodeMorse(attr string) (input, output string) {
	result := []string{}
	split_string := strings.Split(attr, ",")

	for _, val := range split_string {
		for key, value := range Morse {
			if val == value {
				result = append(result, key)
			}
		}
	}
	return "Input: " + attr, "Output: " + strings.Join(result, "")
}

func EncodeMorse(attr string) (input, output string) {
	result := []string{}
	string_lower := strings.ToLower(attr)
	trim_string := strings.ReplaceAll(string_lower, " ", "")
	split_string := strings.Split(trim_string, "")

	for _, val := range split_string {
		for key, value := range Morse {
			if val == key {
				result = append(result, value)
			}
		}
	}

	return "Input: " + attr, "Output: " + strings.Join(result, ",")
}
