package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	res, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	//fmt.Println("status code -> ", res.Body)
	scanner := bufio.NewScanner(res.Body)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
