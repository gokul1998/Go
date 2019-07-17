package main

import (
	"flag"
	"fmt"
)

func main() {
	wordptr := flag.String("name", "John", "this is a name")
	ageptr := flag.Int("age", 21, "this is the age")
	flag.Parse()
	fmt.Println("name : ", *wordptr)
	fmt.Println("age : ", *ageptr)
	fmt.Println("tail : ", flag.Args())
}
