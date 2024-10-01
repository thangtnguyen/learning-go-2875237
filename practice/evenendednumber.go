package main

import (
	"fmt"
)

func main() {
	count := 0
	for i := 1000; i <= 9999; i++ {
		for j := i; j <= 9999; j++ {
			value := i * j
			stringValue := fmt.Sprintf("%d", value)
			if stringValue[0] == stringValue[len(stringValue)-1] {
				count++
			}
		}
	}

	fmt.Println(count)
}
