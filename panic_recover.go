package main

import (
	"fmt"
	_ "log"
	"strconv"
)

func main() {
	fmt.Println(convert_string_to_int("2"))    // pass
	fmt.Println(convert_string_to_int("test")) // fail but recover
	fmt.Println(convert_string_to_int("3"))    // pass
}

func convert_string_to_int(text string) int {
	defer func() {
		if restored := recover(); restored != nil {
			fmt.Println("recovered", restored)
		}
	}()

	num, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}
	return num
}
