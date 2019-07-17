package main

import "fmt"
import "strings"

/*
//grouping names by their length
func main() {
	var names = []string{"Katrina", "Evan", "Neil", "Adam", "Martin", "Matt",
		"Emma", "Isabella", "Emily", "Madison",
		"Ava", "Olivia", "Sophia", "Abigail",
		"Elizabeth", "Chloe", "Samantha",
		"Addison", "Natalie", "Mia", "Alexis"}
	m := make(map[int][]string)
	//fmt.Println(m, names)
	for _, name := range names {
		l := len(name)
		m[l] = append(m[l], name)
	}
	fmt.Println(m)
}
*/
func check(names []string, f func(string) bool) []string {
	res := make([]string, 0)
	for _, name := range names {
		if f(name) == true {
			res = append(res, name)
		}
	}
	return res
}
func main() {
	var names = []string{"Katrina", "Evan", "Neil", "Adam", "Martin", "Matt",
		"Emma", "Isabella", "Emily", "Madison",
		"Ava", "Olivia", "Sophia", "Abigail",
		"Elizabeth", "Chloe", "Samantha",
		"Addison", "Natalie", "Mia", "Alexis"}
	fmt.Println(check(names, func(s string) bool {
		return strings.HasPrefix(s, "Ad")
	}))
}
