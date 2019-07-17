package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	fmt.Print(age)
	res := true
	for i := 2; i <= num/2; i++ {
		fmt.Print(i)
		if num%i == 0 {
			res = false
			break
		}
	}
	if res == true {
		fmt.Print("prime number")
	} else {
		fmt.Print("not a prime")
	}
}
