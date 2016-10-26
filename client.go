package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://api.github.com/users/huangjiasingle"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("authorization", "Basic aHVhbmdqaWFzaW5nbGU6aGxoajUyNTExNQ==")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "5fd45ac5-4049-21f6-89de-e6cda8342670")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	sDec := base64.StdEncoding.EncodeToString([]byte("huangjiasingle:hlhj525115"))
	fmt.Println(string(sDec))

	queue := make(chan string, 5)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem, i := range queue {
		fmt.Println(elem)
		fmt.Println(i)
	}
}
