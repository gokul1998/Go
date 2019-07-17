package main

import (
	"fmt"
	"io/ioutil"
)

func fileexists(path, file string) bool {
	c, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return false
	}
	for _, v := range c {
		if v.IsDir() == true && fileexists(path+string('/')+v.Name(), file) == true {
			return true
		} else if v.IsDir() == false && v.Name() == file {
			return true
		}
	}
	return false
}
func main() {
	var p, f string
	fmt.Println("enter the path : ")
	fmt.Scan(&p)
	fmt.Println("enter the file to check : ")
	fmt.Scan(&f)
	res := fileexists(p, f)
	if res == true {
		fmt.Println("file found")
	} else {
		fmt.Println("file not found")
	}
}
