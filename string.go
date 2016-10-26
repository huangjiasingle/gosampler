package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.LastIndex(`001.002.032.000	001.002.063.255 \\n 001.002.032.000	001.002.063.255`, "\n"))
}
