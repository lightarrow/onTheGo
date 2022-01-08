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
	x = 22

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
		if x < 20 || x > 30 {
			fmt.Println("x is out of range")
		}
		a := 11.0
		b := 20.0

		if frac := a / b; frac > 0.5 {
			fmt.Println("a is more than half of b")
		}

	}
	x = 2
	switch x {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("many")

	}
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("----")
	for i := 0; i < 3; i++ {
		if i > 1 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("----")
	for i := 0; i < 3; i++ {
		if i < 1 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("----")
	a := 0
	for a < 3 {
		fmt.Println(a)
		a++
	}

	fmt.Println("----")
	b := 0
	for {
		if b > 2 {
			break
		}
		fmt.Println(b)
		b++
	}
	for i := 1; i < 20; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizz buzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")

		} else {
			fmt.Println(i)
		}

	}
	book := "the colour of magic"
	fmt.Println(len(book))

	//Stings in go are immutable
	// slice
	// multiline strings are backtick quoted

}
