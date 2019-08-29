package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

const file_name = "/tmp/test.txt"

func main() {
	createFile(file_name)
	writeFile(file_name)
	readFile(file_name)
}

func createFile(name_file string) {
	var _, err = os.Stat(name_file)

	if len(name_file) > 0 {
		if os.IsNotExist(err) {
			f, err := os.OpenFile(name_file, os.O_CREATE, 0644)
			logFatal(err)
			defer f.Close()
			fmt.Println(name_file, "has been created")
		}
	} else {
		fmt.Println("Error no file name")
	}
}

func writeFile(name_file string) {
	f, err := os.OpenFile(name_file, os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()
	logFatal(err)
	text := []string{"Hello World", "from golang"}
	for _, value := range text {
		text_time := fmt.Sprintf("%s.%s", value, currentTime())
		fmt.Fprintln(f, text_time)
		logFatal(err)
	}
}

func readFile(name_file string) {
	f, err := os.OpenFile(name_file, os.O_RDONLY, os.ModePerm)
	logFatal(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func currentTime() string {
	return time.Now().Format(time.RFC850)
}
