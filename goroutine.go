package main

import (
	"fmt"
	"time"
)

func list_elem(num int) {
	for i := 0; i < num; i++ {
		fmt.Println(i)
	}
}

func main() {
	go list_elem(5)
	time.Sleep(100)
}
