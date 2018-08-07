package main

import (
	"fmt"
	"time"

	"gopkg.in/robfig/cron.v2"
)

func main() {
	tab := cron.New()
	tab.Start()
	tab.AddFunc("0 0/1 * * * ?", func() {
		fmt.Println(time.Now())
	})
	time.Sleep(time.Second * 3600)

}
