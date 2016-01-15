package main

import "fmt"

func printWorld() {
	fmt.Println("world")
}

type printFunc func()

func(pf printFunc)chainHello() printFunc {
	return func() {
		fmt.Println("hello")
		pf()
	}
}

func main() {
	var printWorldFunc printFunc = printWorld
	printWorldFunc.chainHello()()
}
