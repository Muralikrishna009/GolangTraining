package main

import "fmt"

func makeGreeter() func() string {
	x := 0
	return func() string {
		return "Hello world!" + string(rune(x))
	}
}

func main() {
	greet := makeGreeter()
	fmt.Println(greet())
	fmt.Printf("%T\n", greet)
}
