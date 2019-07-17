package main

import "fmt"

func fun(c chan int) {
	for i := 1; i <= 9; i++ {
		c <- i * i
	}
	close(c)
}

func main() {
	fmt.Println("main started")
	c := make(chan int)
	go fun(c)
	/*for {
		val, ok := <-c
		if ok == true {
			fmt.Println(val, ok)
		} else {
			fmt.Println("loop terminated")
			break
		}
	}*/
	for val := range c {
		fmt.Println(val)
	}
	fmt.Println("main stopped")
}
