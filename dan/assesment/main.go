package main

import "fmt"

func main() {

	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, slice := range slice {
		//fmt.Println(i, slice)
		if slice%2 == 0 {
			fmt.Println(slice, "is odd number")
		}
		if slice%2 != 0 {
			fmt.Println(slice, "is even number")
		}
	}
}
