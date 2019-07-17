package main

import (
	"fmt"
	"sort"
)

/*type byage []string

func (s bylen) Len() int {
	return len(s)
}
func (a bylen) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a bylen) Less(i, j int) bool {
	if len(a[i]) == len(a[j]) {
		return a[i] < a[j]
	}
	return len(a[i]) < len(a[j])
}
func main() {
	/*s := []string{"aaz", "aab", "grtr", "ffef", "zfe"}
	fmt.Println(sort.StringsAreSorted(s))
	sort.Strings(s)
	fmt.Println(sort.StringsAreSorted(s))

	sort.Sort(bylen(s))
	fmt.Println(s)
*/
type byage []person

func (a byage) Len() int {
	return len(a)
}
func (a byage) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a byage) Less(i, j int) bool {
	if a[i].age == a[j].age {
		return a[i].name < a[j].name
	}
	return a[i].age < a[j].age
}

type person struct {
	name string
	age  int
}

func main() {
	var n int
	fmt.Println("enter the number of persons : ")
	fmt.Scan(&n)
	var a = make([]person, n)
	//fmt.Println(a)
	for i := 1; i <= n; i++ {
		fmt.Println("enter name of person ", i)
		fmt.Scan(&a[i-1].name)
		fmt.Println("enter age of person ", i)
		fmt.Scan(&a[i-1].age)
	}
	fmt.Println(a)
	sort.Sort(byage(a))
	fmt.Println(a)
}
