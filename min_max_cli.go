package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Error al menos dos parametros")
		os.Exit(1)
	}
	arguments := os.Args[1:]
	parse_data := convert_int(arguments)
	fmt.Println("MIN: ", min_max(parse_data)[0])
	fmt.Println("MAX: ", min_max(parse_data)[1])
}

func min_max(data []int) []int {
	values := []int{}
	min := data[0]
	max := data[0]
	for _, value := range data {
		if min < value {
			min = value
		}
		if max > value {
			max = value
		}
	}
	values = append(values, max)
	values = append(values, min)
	return values
}

func convert_int(data []string) []int {
	convert_num := []int{}
	for _, val := range data {
		parse, _ := strconv.Atoi(val)
		convert_num = append(convert_num, parse)
	}
	return convert_num
}
