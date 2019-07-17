package main

import (
	"fmt"
	"io/ioutil"
)

<<<<<<< HEAD
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
=======
func fileexists(path, file string) (bool, string) {
	c, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return false, " "
	}
	for _, v := range c {
		if v.IsDir() == true && fileexists(path+string('/')+v.Name(), file) == true {
			return false, path + string('/') + v.Name()
		} else if v.IsDir() == false && v.Name() == file {
			return false, path
		}
	}
	return false, " "
>>>>>>> 0af78c263f31fd7474b190aa0e2cb2aa554b0887
}
func main() {
	var p, f string
	fmt.Println("enter the path : ")
	fmt.Scan(&p)
	fmt.Println("enter the file to check : ")
	fmt.Scan(&f)
<<<<<<< HEAD
	res := fileexists(p, f)
	if res == true {
		fmt.Println("file found")
=======
	res, path := fileexists(p, f)
	if res == true {
		fmt.Println("file found", path)
>>>>>>> 0af78c263f31fd7474b190aa0e2cb2aa554b0887
	} else {
		fmt.Println("file not found")
	}
}
