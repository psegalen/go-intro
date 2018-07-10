package main

import "fmt"

var helloMessage = "Hello World"

func main() {
	fmt.Println(hello())
}

func hello() string {
	helloMessage = "Hi Planet"
	return helloMessage
}
