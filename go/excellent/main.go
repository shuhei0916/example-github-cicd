package main

import "fmt"

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}

func main() {
	fmt.Println(1 + 2)
	fmt.Println(EvenOrOdd(57))
	fmt.Println(EvenOrOdd(12))
}
