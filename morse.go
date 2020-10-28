package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type CodeMorse struct {
	Input string `json:"input"`
}

type Response struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

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
	decode_morse := CodeMorse{"****,*,*-**,*-**,---,*--,---,*-*,*-**,-**"}
	encode_morse := CodeMorse{"Hello world"}

	fmt.Println(decode_morse.DecodeMorse())
	fmt.Println(encode_morse.EncodeMorse())
}

func logErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func (c *CodeMorse) DecodeMorse() string {
	result := []string{}
	split_string := strings.Split(c.Input, ",")

	for _, val := range split_string {
		for key, value := range Morse {
			if val == value {
				result = append(result, key)
			}
		}
	}

	response := Response{Input: c.Input, Output: strings.Join(result, "")}
	bytes, err := json.Marshal(response)
	logErr(err)
	return string(bytes)
}

func (c *CodeMorse) EncodeMorse() string {
	result := []string{}
	string_lower := strings.ToLower(c.Input)
	trim_string := strings.ReplaceAll(string_lower, " ", "")
	split_string := strings.Split(trim_string, "")

	for _, val := range split_string {
		for key, value := range Morse {
			if val == key {
				result = append(result, value)
			}
		}
	}

	response := Response{Input: c.Input, Output: strings.Join(result, ",")}
	bytes, err := json.Marshal(response)
	logErr(err)
	return string(bytes)
}
