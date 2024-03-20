package main

import "fmt"

// const (
// 	_  = iota             // 0
// 	KB = 1 << (iota * 10) // 1 << (1 * 10)
// 	MB = 1 << (iota * 10) // 1 << (2 * 10)
// 	GB = 1 << (iota * 10) // 1 << (3 * 10)
// 	TB = 1 << (iota * 10) // 1 << (4 * 10)
// )

const (
	num1 = iota
	num2
	num3
)

func main() {
	// fmt.Println("binary\t\tdecimal")
	// fmt.Printf("%b\t", KB)
	// fmt.Printf("%d\n", KB)
	// fmt.Printf("%b\t", MB)
	// fmt.Printf("%d\n", MB)
	// fmt.Printf("%b\t", GB)
	// fmt.Printf("%d\n", GB)
	// fmt.Printf("%b\t", TB)
	// fmt.Printf("%d\n", TB)

	// fmt.Println(num1)
	// fmt.Println(1 << num2)
	// fmt.Println(num3)

	for i := 0; i < 10; i++ {
		fmt.Printf("%d - %b \t %d - %b \t %d - %b\n", i, i, i<<2, i<<2, i>>2, i>>2)
	}

}
