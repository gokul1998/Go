package main

import "fmt"

func sum(a []int, c chan int) {
	s := 0
	for _, val := range a {
		fmt.Println(val)
		s += val
	}
	c <- s
}
func main() {
	a := []int{1, 2, 8, 6, 5, 4}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
