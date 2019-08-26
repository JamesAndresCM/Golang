package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Error al menos un parametro")
		os.Exit(1)
	}

	argument := os.Args

	total := 0
	for _, val := range argument[1:] {
		parse, _ := strconv.Atoi(val)
		total += parse
	}
	fmt.Println(total)

}
