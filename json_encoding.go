package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type location struct {
  Lat float64 `json:"latitude"`
  Long float64 `json:"longitude"`
}

func exitOnErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func main() {
	curiosity := location{-3.94850, 123.2837}
	bytes, err := json.Marshal(curiosity)
	exitOnErr(err)
	fmt.Println(string(bytes))
}
