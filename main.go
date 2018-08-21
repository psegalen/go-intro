package main

import (
	"fmt"

	"github.com/psegalen/go-intro/pierre"
)

var helloMessage = "Hello World"

func main() {
	fmt.Println(pierre.Hello())
}

func hello() string {
	helloMessage = "Hi Planet"
	return helloMessage
}
