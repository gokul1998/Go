package main

import "fmt"

type emps interface {
	welcome()
	balance()
}
type juniorEmp struct {
	name   string
	age    int
	salary int
	months int
}
type seniorEmp struct {
	name   string
	salary int
	months int
}

func (j juniorEmp) welcome() {
	fmt.Println("welcome juniors")
}
func (j juniorEmp) balance() {
	fmt.Println(j.salary * j.months)
}
func (s seniorEmp) welcome() {
	fmt.Println("welcome senior")
}
func (s seniorEmp) balance() {
	fmt.Println(s.salary * s.months * 2)
}
func calculate(e emps) {
	fmt.Println(e)
	e.welcome()
	e.balance()
}
func main() {
	a := juniorEmp{"gokul", 21, 100, 5}
	b := seniorEmp{"krishna", 200, 6}
	calculate(a)
	calculate(b)
}
