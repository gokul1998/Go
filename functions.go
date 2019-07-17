package main

import "fmt"
import "errors"

func isprime(n int) (int, bool, error) {
	res := true
	if n == 1 {
		return -1, false, errors.New("neither prime nor composite")
	}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			res = false
			break
		}
	}
	return n, res, nil
}
func main() {
	var a, b int
	fmt.Println("enter two integers : ")
	fmt.Scan(&a, &b)
	for i := a; i <= b; i++ {
		num, res, err := isprime(i)
		if res == true {
			fmt.Printf("%d is prime", num)
		} else if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%d is not prime", num)
		}
		fmt.Println()
	}
}

/*
func variadic(num ...int) int {
	fmt.Println(num)
	sum := 0
	for _, n := range num {
		sum += n
	}
	return sum
}
func main() {
	fmt.Println(variadic(1, 2, 3))
	fmt.Println(variadic(1, 2, 3, 4, 5))
	sl := []int{9, 8, 7}
	fmt.Println(variadic(sl...))
}*/
