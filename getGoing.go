// package onTheGo
/*
 * Author Ronald Johnson 2022
 */
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome Gophers")
	x, y := 1.0, 2.0
	fmt.Printf("x=%v, type of %T\n", x, x)
	fmt.Printf("y=%v, type of %T\n", y, y)

	mean := (x + y) / 2.0
	fmt.Printf("result: %v, type of %T\n", mean, mean)

	// conditions
	if x > 5 {
		fmt.Println("x is big")
		if x > 100 {
			fmt.Println("x is very big")
		} else {
			fmt.Println("x is not that big")
		}
		if x > 5 && x < 15 {
			fmt.Println("x is just right")
		}
	}

}
