package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println("tick at ", t)
		}
	}()
	time.Sleep(2200 * time.Millisecond)
	ticker.Stop()
	fmt.Println("ticket stopped ")

}
