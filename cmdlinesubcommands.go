package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("-----subcommands------")
	foocmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooenable := foocmd.Bool("boolean", false, "this is a boolean")
	fooname := foocmd.String("name", "John", "this is a name in foo")
	barcmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barname := barcmd.String("name", "Cena", "this is a name in bar")
	barage := barcmd.Int("age", 99, "this is a age in bar cmd")
	switch os.Args[1] {
	case "foo":
		foocmd.Parse(os.Args[2:])
		fmt.Println("subcommand FOO : ")
		fmt.Println("boolean value : ", *fooenable)
		fmt.Println("name : ", *fooname)
		fmt.Println("tail : ", foocmd.Args())
	case "bar":
		barcmd.Parse(os.Args[2:])
		fmt.Println("BAR subcommand :")
		fmt.Println("name : ", *barname)
		fmt.Println("age : ", *barage)
		fmt.Println("tail : ", barcmd.Args())
	default:
		fmt.Println("expected foo or bar subcommands")
		os.Exit(1)
	}
}
