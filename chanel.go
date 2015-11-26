//计算1000以下能被3或者5整除的整数的和

package main

import (
	"fmt"
	"runtime"
	"time"
)

//计算某个范围内能被某个书整除的整数的和
func get_sum_divisible(num, divier int, result chan int) {
	sum := 0
	for i := 0; i < divier; i++ {
		if i%num == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
	result <- sum
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	LIMIT := 1000
	job1 := make(chan int, 1)
	job2 := make(chan int, 2)

	t_start := time.Now()
	go get_sum_divisible(3, LIMIT, job2)
	go get_sum_divisible(5, LIMIT, job2)
	go get_sum_divisible(15, LIMIT, job1)

	sum15 := <-job1
	sum3, sum5 := <-job2, <-job2
	sum := sum3 + sum5 - sum15
	t_end := time.Now()
	fmt.Println(sum)
	fmt.Println(t_end.Sub(t_start))
}
