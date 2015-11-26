package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func list_elem(num int, tag string) {
	for i := 0; i < num; i++ {
		fmt.Println(tag, i)
		tick := time.Duration(rand.Intn(10)) //rand.Intn(n) 返回 0-n以内的伪随机数，如果n<=0，则会报panic
		time.Sleep(time.Millisecond * tick)
	}
}

func main() {
	// go list_elem(10, "tag_a")
	// go list_elem(20, "tag_b")
	// // var input string
	// // fmt.Scanln(&input)
	// fmt.Println(runtime.NumCPU())
	// time.Sleep(2 * 1e9)
	runtime.GOMAXPROCS(runtime.NumCPU())
	ch := make(chan int)

	go func(ch chan int) {
		time.Sleep(1 * 1e9)
		ch <- 1
	}(ch)

	go func(ch chan int) {
		for {
			select {
			case <-ch:
				println("got!!!")
				return
			default:
				println("waiting...")
				time.Sleep(1 * 1e9)
			}
		}
	}(ch)

	time.Sleep(2 * 1e9)
}
