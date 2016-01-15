package main

import (
	"fmt"
)

type MyError struct {
	s string
}

func (f MyError) Error() string {
	return f.s
}

func returnsMyError() *MyError {
	return nil
}

func returnsMyErrorAssigned() *MyError {
	var p *MyError = nil
	return p
}

func returnsError() error {
	return nil
}

func returnsErrorAssigned() error {
	var p *MyError = nil
	return p
}

func main() {
	fmt.Printf("*MyError is nil: %v\n", returnsMyError() == nil)
	fmt.Printf("*MyError assigned nil is nil: %v\n", returnsMyErrorAssigned() == nil)
	fmt.Printf("error is nil: %v\n", returnsError() == nil)
	fmt.Printf("error assigned nil is nil: %v\n", returnsErrorAssigned() == nil)
}
