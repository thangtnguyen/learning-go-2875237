package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i <= 20; i++ {
		fmt.Printf("%v\n", FizzBuzz(i))
	}
}

func FizzBuzz(value int) string {
	if value%3 == 0 && value%5 == 0 {
		return "fizz buzz"
	} else if value%3 == 0 {
		return "fizz"
	} else if value%5 == 0 {
		return "buzz"
	} else {
		return strconv.Itoa(value)
	}
}
