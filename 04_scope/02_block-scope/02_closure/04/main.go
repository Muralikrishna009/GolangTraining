package main

import "fmt"

// func wrapper() func() int {
// 	x := 0
// 	return func() int {
// 		x++
// 		return x
// 	}
// }

func wrapper() func() int {
	x := 1
	return func() int {
		x++
		return x
	}
}

func main() {
	x  := 1
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(x)
}

/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope
*/
