package main

import "fmt"

func main() {
	n := average(43, 56, 87, 12, 45, 57)
	s :=[]float64{1,2,3,4,5}
	sAns := average(s...)
	fmt.Println(n)
	fmt.Println(sAns)
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
