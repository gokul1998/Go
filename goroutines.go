package main

import (
	"fmt"
	"time"
)

func f1() {
	for i := 0; i < 1000; i++ {
		time.Sleep(time.Millisecond)
		fmt.Println(i)
	}
}
func main() {
	fmt.Println("main started")
	go f1()
	time.Sleep(15 * time.Millisecond)
	fmt.Println("main terminated")
	fmt.Scanln()
}
