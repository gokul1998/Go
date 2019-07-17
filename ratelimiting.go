package main

import "fmt"
import "time"

func main() {
	/*requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	limit := time.Tick(200 * time.Millisecond)
	for val := range requests {
		<-limit
		fmt.Println("requests ", time.Now(), val)
	}*/
	burstylimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstylimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstylimiter <- t
		}
	}()
	burstyrequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyrequests <- i
	}
	close(burstyrequests)
	for req := range burstyrequests {
		<-burstylimiter
		fmt.Println("request ", req, time.Now())
	}
}
