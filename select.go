package main

import (
	"fmt"
)

var c = make(chan int)

func h() {
	fmt.Println("xxx")

}

func main() {
	go h()

	select {}

	fmt.Println("123123")
}
