package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func init() {
	fmt.Println("Hello Otus!")

}

func main() {
	fmt.Println(reverse.String("Hello Otus!"))

}
