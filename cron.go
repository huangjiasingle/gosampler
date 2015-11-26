//本列子主要提供一种程序思维，实现心跳健康检查
package main

import (
	"fmt"
	"time"
)

func Heartbeat() {
	for {
		heartbeat()
		time.Sleep(time.Second * 5)
	}
}

func heartbeat() {
	fmt.Println("正在执行心跳健康检查。。。。")
}

func main() {
	Heartbeat()
}
