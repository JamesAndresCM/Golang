package main

import (
	"fmt"
	"sync"
)

func task(msg string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Println(msg)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go task("jaime", &wg)
	wg.Add(1)
	go task("prueba", &wg)
	wg.Wait()
}
